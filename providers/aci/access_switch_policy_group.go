package aci

import (
	"fmt"
	"strconv"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessSwitchPolicyGroupClassName = "infraAccNodePGrp"

type AccessSwitchPolicyGroupGenerator struct {
	ACIService
}

func (a *AccessSwitchPolicyGroupGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessSwitchPolicyGroupClassName)
	AccessSwitchPolicyGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	AccessSwitchPolicyGroupCount, err := strconv.Atoi(stripQuotes(AccessSwitchPolicyGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < AccessSwitchPolicyGroupCount; i++ {
		AccessSwitchPolicyGroupAttr := AccessSwitchPolicyGroupCont.S("imdata").Index(i).S(accessSwitchPolicyGroupClassName, "attributes")
		AccessSwitchPolicyGroupDN := G(AccessSwitchPolicyGroupAttr, "dn")
		name := G(AccessSwitchPolicyGroupAttr, "name")
		if filterChildrenDn(AccessSwitchPolicyGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				AccessSwitchPolicyGroupDN,
				name,
				"aci_access_switch_policy_group",
				"aci",
				map[string]string{},
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
