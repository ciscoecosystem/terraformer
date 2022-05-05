package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tACACSPlusProviderClassName = "aaaTacacsPlusProvider"

type TACACSPlusProviderGenerator struct {
	ACIService
}

func (a *TACACSPlusProviderGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tACACSPlusProviderClassName)

	TACACSPlusProviderCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	TACACSPlusProviderCount, err := strconv.Atoi(stripQuotes(TACACSPlusProviderCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < TACACSPlusProviderCount; i++ {
		TACACSPlusProviderAttr := TACACSPlusProviderCont.S("imdata").Index(i).S(tACACSPlusProviderClassName, "attributes")
		TACACSPlusProviderDN := G(TACACSPlusProviderAttr, "dn")
		if filterChildrenDn(TACACSPlusProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				TACACSPlusProviderDN,
				resourceNamefromDn(tACACSPlusProviderClassName, TACACSPlusProviderDN, i),
				"aci_tacacs_provider",
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
