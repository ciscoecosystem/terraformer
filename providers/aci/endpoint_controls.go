package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const endpointControlPolicyClassName = "epControlP"

type EndpointControlPolicyGenerator struct {
	ACIService
}

func (a *EndpointControlPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, endpointControlPolicyClassName)

	EndpointControlPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	EndpointControlPolicyCount, err := strconv.Atoi(stripQuotes(EndpointControlPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < EndpointControlPolicyCount; i++ {
		EndpointControlPolicyAttr := EndpointControlPolicyCont.S("imdata").Index(i).S(endpointControlPolicyClassName, "attributes")
		EndpointControlPolicyDN := G(EndpointControlPolicyAttr, "dn")
		if filterChildrenDn(EndpointControlPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				EndpointControlPolicyDN,
				resourceNamefromDn(endpointControlPolicyClassName, EndpointControlPolicyDN, i),
				"aci_endpoint_controls",
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
