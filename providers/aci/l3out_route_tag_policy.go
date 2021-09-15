package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutRouteTagPolicyClassName = "l3extRouteTagPol"

type L3OutRouteTagPolicyGenerator struct {
	ACIService
}

func (a *L3OutRouteTagPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutRouteTagPolicyClassName)

	L3OutRouteTagPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutRouteTagPolicyCount, err := strconv.Atoi(stripQuotes(L3OutRouteTagPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L3OutRouteTagPolicyCount; i++ {
		L3OutRouteTagPolicyDN := stripQuotes(L3OutRouteTagPolicyCont.S("imdata").Index(i).S(L3OutRouteTagPolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(L3OutRouteTagPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutRouteTagPolicyDN,
				L3OutRouteTagPolicyDN,
				"aci_l3out_route_tag_policy",
				"aci",
				[]string{
					"name_alias",
					"tag",
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
