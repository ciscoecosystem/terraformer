package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vpcExplicitProtectionGroupClassName = "fabricExplicitGEp"

type VPCExplicitProtectionGroupGenerator struct {
	ACIService
}

func (a *VPCExplicitProtectionGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client:= clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vpcExplicitProtectionGroupClassName)

	VPCExplicitProtectionGroupsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	VPCExplicitProtectionGroupCount, err := strconv.Atoi(stripQuotes(VPCExplicitProtectionGroupsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < VPCExplicitProtectionGroupCount; i++ {
		VPCExplicitProtectionGroupDN := VPCExplicitProtectionGroupsCont.S("imdata").Index(i).S(vpcExplicitProtectionGroupClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(VPCExplicitProtectionGroupDN),
			stripQuotes(VPCExplicitProtectionGroupDN),
			"aci_vpc_explicit_protection_group",
			"aci",
			[]string{
				"vpc_domain_policy",
				"annotation",
				"vpc_explicit_protection_group_id",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
