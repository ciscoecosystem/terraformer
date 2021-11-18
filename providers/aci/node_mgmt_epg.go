package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const NodeMgmtEPGClassName = "mgmtInB"

type NodeMgmtEPGGenerator struct {
	ACIService
}

func (a *NodeMgmtEPGGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, NodeMgmtEPGClassName)

	NodeMgmtEPGCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	NodeMgmtEPGCount, err := strconv.Atoi(stripQuotes(NodeMgmtEPGCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < NodeMgmtEPGCount; i++ {
		NodeMgmtEPGDN := NodeMgmtEPGCont.S("imdata").Index(i).S(NodeMgmtEPGClassName, "attributes", "dn").String()
		if filterChildrenDn(NodeMgmtEPGDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				stripQuotes(NodeMgmtEPGDN),
				stripQuotes(NodeMgmtEPGDN),
				"aci_node_mgmt_epg",
				"aci",
				map[string]string{
					"type": "in_band",
				},
				[]string{
					"name_alias",
					"prio",
					"encap",
					"exception_tag",
					"flood_on_encap",
					"match_t",
					"pref_gr_memb",
					"relation_fv_rs_sec_inherited",
					"relation_fv_rs_prov",
					"relation_fv_rs_cons_if",
					"relation_fv_rs_cust_qos_pol",
					"relation_mgmt_rs_mgmt_bd",
					"relation_fv_rs_cons",
					"relation_fv_rs_prot_by",
					"relation_mgmt_rs_in_b_st_node",
					"relation_fv_rs_intra_epg",
					"relation_mgmt_rs_oo_b_prov",
					"relation_mgmt_rs_oo_b_st_node",
					"relation_mgmt_rs_oo_b_ctx",
					"annotation",
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	dnURL = fmt.Sprintf("%s/%s.json", baseURL, "mgmtOoB")

	NodeMgmtEPGCont, err = client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	NodeMgmtEPGCount, err = strconv.Atoi(stripQuotes(NodeMgmtEPGCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < NodeMgmtEPGCount; i++ {
		NodeMgmtEPGDN := NodeMgmtEPGCont.S("imdata").Index(i).S("mgmtOoB", "attributes", "dn").String()
		if filterChildrenDn(NodeMgmtEPGDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				stripQuotes(NodeMgmtEPGDN),
				stripQuotes(NodeMgmtEPGDN),
				"aci_node_mgmt_epg",
				"aci",
				map[string]string{
					"type": "out_of_band",
				},
				[]string{
					"management_profile_dn",
					"name_alias",
					"prio",
					"encap",
					"exception_tag",
					"flood_on_encap",
					"match_t",
					"pref_gr_memb",
					"relation_fv_rs_sec_inherited",
					"relation_fv_rs_prov",
					"relation_fv_rs_cons_if",
					"relation_fv_rs_cust_qos_pol",
					"relation_mgmt_rs_mgmt_bd",
					"relation_fv_rs_cons",
					"relation_fv_rs_prot_by",
					"relation_mgmt_rs_in_b_st_node",
					"relation_fv_rs_intra_epg",
					"relation_mgmt_rs_oo_b_prov",
					"relation_mgmt_rs_oo_b_st_node",
					"relation_mgmt_rs_oo_b_ctx",
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
