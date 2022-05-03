package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const rtctrlSetAddCommClassName = "rtctrlSetAddComm"

type RtctrlSetAddCommGenerator struct {
	ACIService
}

func (a *RtctrlSetAddCommGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, rtctrlSetAddCommClassName)

	RtctrlSetAddCommCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RtctrlSetAddCommCount, err := strconv.Atoi(stripQuotes(RtctrlSetAddCommCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RtctrlSetAddCommCount; i++ {
		RtctrlSetAddCommAttr := RtctrlSetAddCommCont.S("imdata").Index(i).S(rtctrlSetAddCommClassName, "attributes")
		RtctrlSetAddCommDN := G(RtctrlSetAddCommAttr, "dn")
		community := G(RtctrlSetAddCommAttr, "community")
		if filterChildrenDn(RtctrlSetAddCommDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RtctrlSetAddCommDN,
				resourceNamefromDn(rtctrlSetAddCommClassName, RtctrlSetAddCommDN, i),
				"aci_action_rule_additional_communities",
				"aci",
				map[string]string{
					"action_rule_profile_dn": GetParentDn(RtctrlSetAddCommDN, fmt.Sprintf("/saddcomm-%s", community)),
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
