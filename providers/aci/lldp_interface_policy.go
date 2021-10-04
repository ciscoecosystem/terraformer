package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const LLDPInterfacePolicyClassName = "lldpIfPol"

type LLDPInterfacePolicyGenerator struct {
	ACIService
}

func (a *LLDPInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, LLDPInterfacePolicyClassName)

	LLDPInterfacePolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LLDPInterfacePolicyCount, err := strconv.Atoi(stripQuotes(LLDPInterfacePolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LLDPInterfacePolicyCount; i++ {
		LLDPInterfacePolicyDN := LLDPInterfacePolicysCont.S("imdata").Index(i).S(LLDPInterfacePolicyClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(LLDPInterfacePolicyDN),
			stripQuotes(LLDPInterfacePolicyDN),
			"aci_lldp_interface_policy",
			"aci",
			[]string{
				"admin_rx_st",
				"admin_tx_st",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
