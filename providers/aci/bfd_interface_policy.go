package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const bFDInterfacePolicyClassName = "bfdIfPol"

type BFDInterfacePolicyGenerator struct {
	ACIService
}

func (a *BFDInterfacePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, bFDInterfacePolicyClassName)
	BFDInterfacePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	BFDInterfacePolicyCount, err := strconv.Atoi(stripQuotes(BFDInterfacePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < BFDInterfacePolicyCount; i++ {
		BFDInterfacePolicyAttr := BFDInterfacePolicyCont.S("imdata").Index(i).S(bFDInterfacePolicyClassName, "attributes")
		BFDInterfacePolicyDN := G(BFDInterfacePolicyAttr, "dn")
		name := G(BFDInterfacePolicyAttr, "name")
		if filterChildrenDn(BFDInterfacePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				BFDInterfacePolicyDN,
				resourceNamefromDn(bFDInterfacePolicyClassName, BFDInterfacePolicyDN, i),
				"aci_bfd_interface_policy",
				"aci",
				map[string]string{
					"tenant_dn": GetParentDn(BFDInterfacePolicyDN, fmt.Sprintf("/bfdIfPol-%s", name)),
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
