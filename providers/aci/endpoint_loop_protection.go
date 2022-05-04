package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const ePLoopProtectionPolicyClassName = "epLoopProtectP"

type EPLoopProtectionPolicyGenerator struct {
	ACIService
}

func (a *EPLoopProtectionPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, ePLoopProtectionPolicyClassName)

	EPLoopProtectionPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	EPLoopProtectionPolicyCount, err := strconv.Atoi(stripQuotes(EPLoopProtectionPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < EPLoopProtectionPolicyCount; i++ {
		EPLoopProtectionPolicyAttr := EPLoopProtectionPolicyCont.S("imdata").Index(i).S(ePLoopProtectionPolicyClassName, "attributes")
		EPLoopProtectionPolicyDN := G(EPLoopProtectionPolicyAttr, "dn")
		//name := G(EPLoopProtectionPolicyAttr, "name")
		if filterChildrenDn(EPLoopProtectionPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				EPLoopProtectionPolicyDN,
				resourceNamefromDn(ePLoopProtectionPolicyClassName, EPLoopProtectionPolicyDN, i),
				"aci_endpoint_loop_protection",
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
