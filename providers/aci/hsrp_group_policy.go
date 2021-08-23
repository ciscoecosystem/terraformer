package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const HSRPGroupPolicyClassName = "hsrpGroupPol"

type HSRPGroupPolicyGenerator struct {
	ACIService
}

func (a *HSRPGroupPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, HSRPGroupPolicyClassName)

	HSRPGroupPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	HSRPGroupPolicyCount, err := strconv.Atoi(stripQuotes(HSRPGroupPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < HSRPGroupPolicyCount; i++ {
		HSRPGroupPolicyDN := stripQuotes(HSRPGroupPolicyCont.S("imdata").Index(i).S(HSRPGroupPolicyClassName, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			HSRPGroupPolicyDN,
			HSRPGroupPolicyDN,
			"aci_hsrp_group_policy",
			"aci",
			[]string{
				"ctrl",
				"hello_intvl",
				"hold_intvl",
				"key",
				"name_alias",
				"preempt_delay_min",
				"preempt_delay_reload",
				"preempt_delay_sync",
				"prio",
				"timeout",
				"hsrp_group_policy_type",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
