package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BgpBestPathPolicyClass = "bgpBestPathCtrlPol"

type BgpBestPathPolicyGenerator struct {
	ACIService
}

func (a *BgpBestPathPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BgpBestPathPolicyClass)

	BgpBestPathPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BgpBestPathPolicyCount, err := strconv.Atoi(stripQuotes(BgpBestPathPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BgpBestPathPolicyCount; i++ {
		BgpBestPathPolicyDN := stripQuotes(BgpBestPathPolicyCont.S("imdata").Index(i).S(BgpBestPathPolicyClass, "attributes", "dn").String())
		if filterChildrenDn(BgpBestPathPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				BgpBestPathPolicyDN,
				fmt.Sprintf("%s_%s_%d", BgpBestPathPolicyClass, GetMOName(BgpBestPathPolicyDN), i),
				"aci_bgp_best_path_policy",
				"aci",
				[]string{
					"name_alias",
					"annotation",
					"description",
					"ctrl",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
