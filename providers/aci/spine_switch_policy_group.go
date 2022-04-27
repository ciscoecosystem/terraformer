package aci

import (
	"fmt"
	"strconv"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const spineSwitchPolicyGroupClassName = "infraSpineAccNodePGrp"

type SpineSwitchPolicyGroupGenerator struct {
	ACIService
}
func (a *SpineSwitchPolicyGroupGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, spineSwitchPolicyGroupClassName)
	SpineSwitchPolicyGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	SpineSwitchPolicyGroupCount, err := strconv.Atoi(stripQuotes(SpineSwitchPolicyGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < SpineSwitchPolicyGroupCount; i++ {
		SpineSwitchPolicyGroupAttr := SpineSwitchPolicyGroupCont.S("imdata").Index(i).S(spineSwitchPolicyGroupClassName, "attributes")
		SpineSwitchPolicyGroupDN := G(SpineSwitchPolicyGroupAttr,"dn")
		name := G(SpineSwitchPolicyGroupAttr,"name")
		if filterChildrenDn(SpineSwitchPolicyGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					SpineSwitchPolicyGroupDN,
					name,
					"aci_spine_switch_policy_group",
					"aci",
					map[string]string{
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