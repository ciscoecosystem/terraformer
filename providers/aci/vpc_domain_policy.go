package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vPCDomainPolicyClassName = "vpcInstPol"

type VPCDomainPolicyGenerator struct {
	ACIService
}

func (a *VPCDomainPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vPCDomainPolicyClassName)

	VPCDomainPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	VPCDomainPolicyCount, err := strconv.Atoi(stripQuotes(VPCDomainPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < VPCDomainPolicyCount; i++ {
		VPCDomainPolicyAttr := VPCDomainPolicyCont.S("imdata").Index(i).S(vPCDomainPolicyClassName, "attributes")
		VPCDomainPolicyDN := G(VPCDomainPolicyAttr, "dn")
		name := G(VPCDomainPolicyAttr, "name")
		if filterChildrenDn(VPCDomainPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				VPCDomainPolicyDN,
				name,
				"aci_vpc_domain_policy",
				"aci",
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
