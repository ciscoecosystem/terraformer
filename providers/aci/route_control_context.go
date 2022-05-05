package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const routeControlContextClassName = "rtctrlCtxP"

type RouteControlContextGenerator struct {
	ACIService
}

func (a *RouteControlContextGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, routeControlContextClassName)

	RouteControlContextCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RouteControlContextCount, err := strconv.Atoi(stripQuotes(RouteControlContextCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RouteControlContextCount; i++ {
		RouteControlContextAttr := RouteControlContextCont.S("imdata").Index(i).S(routeControlContextClassName, "attributes")
		RouteControlContextDN := G(RouteControlContextAttr, "dn")
		name := G(RouteControlContextAttr, "name")
		if filterChildrenDn(RouteControlContextDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RouteControlContextDN,
				resourceNamefromDn(routeControlContextClassName, RouteControlContextDN, i),
				"aci_route_control_context",
				"aci",
				map[string]string{
					"route_control_profile_dn": GetParentDn(RouteControlContextDN, fmt.Sprintf("/ctx-%s", name)),
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
