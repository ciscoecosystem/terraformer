package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const dhcpOptionPolicyClassName = "dhcpOptionPol"

type DhcpOptionPolicyGenerator struct {
	ACIService
}

func (a *DhcpOptionPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, dhcpOptionPolicyClassName)

	DhcpOptionPoliciesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	DhcpOptionPolicyCount, err := strconv.Atoi(stripQuotes(DhcpOptionPoliciesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < DhcpOptionPolicyCount; i++ {
		DhcpOptionPolicyDN := stripQuotes(DhcpOptionPoliciesCont.S("imdata").Index(i).S(dhcpOptionPolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(DhcpOptionPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(DhcpOptionPolicyDN),
				stripQuotes(DhcpOptionPolicyDN),
				"aci_dhcp_option_policy",
				"aci",
				[]string{
					"name_alias",
					"dhcp_option",
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
