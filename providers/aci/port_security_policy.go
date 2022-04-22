package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const PortSecurityPolicyClass = "l2PortSecurityPol"

type PortSecurityPolicyGenerator struct {
	ACIService
}

func (a *PortSecurityPolicyGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, PortSecurityPolicyClass)

	PortSecurityPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	PortSecurityPolicyCount, err := strconv.Atoi(stripQuotes(PortSecurityPolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < PortSecurityPolicyCount; i++ {
		PortSecurityPolicyProfileDN := stripQuotes(PortSecurityPolicyCont.S("imdata").Index(i).S(PortSecurityPolicyClass, "attributes", "dn").String())
		if filterChildrenDn(PortSecurityPolicyProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				PortSecurityPolicyProfileDN,
				resourceNamefromDn(PortSecurityPolicyClass, (PortSecurityPolicyProfileDN), i),
				"aci_port_security_policy",
				"aci",
				[]string{
					"maximum",
					"name_alias",
					"timeout",
					"violation",
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
