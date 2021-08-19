package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const DHCPRelayPolicyClass = "dhcpRelayP"

type DHCPRelayPolicyGenerator struct {
	ACIService
}

func (a *DHCPRelayPolicyGenerator) InitResources() error {
	client, err := a.createClient()

	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, DHCPRelayPolicyClass)

	DHCPRelayPolicyCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	DHCPRelayPolicyCount, err := strconv.Atoi(stripQuotes(DHCPRelayPolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < DHCPRelayPolicyCount; i++ {
		DHCPRelayPolicyDN := stripQuotes(DHCPRelayPolicyCont.S("imdata").Index(i).S(DHCPRelayPolicyClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			DHCPRelayPolicyDN,
			DHCPRelayPolicyDN,
			"aci_dhcp_relay_policy",
			"aci",
			[]string{
				"annotation",
				"mode",
				"name_alias",
				"owner",
				"relation_dhcp_rs_prov",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
