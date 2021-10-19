package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const HSRPInterfacePolicyClass = "hsrpIfPol"

type HSRPInterfacePolicyGenerator struct {
	ACIService
}

func (a *HSRPInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, HSRPInterfacePolicyClass)

	HSRPInterfacePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	HSRPInterfacePolicyCount, err := strconv.Atoi(stripQuotes(HSRPInterfacePolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < HSRPInterfacePolicyCount; i++ {
		HSRPInterfacePolicyDN := stripQuotes(HSRPInterfacePolicyCont.S("imdata").Index(i).S(HSRPInterfacePolicyClass, "attributes", "dn").String())
		if filterChildrenDn(HSRPInterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				HSRPInterfacePolicyDN,
				HSRPInterfacePolicyDN,
				"aci_hsrp_interface_policy",
				"aci",
				[]string{
					"ctrl",
					"delay",
					"name_alias",
					"reload_delay",
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
