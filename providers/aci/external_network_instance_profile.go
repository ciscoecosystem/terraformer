package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const ExtNetInsProClass = "l3extInstP"

type ExtNetInsProGenerator struct {
	ACIService
}

func (a *ExtNetInsProGenerator) InitResources() error {
	client, err := a.createClient()

	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, ExtNetInsProClass)

	extNetInsProCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	extNetInsProCount, err := strconv.Atoi(stripQuotes(extNetInsProCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < extNetInsProCount; i++ {
		extNetInsProDN := stripQuotes(extNetInsProCont.S("imdata").Index(i).S(ExtNetInsProClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			extNetInsProDN,
			extNetInsProDN,
			"aci_external_network_instance_profile",
			"aci",
			[]string{
				"exception_tag",
				"flood_on_encap",
				"match_t",
				"name_alias",
				"pref_gr_memb",
				"prio",
				"target_dscp",
				"relation_fv_rs_sec_inherited",
				"relation_fv_rs_prov",
				"relation_l3ext_rs_l3_inst_p_to_dom_p",
				"relation_l3ext_rs_inst_p_to_nat_mapping_epg",
				"relation_fv_rs_cons_if",
				"relation_fv_rs_cust_qos_pol",
				"relation_l3ext_rs_inst_p_to_profile",
				"relation_fv_rs_cons",
				"relation_fv_rs_prot_by",
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
