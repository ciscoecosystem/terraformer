package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const matchRuleClassName = "rtctrlSubjP"

type MatchRuleGenerator struct {
	ACIService
}

func (a *MatchRuleGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, matchRuleClassName)

	MatchRuleCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MatchRuleCount, err := strconv.Atoi(stripQuotes(MatchRuleCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MatchRuleCount; i++ {
		MatchRuleAttr := MatchRuleCont.S("imdata").Index(i).S(matchRuleClassName, "attributes")
		MatchRuleDN := G(MatchRuleAttr, "dn")
		name := G(MatchRuleAttr, "name")
		if filterChildrenDn(MatchRuleDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				MatchRuleDN,
				resourceNamefromDn(matchRuleClassName, MatchRuleDN, i),
				"aci_match_rule",
				"aci",
				map[string]string{
					"tenant_dn": GetParentDn(MatchRuleDN, fmt.Sprintf("/subj-%s", name)),
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
