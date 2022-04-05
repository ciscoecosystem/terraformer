package mso

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DhcpOptionPolicyOptionGenerator struct {
	MSOService
}

func (a *DhcpOptionPolicyOptionGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getDhcpContainer(mso)
	if err != nil {
		return err
	}
	dhcpLength := len(con.S("DhcpRelayPolicies").Data().([]interface{}))
	for i := 0; i < dhcpLength; i++ {
		dhcpCon := con.S("DhcpRelayPolicies").Index(i)
		optionPolicyName := stripQuotes(dhcpCon.S("name").String())
		dhcpOptionCon := dhcpCon.S("dhcpOption")
		dhcpOptionCount := 0
		if dhcpCon.Exists("dhcpOption") {
			dhcpOptionCount = len(dhcpOptionCon.Data().([]interface{}))
		}
		for j := 0; j < dhcpOptionCount; j++ {
			optionName := stripQuotes(dhcpOptionCon.Index(j).S("name").String())
			dhcpOptionId := fmt.Sprintf("%s/%s", optionPolicyName, optionName)
			dhcpOptionName := fmt.Sprintf("%s_%s", optionPolicyName, optionName)
			resource := terraformutils.NewResource(
				dhcpOptionId,
				dhcpOptionName,
				"mso_dhcp_option_policy_option",
				"mso",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
