package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpineInterfaceProfileClassName = "infraSpAccPortP"

type SpineInterfaceProfileGenerator struct {
	ACIService
}

func (a *SpineInterfaceProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpineInterfaceProfileClassName)

	SpineInterfaceProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SpineInterfaceProfileCount, err := strconv.Atoi(stripQuotes(SpineInterfaceProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SpineInterfaceProfileCount; i++ {
		SpineInterfaceProfileDN := SpineInterfaceProfileCont.S("imdata").Index(i).S(SpineInterfaceProfileClassName, "attributes", "dn").String()
		if filterChildrenDn(SpineInterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(SpineInterfaceProfileDN),
				stripQuotes(SpineInterfaceProfileDN),
				"aci_spine_interface_profile",
				"aci",
				[]string{
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
