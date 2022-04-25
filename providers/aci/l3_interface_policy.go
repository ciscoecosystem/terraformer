package aci

import (
	"fmt"
	"strconv"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l3InterfacePolicyClassName = "l3IfPol"

type L3InterfacePolicyGenerator struct {
	ACIService
}

func (a *L3InterfacePolicyGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l3InterfacePolicyClassName)
	L3InterfacePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	L3InterfacePolicyCount, err := strconv.Atoi(stripQuotes(L3InterfacePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < L3InterfacePolicyCount; i++ {
		L3InterfacePolicyAttr := L3InterfacePolicyCont.S("imdata").Index(i).S(l3InterfacePolicyClassName, "attributes")
		L3InterfacePolicyDN := G(L3InterfacePolicyAttr, "dn")
		name := G(L3InterfacePolicyAttr, "name")
		if filterChildrenDn(L3InterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				L3InterfacePolicyDN,
				name,
				"aci_l3_interface_policy",
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
