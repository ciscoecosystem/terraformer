package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const leafBreakoutPortGrpClass = "infraBrkoutPortGrp"

type LeafBreakoutPortGrpGenerator struct {
	ACIService
}

func (a *LeafBreakoutPortGrpGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, leafBreakoutPortGrpClass)

	leafBreakoutPortGrpCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(leafBreakoutPortGrpCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	leafBreakoutPortGrpCount, err := strconv.Atoi(stripQuotes(leafBreakoutPortGrpCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < leafBreakoutPortGrpCount; i++ {
		leafBreakoutPortGrpProfileDN := stripQuotes(leafBreakoutPortGrpCont.S("imdata").Index(i).S(leafBreakoutPortGrpClass, "attributes", "dn").String())
		if filterChildrenDn(leafBreakoutPortGrpProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				leafBreakoutPortGrpProfileDN,
				resourceNamefromDn(leafBreakoutPortGrpClass, (leafBreakoutPortGrpProfileDN), i),
				"aci_leaf_breakout_port_group",
				"aci",
				[]string{
					"name_alias",
					"brkout_map",
					"relation_infra_rs_mon_brkout_infra_pol",
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
