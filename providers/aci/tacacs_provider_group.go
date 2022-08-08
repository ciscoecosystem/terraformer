package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tACACSPlusProviderGroupClassName = "aaaTacacsPlusProviderGroup"

type TACACSPlusProviderGroupGenerator struct {
	ACIService
}

func (a *TACACSPlusProviderGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tACACSPlusProviderGroupClassName)

	TACACSPlusProviderGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	TACACSPlusProviderGroupCount, err := strconv.Atoi(stripQuotes(TACACSPlusProviderGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < TACACSPlusProviderGroupCount; i++ {
		TACACSPlusProviderGroupAttr := TACACSPlusProviderGroupCont.S("imdata").Index(i).S(tACACSPlusProviderGroupClassName, "attributes")
		TACACSPlusProviderGroupDN := G(TACACSPlusProviderGroupAttr, "dn")
		if filterChildrenDn(TACACSPlusProviderGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				TACACSPlusProviderGroupDN,
				resourceNamefromDn(tACACSPlusProviderGroupClassName, TACACSPlusProviderGroupDN, i),
				"aci_tacacs_provider_group",
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
