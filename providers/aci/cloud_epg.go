package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cloudEPGClass = "cloudEPg"

type CloudEPGGenerator struct {
	ACIService
}

func (a *CloudEPGGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, cloudEPGClass)

	cloudEPGCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(cloudEPGCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	cloudEPGCount, err := strconv.Atoi(stripQuotes(cloudEPGCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < cloudEPGCount; i++ {
		cloudEPGProfileDN := stripQuotes(cloudEPGCont.S("imdata").Index(i).S(cloudEPGClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			cloudEPGProfileDN,
			cloudEPGProfileDN,
			"aci_cloud_epg",
			"aci",
			[]string{
				"exception_tag",
				"flood_on_encap",
				"match_t",
				"name_alias",
				"pref_gr_memb",
				"prio",
				"relation_fv_rs_sec_inherited",
				"relation_fv_rs_prov",
				"relation_fv_rs_cons_if",
				"relation_fv_rs_cust_qos_pol",
				"relation_fv_rs_cons",
				"relation_cloud_rs_cloud_epg_ctx",
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
