package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const interfaceFCPolicyClassName = "fcIfPol"

type InterfaceFCPolicyGenerator struct {
	ACIService
}

func (a *InterfaceFCPolicyGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, interfaceFCPolicyClassName)

	InterfaceFCPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	InterfaceFCPolicyCount, err := strconv.Atoi(stripQuotes(InterfaceFCPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < InterfaceFCPolicyCount; i++ {
		InterfaceFCPolicyAttr := InterfaceFCPolicyCont.S("imdata").Index(i).S(interfaceFCPolicyClassName, "attributes")
		InterfaceFCPolicyDN := G(InterfaceFCPolicyAttr, "dn")
		if filterChildrenDn(InterfaceFCPolicyDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				InterfaceFCPolicyDN,
				InterfaceFCPolicyDN,
				"aci_interface_fc_policy",
				"aci",
				map[string]string{},
				[]string{
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
