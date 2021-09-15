package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutHSRPInterfaceProfileClass = "hsrpIfP"

type L3OutHSRPInterfaceProfileGenerator struct {
	ACIService
}

func (a *L3OutHSRPInterfaceProfileGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutHSRPInterfaceProfileClass)

	L3OutHSRPInterfaceProfileCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3OutHSRPInterfaceProfileCount, err := strconv.Atoi(stripQuotes(L3OutHSRPInterfaceProfileCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3OutHSRPInterfaceProfileCount; i++ {
		L3OutHSRPInterfaceProfileDN := stripQuotes(L3OutHSRPInterfaceProfileCont.S("imdata").Index(i).S(L3OutHSRPInterfaceProfileClass, "attributes", "dn").String())
		if filterChildrenDn(L3OutHSRPInterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutHSRPInterfaceProfileDN,
				L3OutHSRPInterfaceProfileDN,
				"aci_l3out_hsrp_interface_profile",
				"aci",
				[]string{
					"name_alias",
					"version",
					"relation_hsrp_rs_if_pol",
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
