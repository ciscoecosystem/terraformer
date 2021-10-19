package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpinePortPolicyGroupClass = "infraSpAccPortGrp"

type SpinePortPolicyGroupGenerator struct {
	ACIService
}

func (a *SpinePortPolicyGroupGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpinePortPolicyGroupClass)

	SpinePortPolicyGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SpinePortPolicyGroupCount, err := strconv.Atoi(stripQuotes(SpinePortPolicyGroupCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < SpinePortPolicyGroupCount; i++ {
		SpinePortPolicyGroupProfileDN := stripQuotes(SpinePortPolicyGroupCont.S("imdata").Index(i).S(SpinePortPolicyGroupClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			SpinePortPolicyGroupProfileDN,
			SpinePortPolicyGroupProfileDN,
			"aci_spine_port_policy_group",
			"aci",
			[]string{
				"relation_infra_rs_h_if_pol",
				"relation_infra_rs_cdp_if_pol",
				"relation_infra_rs_copp_if_pol",
				"relation_infra_rs_att_ent_p",
				"relation_infra_rs_macsec_if_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
