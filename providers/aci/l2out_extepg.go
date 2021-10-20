package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l2OutExtEPGClass = "l2extInstP"

type L2OutExtEPGGenerator struct {
	ACIService
}

func (a *L2OutExtEPGGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l2OutExtEPGClass)

	l2OutExtEPGCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(l2OutExtEPGCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	l2OutExtEPGCount, err := strconv.Atoi(stripQuotes(l2OutExtEPGCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l2OutExtEPGCount; i++ {
		l2OutExtEPGProfileDN := stripQuotes(l2OutExtEPGCont.S("imdata").Index(i).S(l2OutExtEPGClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			l2OutExtEPGProfileDN,
			l2OutExtEPGProfileDN,
			"aci_l2out_extepg",
			"aci",
			[]string{
				"name_alias",
				"exception_tag",
				"flood_on_encap",
				"match_t",
				"pref_gr_memb",
				"prio",
				"target_dscp",
				"relation_fv_rs_sec_inherited",
				"relation_fv_rs_prov",
				"relation_fv_rs_cons_if",
				"relation_fv_rs_cust_qos_pol",
				"relation_fv_rs_cons",
				"relation_l2ext_rs_l2_inst_p_to_dom_p",
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