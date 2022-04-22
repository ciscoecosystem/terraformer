package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutFloatingSviClassName = "l3extVirtualLIfP"

type L3OutFloatingSviGenerator struct {
	ACIService
}

func (a *L3OutFloatingSviGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutFloatingSviClassName)

	L3OutFloatingSviCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutFloatingSviCount, err := strconv.Atoi(stripQuotes(L3OutFloatingSviCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L3OutFloatingSviCount; i++ {
		L3OutFloatingSviDN := stripQuotes(L3OutFloatingSviCont.S("imdata").Index(i).S(L3OutFloatingSviClassName, "attributes", "dn").String())
		if filterChildrenDn(L3OutFloatingSviDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutFloatingSviDN,
				resourceNamefromDn(L3OutFloatingSviClassName, L3OutFloatingSviDN, i),
				"aci_l3out_floating_svi",
				"aci",
				[]string{
					"addr",
					"autostate",
					"encap_scope",
					"if_inst_t",
					"ipv6_dad",
					"ll_addr",
					"mac",
					"mode",
					"mtu",
					"target_dscp",
					"relation_l3ext_rs_dyn_path_att",
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
