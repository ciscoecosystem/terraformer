package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vlanPoolClass = "fvnsVlanInstP"

type VlanPoolGenerator struct {
	ACIService
}

func (a *VlanPoolGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vlanPoolClass)

	vlanPoolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(vlanPoolCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	vlanPoolCount, err := strconv.Atoi(stripQuotes(vlanPoolCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < vlanPoolCount; i++ {
		vlanPoolProfileDN := stripQuotes(vlanPoolCont.S("imdata").Index(i).S(vlanPoolClass, "attributes", "dn").String())
		if filterChildrenDn(vlanPoolProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				vlanPoolProfileDN,
				vlanPoolProfileDN,
				"aci_vlan_pool",
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
