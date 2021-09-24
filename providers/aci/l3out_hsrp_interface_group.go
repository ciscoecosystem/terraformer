package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutHSRPInterfaceGroupClass = "hsrpGroupP"

type L3OutHSRPInterfaceGroupGenerator struct {
	ACIService
}

func (a *L3OutHSRPInterfaceGroupGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutHSRPInterfaceGroupClass)

	L3OutHSRPInterfaceGroupCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3OutHSRPInterfaceGroupCount, err := strconv.Atoi(stripQuotes(L3OutHSRPInterfaceGroupCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3OutHSRPInterfaceGroupCount; i++ {
		L3OutHSRPInterfaceGroupDN := stripQuotes(L3OutHSRPInterfaceGroupCont.S("imdata").Index(i).S(L3OutHSRPInterfaceGroupClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			L3OutHSRPInterfaceGroupDN,
			L3OutHSRPInterfaceGroupDN,
			"aci_l3out_hsrp_interface_group",
			"aci",
			[]string{
				"config_issues",
				"group_af",
				"group_id",
				"group_name",
				"ip",
				"ip_obtain_mode",
				"mac",
				"name_alias",
				"relation_hsrp_rs_group_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
