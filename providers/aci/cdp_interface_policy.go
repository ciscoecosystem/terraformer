package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CDPInterfacePolicyClassName = "cdpIfPol"

type CDPInterfacePolicyGenerator struct {
	ACIService
}

func (a *CDPInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CDPInterfacePolicyClassName)

	CDPInterfacePolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	CDPInterfacePolicyCount, err := strconv.Atoi(stripQuotes(CDPInterfacePolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < CDPInterfacePolicyCount; i++ {
		CDPInterfacePolicyDN := stripQuotes(CDPInterfacePolicysCont.S("imdata").Index(i).S(CDPInterfacePolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(CDPInterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CDPInterfacePolicyDN,
				fmt.Sprintf("%s_%s_%d", CDPInterfacePolicyClassName, GetMOName(CDPInterfacePolicyDN), i),
				"aci_cdp_interface_policy",
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
