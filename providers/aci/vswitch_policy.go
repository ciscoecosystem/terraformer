package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const VswitchPolicyClass = "vmmVSwitchPolicyCont"

type VswitchPolicyGenerator struct {
	ACIService
}

func (a *VswitchPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, VswitchPolicyClass)

	VswitchPolicyCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	VswitchPolicyCount, err := strconv.Atoi(stripQuotes(VswitchPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < VswitchPolicyCount; i++ {
		VswitchPolicyDN := VswitchPolicyCont.S("imdata").Index(i).S(VswitchPolicyClass, "attributes", "dn").String()
		if filterChildrenDn(VswitchPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(VswitchPolicyDN),
				stripQuotes(VswitchPolicyDN),
				"aci_vswitch_policy",
				"aci",
				[]string{
					"relation_vmm_rs_vswitch_exporter_pol",
					"relation_vmm_rs_vswitch_override_cdp_if_pol",
					"relation_vmm_rs_vswitch_override_fw_pol",
					"relation_vmm_rs_vswitch_override_lacp_pol",
					"relation_vmm_rs_vswitch_override_lldp_if_pol",
					"relation_vmm_rs_vswitch_override_mcp_if_pol",
					"relation_vmm_rs_vswitch_override_mtu_pol",
					"relation_vmm_rs_vswitch_override_stp_pol",
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
