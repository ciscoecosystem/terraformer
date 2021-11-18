package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L2OutsideClassName = "l2extOut"

type L2OutsideGenerator struct {
	ACIService
}

func (a *L2OutsideGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L2OutsideClassName)

	L2OutsideCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L2OutsideCount, err := strconv.Atoi(stripQuotes(L2OutsideCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L2OutsideCount; i++ {
		L2OutsideDN := L2OutsideCont.S("imdata").Index(i).S(L2OutsideClassName, "attributes", "dn").String()
		if filterChildrenDn(L2OutsideDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(L2OutsideDN),
				stripQuotes(L2OutsideDN),
				"aci_l2_outside",
				"aci",
				[]string{
					"name_alias",
					"target_dscp",
					"relation_l2ext_rs_e_bd",
					"relation_l2ext_rs_l2_dom_att",
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
