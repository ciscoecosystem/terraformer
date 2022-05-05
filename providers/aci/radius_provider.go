package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const rADIUSProviderClassName = "aaaRadiusProvider"

type RADIUSProviderGenerator struct {
	ACIService
}

func (a *RADIUSProviderGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, rADIUSProviderClassName)

	RADIUSProviderCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RADIUSProviderCount, err := strconv.Atoi(stripQuotes(RADIUSProviderCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RADIUSProviderCount; i++ {
		RADIUSProviderAttr := RADIUSProviderCont.S("imdata").Index(i).S(rADIUSProviderClassName, "attributes")
		RADIUSProviderDN := G(RADIUSProviderAttr, "dn")
		if filterChildrenDn(RADIUSProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RADIUSProviderDN,
				resourceNamefromDn(rADIUSProviderClassName, RADIUSProviderDN, i),
				"aci_radius_provider",
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
