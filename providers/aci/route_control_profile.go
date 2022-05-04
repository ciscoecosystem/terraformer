package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const routeControlProfileClassName = "rtctrlProfile"

type RouteControlProfileGenerator struct {
	ACIService
}

func (a *RouteControlProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, routeControlProfileClassName)

	RouteControlProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RouteControlProfileCount, err := strconv.Atoi(stripQuotes(RouteControlProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RouteControlProfileCount; i++ {
		RouteControlProfileAttr := RouteControlProfileCont.S("imdata").Index(i).S(routeControlProfileClassName, "attributes")
		RouteControlProfileDN := G(RouteControlProfileAttr, "dn")
		name := G(RouteControlProfileAttr, "name")
		if filterChildrenDn(RouteControlProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RouteControlProfileDN,
				resourceNamefromDn(routeControlProfileClassName, RouteControlProfileDN, i),
				"aci_route_control_profile",
				"aci",
				map[string]string{
					"parent_dn": GetParentDn(RouteControlProfileDN, fmt.Sprintf("/prof-%s", name)),
				},
				[]string{
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
