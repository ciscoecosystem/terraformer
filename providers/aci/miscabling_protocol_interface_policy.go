package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const MiscablingProtocolInterfacePolicyClassName = "mcpIfPol"

type MiscablingProtocolInterfacePolicyGenerator struct {
	ACIService
}

func (a *MiscablingProtocolInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, MiscablingProtocolInterfacePolicyClassName)

	MiscablingProtocolInterfacePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MiscablingProtocolInterfacePolicyCount, err := strconv.Atoi(stripQuotes(MiscablingProtocolInterfacePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MiscablingProtocolInterfacePolicyCount; i++ {
		MiscablingProtocolInterfacePolicyDN := MiscablingProtocolInterfacePolicyCont.S("imdata").Index(i).S(MiscablingProtocolInterfacePolicyClassName, "attributes", "dn").String()
		if filterChildrenDn(MiscablingProtocolInterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(MiscablingProtocolInterfacePolicyDN),
				stripQuotes(MiscablingProtocolInterfacePolicyDN),
				"aci_miscabling_protocol_interface_policy",
				"aci",
				[]string{
					"admin_st",
					"name_alias",
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
