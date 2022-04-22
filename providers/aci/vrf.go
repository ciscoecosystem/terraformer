package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const VRFClass = "fvCtx"

type VRFGenerator struct {
	ACIService
}

func (a *VRFGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, VRFClass)

	vrfCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	vrfCount, err := strconv.Atoi(stripQuotes(vrfCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < vrfCount; i++ {
		vrfDN := stripQuotes(vrfCont.S("imdata").Index(i).S(VRFClass, "attributes", "dn").String())
		if filterChildrenDn(vrfDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				vrfDN,
				resourceNamefromDn(VRFClass, (vrfDN), i),
				"aci_vrf",
				"aci",
				[]string{
					"bd_enforced_enable",
					"ip_data_plane_learning",
					"knw_mcast_act",
					"name_alias",
					"pc_enf_dir",
					"pc_enf_pref",
					"relation_fv_rs_ospf_ctx_pol",
					"relation_fv_rs_vrf_validation_pol",
					"relation_fv_rs_ctx_mcast_to",
					"relation_fv_rs_ctx_to_eigrp_ctx_af_pol",
					"relation_fv_rs_ctx_to_ospf_ctx_pol",
					"relation_fv_rs_ctx_to_ep_ret",
					"relation_fv_rs_bgp_ctx_pol",
					"relation_fv_rs_ctx_mon_pol",
					"relation_fv_rs_ctx_to_ext_route_tag_pol",
					"relation_fv_rs_ctx_to_bgp_ctx_af_pol",
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
