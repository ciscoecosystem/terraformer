package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const applicationEPGClass = "fvAEPg"

type ApplicationEPGGenerator struct {
	ACIService
}

func (a *ApplicationEPGGenerator) InitResources() error {

	client, err := a.createClient()

	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, applicationEPGClass)

	epgCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	epgCount, err := strconv.Atoi(stripQuotes(epgCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < epgCount; i++ {
		epgProfileDN := stripQuotes(epgCont.S("imdata").Index(i).S(applicationEPGClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			epgProfileDN,
			epgProfileDN,
			"aci_application_epg",
			"aci",
			[]string{
				"exception_tag",
				"flood_on_encap",
				"fwd_ctrl",
				"has_mcast_source",
				"is_attr_based_epg",
				"match_t",
				"name_alias",
				"pc_enf_pref",
				"pref_gr_memb",
				"prio",
				"shutdown",
				"relation_fv_rs_bd",
				"relation_fv_rs_cust_qos_pol",
				"relation_fv_rs_fc_path_att",
				"relation_fv_rs_prov",
				"relation_fv_rs_graph_def",
				"relation_fv_rs_cons_if",
				"relation_fv_rs_sec_inherited",
				"relation_fv_rs_node_att",
				"relation_fv_rs_dpp_pol",
				"relation_fv_rs_cons",
				"relation_fv_rs_prov_def",
				"relation_fv_rs_trust_ctrl",
				"relation_fv_rs_prot_by",
				"relation_fv_rs_aepg_mon_pol",
				"relation_fv_rs_intra_epg",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
