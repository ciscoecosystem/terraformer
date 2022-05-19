package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const bGPAddressFamilyContextPolicyClassName = "fvRsCtxToBgpCtxAfPol"

type BGPAddressFamilyContextPolicyGenerator struct {
	ACIService
}

func (a *BGPAddressFamilyContextPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, bGPAddressFamilyContextPolicyClassName)
	BGPAddressFamilyContextPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	BGPAddressFamilyContextPolicyCount, err := strconv.Atoi(stripQuotes(BGPAddressFamilyContextPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < BGPAddressFamilyContextPolicyCount; i++ {
		BGPAddressFamilyContextPolicyAttr := BGPAddressFamilyContextPolicyCont.S("imdata").Index(i).S(bGPAddressFamilyContextPolicyClassName, "attributes")
		BGPAddressFamilyContextPolicyDN := G(BGPAddressFamilyContextPolicyAttr, "dn")
		tnBgpCtxAfPolName := G(BGPAddressFamilyContextPolicyAttr, "tnBgpCtxAfPolName")
		af := G(BGPAddressFamilyContextPolicyAttr, "af")
		if filterChildrenDn(BGPAddressFamilyContextPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				BGPAddressFamilyContextPolicyDN,
				resourceNamefromDn(bGPAddressFamilyContextPolicyClassName, BGPAddressFamilyContextPolicyDN, i),
				"aci_vrf_to_bgp_address_family_context",
				"aci",
				map[string]string{
					"vrf_dn": GetParentDn(BGPAddressFamilyContextPolicyDN, fmt.Sprintf("/rsctxToBgpCtxAfPol-[%s]-%s", tnBgpCtxAfPolName, af)),
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
