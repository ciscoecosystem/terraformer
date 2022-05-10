package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const interfaceProfileClassName = "infraRsSpAccPortP"

type InterfaceProfileGenerator struct {
	ACIService
}

func (a *InterfaceProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, interfaceProfileClassName)

	InterfaceProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	InterfaceProfileCount, err := strconv.Atoi(stripQuotes(InterfaceProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < InterfaceProfileCount; i++ {
		InterfaceProfileAttr := InterfaceProfileCont.S("imdata").Index(i).S(interfaceProfileClassName, "attributes")
		InterfaceProfileDN := G(InterfaceProfileAttr, "dn")
		tDn := G(InterfaceProfileAttr, "tDn")
		if filterChildrenDn(InterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				InterfaceProfileDN,
				resourceNamefromDn(interfaceProfileClassName, InterfaceProfileDN, i),
				"aci_spine_interface_profile_selector",
				"aci",
				map[string]string{
					"spine_profile_dn": GetParentDn(InterfaceProfileDN, fmt.Sprintf("/rsspAccPortP-[%s]", tDn)),
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
