package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l2DomClass = "l2extDomP"

type L2DomGenerator struct {
	ACIService
}

func (a *L2DomGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l2DomClass)

	l2DomCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(l2DomCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	l2DomCount, err := strconv.Atoi(stripQuotes(l2DomCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l2DomCount; i++ {
		l2DomProfileDN := stripQuotes(l2DomCont.S("imdata").Index(i).S(l2DomClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			l2DomProfileDN,
			l2DomProfileDN,
			"aci_l2_domain",
			"aci",
			[]string{
				"name_alias",
				"relation_infra_rs_vlan_ns",
				"relation_infra_rs_vlan_ns_def",
				"relation_infra_rs_vip_addr_ns",
				"relation_extnw_rs_out",
				"relation_infra_rs_dom_vxlan_ns_def",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}