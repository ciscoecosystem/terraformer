package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const samlProviderClassName = "aaaSamlProvider"

type SAMLProviderGenerator struct {
	ACIService
}

func (a *SAMLProviderGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, samlProviderClassName)
	samlProviderCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	samlProviderCount, err := strconv.Atoi(stripQuotes(samlProviderCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < samlProviderCount; i++ {
		samlProviderAttr := samlProviderCont.S("imdata").Index(i).S(samlProviderClassName, "attributes")
		samlProviderDN := G(samlProviderAttr, "dn")
		if filterChildrenDn(samlProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				samlProviderDN,
				resourceNamefromDn(samlProviderClassName, samlProviderDN, i),
				"aci_saml_provider",
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
