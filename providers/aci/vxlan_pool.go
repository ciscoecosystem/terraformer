package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vxlanPoolClass = "fvnsVxlanInstP"

type VxlanPoolGenerator struct {
	ACIService
}

func (a *VxlanPoolGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vxlanPoolClass)

	vxlanPoolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(vxlanPoolCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	vxlanPoolCount, err := strconv.Atoi(stripQuotes(vxlanPoolCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < vxlanPoolCount; i++ {
		vxlanPoolProfileDN := stripQuotes(vxlanPoolCont.S("imdata").Index(i).S(vxlanPoolClass, "attributes", "dn").String())
		if filterChildrenDn(vxlanPoolProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				vxlanPoolProfileDN,
				vxlanPoolProfileDN,
				"aci_vxlan_pool",
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
