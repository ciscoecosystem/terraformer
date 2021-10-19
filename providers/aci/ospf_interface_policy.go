package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const ospfInterfacePolicyClassName = "ospfIfPol"

type ospfInterfacePolicyGenerator struct {
	ACIService
}

func (a *ospfInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, ospfInterfacePolicyClassName)

	ospfInterfacePoliciesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ospfInterfacePolicyCount, err := strconv.Atoi(stripQuotes(ospfInterfacePoliciesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ospfInterfacePolicyCount; i++ {
		ospfInterfacePolicyDN := stripQuotes(ospfInterfacePoliciesCont.S("imdata").Index(i).S(ospfInterfacePolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(ospfInterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				ospfInterfacePolicyDN,
				ospfInterfacePolicyDN,
				"aci_ospf_interface_policy",
				"aci",
				[]string{
					"cost",
					"ctrl",
					"dead_intvl",
					"hello_intvl",
					"name_alias",
					"nw_t",
					"pfx_suppress",
					"prio",
					"rexmit_intvl",
					"xmit_delay",
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
