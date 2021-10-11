package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const leafAccBunPolGClass = "infraAccBndlGrp"

type LeafAccBunPolGGenerator struct {
	ACIService
}

func (a *LeafAccBunPolGGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, leafAccBunPolGClass)

	leafAccBunPolGCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(leafAccBunPolGCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	leafAccBunPolGCount, err := strconv.Atoi(stripQuotes(leafAccBunPolGCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < leafAccBunPolGCount; i++ {
		leafAccBunPolGProfileDN := stripQuotes(leafAccBunPolGCont.S("imdata").Index(i).S(leafAccBunPolGClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			leafAccBunPolGProfileDN,
			leafAccBunPolGProfileDN,
			"aci_leaf_access_bundle_policy_group",
			"aci",
			[]string{
				"lag_t",
				"name_alias",
				"relation_infra_rs_span_v_src_grp",
				"relation_infra_rs_stormctrl_if_pol",
				"relation_infra_rs_lldp_if_pol",
				"relation_infra_rs_macsec_if_pol",
				"relation_infra_rs_qos_dpp_if_pol",
				"relation_infra_rs_h_if_pol",
				"relation_infra_rs_netflow_monitor_pol",
				"relation_infra_rs_l2_port_auth_pol",
				"relation_infra_rs_mcp_if_pol",
				"relation_infra_rs_l2_port_security_pol",
				"relation_infra_rs_copp_if_pol",
				"relation_infra_rs_span_v_dest_grp",
				"relation_infra_rs_lacp_pol",
				"relation_infra_rs_cdp_if_pol",
				"relation_infra_rs_qos_pfc_if_pol",
				"relation_infra_rs_qos_sd_if_pol",
				"relation_infra_rs_mon_if_infra_pol",
				"relation_infra_rs_fc_if_pol",
				"relation_infra_rs_qos_ingress_dpp_if_pol",
				"relation_infra_rs_qos_egress_dpp_if_pol",
				"relation_infra_rs_l2_if_pol",
				"relation_infra_rs_stp_if_pol",
				"relation_infra_rs_att_ent_p",
				"relation_infra_rs_l2_inst_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}