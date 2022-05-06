package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const matchRuleBasedonCommunityRegularExpressionClassName = "rtctrlMatchCommRegexTerm"

type MatchRuleBasedonCommunityRegularExpressionGenerator struct {
	ACIService
}

func (a *MatchRuleBasedonCommunityRegularExpressionGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, matchRuleBasedonCommunityRegularExpressionClassName)

	MatchRuleBasedonCommunityRegularExpressionCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MatchRuleBasedonCommunityRegularExpressionCount, err := strconv.Atoi(stripQuotes(MatchRuleBasedonCommunityRegularExpressionCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MatchRuleBasedonCommunityRegularExpressionCount; i++ {
		MatchRuleBasedonCommunityRegularExpressionAttr := MatchRuleBasedonCommunityRegularExpressionCont.S("imdata").Index(i).S(matchRuleBasedonCommunityRegularExpressionClassName, "attributes")
		MatchRuleBasedonCommunityRegularExpressionDN := G(MatchRuleBasedonCommunityRegularExpressionAttr,"dn")
		commType := G(MatchRuleBasedonCommunityRegularExpressionAttr,"commType")
		if filterChildrenDn(MatchRuleBasedonCommunityRegularExpressionDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					MatchRuleBasedonCommunityRegularExpressionDN,
					resourceNamefromDn(matchRuleBasedonCommunityRegularExpressionClassName,MatchRuleBasedonCommunityRegularExpressionDN,i),
					"aci_match_regex_community_terms",
					"aci",
					map[string]string{
						"match_rule_dn": GetParentDn(MatchRuleBasedonCommunityRegularExpressionDN, fmt.Sprintf("/commrxtrm-%s", commType)),
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