package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vmmDomClass = "vmmDomP"

type VmmDomGenerator struct {
	ACIService
}

func (a *VmmDomGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vmmDomClass)

	vmmDomCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(vmmDomCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	vmmDomCount, err := strconv.Atoi(stripQuotes(vmmDomCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < vmmDomCount; i++ {
		vmmDomProfileDN := stripQuotes(vmmDomCont.S("imdata").Index(i).S(vmmDomClass, "attributes", "dn").String())
		if filterChildrenDn(vmmDomProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				vmmDomProfileDN,
				vmmDomProfileDN,
				"aci_vmm_domain",
				"aci",
				[]string{
					"ctrl_knob",
					"delimiter",
					"enable_ave",
					"enable_tag",
					"encap_mode",
					"enf_pref",
					"ep_inventory_type",
					"ep_ret_time",
					"hv_avail_monitor",
					"mcast_addr",
					"mode",
					"pref_encap_mode",
					"relation_vmm_rs_pref_enhanced_lag_pol",
					"relation_infra_rs_vlan_ns",
					"relation_vmm_rs_dom_mcast_addr_ns",
					"relation_vmm_rs_default_cdp_if_pol",
					"relation_vmm_rs_default_lacp_lag_pol",
					"relation_infra_rs_vlan_ns_def",
					"relation_infra_rs_vip_addr_ns",
					"relation_vmm_rs_default_lldp_if_pol",
					"relation_vmm_rs_default_stp_if_pol",
					"relation_infra_rs_dom_vxlan_ns_def",
					"relation_vmm_rs_default_fw_pol",
					"relation_vmm_rs_default_l2_inst_pol",
					"config_infra_pg",
					"ave_time_out",
					"arp_learning",
					"access_mode",
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
