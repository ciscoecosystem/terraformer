package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutsideClass = "l3extOut"

type L3OutsideGenerator struct {
	ACIService
}

func (a *L3OutsideGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutsideClass)

	l3OutCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	l3OutCount, err := strconv.Atoi(stripQuotes(l3OutCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3OutCount; i++ {
		l3OutDN := stripQuotes(l3OutCont.S("imdata").Index(i).S(L3OutsideClass, "attributes", "dn").String())
		if filterChildrenDn(l3OutDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				l3OutDN,
				resourceNamefromDn(L3OutsideClass, (l3OutDN), i),
				"aci_l3_outside",
				"aci",
				[]string{
					"enforce_rtctrl",
					"name_alias",
					"target_dscp",
					"relation_l3ext_rs_dampening_pol",
					"relation_l3ext_rs_ectx",
					"relation_l3ext_rs_out_to_bd_public_subnet_holder",
					"relation_l3ext_rs_interleak_pol",
					"relation_l3ext_rs_l3_dom_att",
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