package mso

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DhcpRelayPolicyProviderGenerator struct {
	MSOService
}

func (a *DhcpRelayPolicyProviderGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getDhcpRelayContainer(mso)
	if err != nil {
		return err
	}
	dhcpLength := len(con.S("DhcpRelayPolicies").Data().([]interface{}))
	for i := 0; i < dhcpLength; i++ {
		dhcpCon := con.S("DhcpRelayPolicies").Index(i)
		dhcpProviderLen := 0
		if dhcpCon.Exists("provider") {
			dhcpProviderLen = len(dhcpCon.S("provider").Data().([]interface{}))
		}
		for j := 0; j < dhcpProviderLen; j++ {
			policyName := stripQuotes(dhcpCon.S("name").String())
			providerCon := dhcpCon.S("provider").Index(j)
			epgRef := stripQuotes(providerCon.S("epgRef").String())
			extepgRef := stripQuotes(providerCon.S("externalEpgRef").String())
			addr := stripQuotes(providerCon.S("addr").String())
			var id string
			var name string
			if epgRef != "" {
				id = fmt.Sprintf("%s%s/%s", policyName, epgRef, addr)
				name = policyName + "_" + strings.Replace(epgRef, "/", "_", -1) + "_" + strconv.Itoa(int(hash(addr)))
			} else {
				id = fmt.Sprintf("%s%s/%s", policyName, extepgRef, addr)
				name = policyName + "_" + strings.Replace(extepgRef, "/", "_", -1) + "_" + strconv.Itoa(int(hash(addr)))
			}

			resource := terraformutils.NewResource(
				id,
				name,
				"mso_dhcp_relay_policy_provider",
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
