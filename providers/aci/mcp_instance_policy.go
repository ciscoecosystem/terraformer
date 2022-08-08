package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const MiscablingProtocolInstancePolicyClassName = "mcpInstPol"

type MiscablingProtocolInstancePolicyGenerator struct {
	ACIService
}

func (a *MiscablingProtocolInstancePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, MiscablingProtocolInstancePolicyClassName)

	MiscablingProtocolInstancePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MiscablingProtocolInstancePolicyCount, err := strconv.Atoi(stripQuotes(MiscablingProtocolInstancePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MiscablingProtocolInstancePolicyCount; i++ {
		MiscablingProtocolInstancePolicyAttr := MiscablingProtocolInstancePolicyCont.S("imdata").Index(i).S(MiscablingProtocolInstancePolicyClassName, "attributes")
		MiscablingProtocolInstancePolicyDN := G(MiscablingProtocolInstancePolicyAttr, "dn")
		if filterChildrenDn(MiscablingProtocolInstancePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				MiscablingProtocolInstancePolicyDN,
				resourceNamefromDn(MiscablingProtocolInstancePolicyClassName, MiscablingProtocolInstancePolicyDN, i),
				"aci_mcp_instance_policy",
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
