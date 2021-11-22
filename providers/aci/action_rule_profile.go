package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const actionRuleProfClass = "rtctrlAttrP"

type ActionRuleProfGenerator struct {
	ACIService
}

func (a *ActionRuleProfGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, actionRuleProfClass)

	actionRuleProfCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(actionRuleProfCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	actionRuleProfCount, err := strconv.Atoi(stripQuotes(actionRuleProfCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < actionRuleProfCount; i++ {
		actionRuleProfProfileDN := stripQuotes(actionRuleProfCont.S("imdata").Index(i).S(actionRuleProfClass, "attributes", "dn").String())
		if filterChildrenDn(actionRuleProfProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				actionRuleProfProfileDN,
				fmt.Sprintf("%s_%s_%d", actionRuleProfClass, GetMOName(actionRuleProfProfileDN), i),
				"aci_action_rule_profile",
				"aci",
				[]string{
					"name_alias",
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
