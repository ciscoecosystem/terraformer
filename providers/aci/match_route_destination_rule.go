package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const matchRouteDestinationRuleClassName = "rtctrlMatchRtDest"

type MatchRouteDestinationRuleGenerator struct {
	ACIService
}

func (a *MatchRouteDestinationRuleGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, matchRouteDestinationRuleClassName)

	MatchRouteDestinationRuleCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MatchRouteDestinationRuleCount, err := strconv.Atoi(stripQuotes(MatchRouteDestinationRuleCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MatchRouteDestinationRuleCount; i++ {
		MatchRouteDestinationRuleAttr := MatchRouteDestinationRuleCont.S("imdata").Index(i).S(matchRouteDestinationRuleClassName, "attributes")
		MatchRouteDestinationRuleDN := G(MatchRouteDestinationRuleAttr,"dn")
		ip := G(MatchRouteDestinationRuleAttr,"ip")
		if filterChildrenDn(MatchRouteDestinationRuleDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					MatchRouteDestinationRuleDN,
					resourceNamefromDn(matchRouteDestinationRuleClassName,MatchRouteDestinationRuleDN,i),
					"aci_match_route_destination_rule",
					"aci",
					map[string]string{
						"match_rule_dn": GetParentDn(MatchRouteDestinationRuleDN, fmt.Sprintf("/dest-[%s]", ip)),
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