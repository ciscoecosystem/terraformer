package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutBGPProtocolProfileClassName = "bgpProtP"

type L3OutBGPProtocolProfileGenerator struct {
	ACIService
}

func (a *L3OutBGPProtocolProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutBGPProtocolProfileClassName)

	L3OutBGPProtocolProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutBGPProtocolProfileCount, err := strconv.Atoi(stripQuotes(L3OutBGPProtocolProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L3OutBGPProtocolProfileCount; i++ {
		L3OutBGPProtocolProfileDN := stripQuotes(L3OutBGPProtocolProfileCont.S("imdata").Index(i).S(L3OutBGPProtocolProfileClassName, "attributes", "dn").String())
		if filterChildrenDn(L3OutBGPProtocolProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutBGPProtocolProfileDN,
				resourceNamefromDn(L3OutBGPProtocolProfileClassName, (L3OutBGPProtocolProfileDN), i),
				"aci_l3out_bgp_protocol_profile",
				"aci",
				[]string{
					"relation_bgp_rs_bgp_node_ctx_pol",
					"name_alias",
					"annotation",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
