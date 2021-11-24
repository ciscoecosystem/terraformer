package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fexBundleGrpClass = "infraFexBndlGrp"

type FexBundleGrpGenerator struct {
	ACIService
}

func (a *FexBundleGrpGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fexBundleGrpClass)

	fexBundleGrpCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(fexBundleGrpCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	fexBundleGrpCount, err := strconv.Atoi(stripQuotes(fexBundleGrpCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < fexBundleGrpCount; i++ {
		fexBundleGrpProfileDN := stripQuotes(fexBundleGrpCont.S("imdata").Index(i).S(fexBundleGrpClass, "attributes", "dn").String())
		if filterChildrenDn(fexBundleGrpProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				fexBundleGrpProfileDN,
				resourceNamefromDn(fexBundleGrpClass, (fexBundleGrpProfileDN), i),
				"aci_fex_bundle_group",
				"aci",
				[]string{
					"name",
					"name_alias",
					"relation_infra_rs_mon_fex_infra_pol",
					"relation_infra_rs_fex_bndl_grp_to_aggr_if",
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
