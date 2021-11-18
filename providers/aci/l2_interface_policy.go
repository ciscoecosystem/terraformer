package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L2InterfacePolicyClass = "l2IfPol"

type L2InterfacePolicyGenerator struct {
	ACIService
}

func (a *L2InterfacePolicyGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L2InterfacePolicyClass)
	L2InterfacePolicyCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L2InterfacePolicysCount, err := strconv.Atoi(stripQuotes(L2InterfacePolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L2InterfacePolicysCount; i++ {
		L2InterfacePolicyDN := stripQuotes(L2InterfacePolicyCont.S("imdata").Index(i).S(L2InterfacePolicyClass, "attributes", "dn").String())
		if filterChildrenDn(L2InterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L2InterfacePolicyDN,
				L2InterfacePolicyDN,
				"aci_l2_interface_policy",
				"aci",
				[]string{
					"qinq",
					"vepa",
					"vlan_scope",
					"name_alias",
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
