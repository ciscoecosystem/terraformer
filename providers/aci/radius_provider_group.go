package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const radiusProviderGroupClassName = "aaaRadiusProviderGroup"

type RadiusProviderGroupGenerator struct {
	ACIService
}

func (a *RadiusProviderGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, radiusProviderGroupClassName)
	radiusProviderGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	radiusProviderGroupCount, err := strconv.Atoi(stripQuotes(radiusProviderGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < radiusProviderGroupCount; i++ {
		radiusProviderGroupAttr := radiusProviderGroupCont.S("imdata").Index(i).S(radiusProviderGroupClassName, "attributes")
		radiusProviderGroupDN := G(radiusProviderGroupAttr, "dn")
		if filterChildrenDn(radiusProviderGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				radiusProviderGroupDN,
				resourceNamefromDn(radiusProviderGroupClassName, radiusProviderGroupDN, i),
				"aci_radius_provider_group",
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
