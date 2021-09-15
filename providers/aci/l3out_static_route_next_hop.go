package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l3OutStaticRouteNextHopClass = "ipNexthopP"

type L3OutStaticRouteNextHopGenerator struct {
	ACIService
}

func (a *L3OutStaticRouteNextHopGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l3OutStaticRouteNextHopClass)

	l3StaticRouteNextHopCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	l3StaticRouteNexHopCount, err := strconv.Atoi(stripQuotes(l3StaticRouteNextHopCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3StaticRouteNexHopCount; i++ {
		l3StaticRouteNextHopProfileDN := stripQuotes(l3StaticRouteNextHopCont.S("imdata").Index(i).S(l3OutStaticRouteNextHopClass, "attributes", "dn").String())
		if filterChildrenDn(l3StaticRouteNextHopProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				l3StaticRouteNextHopProfileDN,
				l3StaticRouteNextHopProfileDN,
				"aci_l3out_static_route_next_hop",
				"aci",
				[]string{
					"nexthop_profile_type",
					"name_alias",
					"pref",
					"relation_ip_rs_nexthop_route_track",
					"relation_ip_rs_nh_track_member",
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
