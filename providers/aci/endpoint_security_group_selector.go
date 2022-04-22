package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const applicationEndpointSecurityGroupSelectorClass = "fvEPSelector"

type ApplicationEndpointSecurityGroupSelectorGenerator struct {
	ACIService
}

func (a *ApplicationEndpointSecurityGroupSelectorGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, applicationEndpointSecurityGroupSelectorClass)

	endpointSecurityGroupSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	endpointSecurityGroupSelectorCount, err := strconv.Atoi(stripQuotes(endpointSecurityGroupSelectorCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < endpointSecurityGroupSelectorCount; i++ {
		endpointSecurityGroupSelectorProfileDN := stripQuotes(endpointSecurityGroupSelectorCont.S("imdata").Index(i).S(applicationEndpointSecurityGroupSelectorClass, "attributes", "dn").String())
		if filterChildrenDn(endpointSecurityGroupSelectorProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				endpointSecurityGroupSelectorProfileDN,
				resourceNamefromDn(applicationEndpointSecurityGroupSelectorClass, (endpointSecurityGroupSelectorProfileDN), i),
				"aci_endpoint_security_group_selector",
				"aci",
				[]string{
					"name",
					"annotation",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
