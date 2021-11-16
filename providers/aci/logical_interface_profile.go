package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const logicalInterfaceProfileClassName = "l3extLIfP"

type LogicalInterfaceProfileGenerator struct {
	ACIService
}

func (a *LogicalInterfaceProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, logicalInterfaceProfileClassName)

	LogicalInterfaceProfilesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LogicalInterfaceProfileCount, err := strconv.Atoi(stripQuotes(LogicalInterfaceProfilesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LogicalInterfaceProfileCount; i++ {
		LogicalInterfaceProfileDN := stripQuotes(LogicalInterfaceProfilesCont.S("imdata").Index(i).S(logicalInterfaceProfileClassName, "attributes", "dn").String())
		if filterChildrenDn(LogicalInterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				LogicalInterfaceProfileDN,
				LogicalInterfaceProfileDN,
				"aci_logical_interface_profile",
				"aci",
				[]string{
					"name_alias",
					"prio",
					"tag",
					"relation_l3ext_rs_l_if_p_to_netflow_monitor_pol",
					"relation_l3ext_rs_egress_qos_dpp_pol",
					"relation_l3ext_rs_ingress_qos_dpp_pol",
					"relation_l3ext_rs_l_if_p_cust_qos_pol",
					"relation_l3ext_rs_arp_if_pol",
					"relation_l3ext_rs_nd_if_pol",
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
