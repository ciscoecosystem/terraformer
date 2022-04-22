package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const StaticNodeMgmtAddressClassName = "mgmtRsInBStNode"

type StaticNodeMgmtAddressGenerator struct {
	ACIService
}

func (a *StaticNodeMgmtAddressGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, StaticNodeMgmtAddressClassName)

	StaticNodeMgmtAddressCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	StaticNodeMgmtAddressCount, err := strconv.Atoi(stripQuotes(StaticNodeMgmtAddressCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < StaticNodeMgmtAddressCount; i++ {
		StaticNodeMgmtAddressDN := stripQuotes(StaticNodeMgmtAddressCont.S("imdata").Index(i).S(StaticNodeMgmtAddressClassName, "attributes", "dn").String())
		if filterChildrenDn(StaticNodeMgmtAddressDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				StaticNodeMgmtAddressDN,
				resourceNamefromDn(StaticNodeMgmtAddressClassName, StaticNodeMgmtAddressDN, i),
				"aci_static_node_mgmt_address",
				"aci",
				map[string]string{
					"type": "in_band",
				},
				[]string{
					"addr",
					"gw",
					"v6_addr",
					"v6_gw",
					"annotation",
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	dnURL = fmt.Sprintf("%s/%s.json", baseURL, "mgmtRsOoBStNode")

	StaticNodeMgmtAddressCont, err = client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	StaticNodeMgmtAddressCount, err = strconv.Atoi(stripQuotes(StaticNodeMgmtAddressCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < StaticNodeMgmtAddressCount; i++ {
		StaticNodeMgmtAddressDN := stripQuotes(StaticNodeMgmtAddressCont.S("imdata").Index(i).S("mgmtRsOoBStNode", "attributes", "dn").String())
		fmt.Printf("StaticNodeMgmtAddressDN: %v\n", StaticNodeMgmtAddressDN)
		if filterChildrenDn(StaticNodeMgmtAddressDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				StaticNodeMgmtAddressDN,
				resourceNamefromDn("mgmtRsOoBStNode",StaticNodeMgmtAddressDN, i),
				"aci_static_node_mgmt_address",
				"aci",
				map[string]string{
					"type": "out_of_band",
				},
				[]string{
					"addr",
					"gw",
					"v6_addr",
					"v6_gw",
					"annotation",
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
