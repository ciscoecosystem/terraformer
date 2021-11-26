package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const VlanVxlanTrafficClassName = "infraProvAcc"

type VlanVxlanTrafficGenerator struct {
	ACIService
}

func (a *VlanVxlanTrafficGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, VlanVxlanTrafficClassName)

	VlanVxlanTrafficsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	VlanVxlanTrafficCount, err := strconv.Atoi(stripQuotes(VlanVxlanTrafficsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < VlanVxlanTrafficCount; i++ {
		VlanVxlanTrafficDN := stripQuotes(VlanVxlanTrafficsCont.S("imdata").Index(i).S(VlanVxlanTrafficClassName, "attributes", "dn").String())
		if filterChildrenDn(VlanVxlanTrafficDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				VlanVxlanTrafficDN,
				resourceNamefromDn(VlanVxlanTrafficClassName, (VlanVxlanTrafficDN), i),
				"aci_vlan_encapsulationfor_vxlan_traffic",
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
