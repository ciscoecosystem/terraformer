package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const duoProviderGroupClassName = "aaaDuoProviderGroup"

type DuoProviderGroupGenerator struct {
	ACIService
}

func (a *DuoProviderGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, duoProviderGroupClassName)

	DuoProviderGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	DuoProviderGroupCount, err := strconv.Atoi(stripQuotes(DuoProviderGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < DuoProviderGroupCount; i++ {
		DuoProviderGroupAttr := DuoProviderGroupCont.S("imdata").Index(i).S(duoProviderGroupClassName, "attributes")
		DuoProviderGroupDN := G(DuoProviderGroupAttr, "dn")
		if filterChildrenDn(DuoProviderGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				DuoProviderGroupDN,
				resourceNamefromDn(duoProviderGroupClassName, DuoProviderGroupDN, i),
				"aci_duo_provider_group",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
