package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l3OutStaticRouteClass = "ipRouteP"

type L3OutStaticRouteGenerator struct {
	ACIService
}

func (a *L3OutStaticRouteGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l3OutStaticRouteClass)

	l3StaticRouteCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(l3StaticRouteCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	l3StaticRouteCount, err := strconv.Atoi(stripQuotes(l3StaticRouteCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3StaticRouteCount; i++ {
		l3StaticRouteProfileDN := stripQuotes(l3StaticRouteCont.S("imdata").Index(i).S(l3OutStaticRouteClass, "attributes", "dn").String())
		if filterChildrenDn(l3StaticRouteProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
			l3StaticRouteProfileDN,
			l3StaticRouteProfileDN,
			"aci_l3out_static_route",
			"aci",
			[]string{
				"aggregate",
				"name_alias",
				"pref",
				"rt_ctrl",
				"relation_ip_rs_route_track",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}}

	return nil
}
