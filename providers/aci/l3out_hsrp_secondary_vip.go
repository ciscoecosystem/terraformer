package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutHSRPSecondaryVipClassName = "hsrpSecVip"

type L3OutHSRPSecondaryVipGenerator struct {
	ACIService
}

func (a *L3OutHSRPSecondaryVipGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutHSRPSecondaryVipClassName)

	L3OutHSRPSecondaryVipCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutHSRPSecondaryVipCount, err := strconv.Atoi(stripQuotes(L3OutHSRPSecondaryVipCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L3OutHSRPSecondaryVipCount; i++ {
		L3OutHSRPSecondaryVipDN := stripQuotes(L3OutHSRPSecondaryVipCont.S("imdata").Index(i).S(L3OutHSRPSecondaryVipClassName, "attributes", "dn").String())
		if filterChildrenDn(L3OutHSRPSecondaryVipDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutHSRPSecondaryVipDN,
				L3OutHSRPSecondaryVipDN,
				"aci_l3out_hsrp_secondary_vip",
				"aci",
				[]string{
					"config_issues",
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
