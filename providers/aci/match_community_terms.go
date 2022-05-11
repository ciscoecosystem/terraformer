package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const matchCommunityTermClassName = "rtctrlMatchCommTerm"

type MatchCommunityTermGenerator struct {
	ACIService
}

func (a *MatchCommunityTermGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, matchCommunityTermClassName)
	MatchCommunityTermCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	MatchCommunityTermCount, err := strconv.Atoi(stripQuotes(MatchCommunityTermCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MatchCommunityTermCount; i++ {
		MatchCommunityTermAttr := MatchCommunityTermCont.S("imdata").Index(i).S(matchCommunityTermClassName, "attributes")
		MatchCommunityTermDN := G(MatchCommunityTermAttr, "dn")
		name := G(MatchCommunityTermAttr, "name")
		if filterChildrenDn(MatchCommunityTermDN, client.parentResource) != "" {
			matchMap, err := getMatchMap(client, MatchCommunityTermDN)
			if err != nil {
				return err
			}
			var resource terraformutils.Resource
			if len(matchMap) > 0 {
				resource = terraformutils.NewResource(
					MatchCommunityTermDN,
					resourceNamefromDn(matchCommunityTermClassName, MatchCommunityTermDN, i),
					"aci_match_community_terms",
					"aci",
					map[string]string{
						"match_rule_dn": GetParentDn(MatchCommunityTermDN, fmt.Sprintf("/commtrm-%s", name)),
					},
					[]string{},
					map[string]interface{}{
						"match_community_factors": matchMap,
					},
				)
			} else {
				resource = terraformutils.NewResource(
					MatchCommunityTermDN,
					resourceNamefromDn(matchCommunityTermClassName, MatchCommunityTermDN, i),
					"aci_match_community_terms",
					"aci",
					map[string]string{
						"match_rule_dn": GetParentDn(MatchCommunityTermDN, fmt.Sprintf("/commtrm-%s", name)),
					},
					[]string{},
					map[string]interface{}{},
				)
			}

			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}

func getMatchMap(c *ACIClient, parentDN string) ([]map[string]string, error) {
	cont, err := c.GetViaURL("api/node/class/rtctrlMatchCommFactor.json")
	if err != nil {
		return nil, err
	}
	matchValues := make([]map[string]string, 0, 1)
	matchData := cont.S("imdata")
	for i := 0; i < len(matchData.Data().([]interface{})); i++ {
		matchCont := matchData.Index(i)
		matchContOut := matchCont.S("rtctrlMatchCommFactor")
		matchAttrCont := matchContOut.S("attributes")
		factorDN := G(matchAttrCont, "dn")
		if filterChildrenDn(factorDN, parentDN) != "" {
			community := G(matchAttrCont, "community")
			scope := G(matchAttrCont, "scope")
			desc := G(matchAttrCont, "descr")
			matchValue := make(map[string]string, 0)
			matchValue["community"] = community
			matchValue["scope"] = scope
			if desc != "" {
				matchValue["description"] = desc
			}
			matchValues = append(matchValues, matchValue)
		}

	}
	return matchValues, nil
}
