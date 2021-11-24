package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BgpAddressFamilyContextClass = "bgpCtxAfPol"

type BgpAddressFamilyContextGenerator struct {
	ACIService
}

func (a *BgpAddressFamilyContextGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BgpAddressFamilyContextClass)

	BgpAddressFamilyContextCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BgpAddressFamilyContextCount, err := strconv.Atoi(stripQuotes(BgpAddressFamilyContextCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BgpAddressFamilyContextCount; i++ {
		BgpAddressFamilyContextDN := stripQuotes(BgpAddressFamilyContextCont.S("imdata").Index(i).S(BgpAddressFamilyContextClass, "attributes", "dn").String())
		if filterChildrenDn(BgpAddressFamilyContextDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				BgpAddressFamilyContextDN,
				resourceNamefromDn(BgpAddressFamilyContextClass, (BgpAddressFamilyContextDN), i),
				"aci_bgp_address_family_context",
				"aci",
				[]string{
					"name_alias",
					"annotation",
					"description",
					"e_dist",
					"i_dist",
					"local_dist",
					"max_ecmp",
					"max_ecmp_ibgp",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
