package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const sAMLProviderGroupClassName = "aaaSamlProviderGroup"

type SAMLProviderGroupGenerator struct {
	ACIService
}

func (a *SAMLProviderGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, sAMLProviderGroupClassName)

	SAMLProviderGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SAMLProviderGroupCount, err := strconv.Atoi(stripQuotes(SAMLProviderGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SAMLProviderGroupCount; i++ {
		SAMLProviderGroupAttr := SAMLProviderGroupCont.S("imdata").Index(i).S(sAMLProviderGroupClassName, "attributes")
		SAMLProviderGroupDN := G(SAMLProviderGroupAttr, "dn")
		if filterChildrenDn(SAMLProviderGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				SAMLProviderGroupDN,
				resourceNamefromDn(sAMLProviderGroupClassName, SAMLProviderGroupDN, i),
				"aci_saml_provider_group",
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
