package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BgpRouteControlProfileClass = "rtctrlProfile"

type BgpRouteControlProfileGenerator struct {
	ACIService
}

func (a *BgpRouteControlProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client:= clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BgpRouteControlProfileClass)

	BgpRouteControlProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BgpRouteControlProfileCount, err := strconv.Atoi(stripQuotes(BgpRouteControlProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BgpRouteControlProfileCount; i++ {
		BgpRouteControlProfileDN := stripQuotes(BgpRouteControlProfileCont.S("imdata").Index(i).S(BgpRouteControlProfileClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			BgpRouteControlProfileDN,
			BgpRouteControlProfileDN,
			"aci_bgp_route_control_profile",
			"aci",
			[]string{
				"name_alias",
				"annotation",
				"description",
				"route_control_profile_type",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
