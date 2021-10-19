package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l3DomPClass = "l3extDomP"

type L3DomPGenerator struct {
	ACIService
}

func (a *L3DomPGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l3DomPClass)

	l3DomPCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(l3DomPCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	l3DomPCount, err := strconv.Atoi(stripQuotes(l3DomPCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3DomPCount; i++ {
		l3DomPProfileDN := stripQuotes(l3DomPCont.S("imdata").Index(i).S(l3DomPClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			l3DomPProfileDN,
			l3DomPProfileDN,
			"aci_l3_domain_profile",
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