package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const physicalDomClass = "physDomP"

type PhysicalDomGenerator struct {
	ACIService
}

func (a *PhysicalDomGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, physicalDomClass)

	physicalDomCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(physicalDomCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	physicalDomCount, err := strconv.Atoi(stripQuotes(physicalDomCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < physicalDomCount; i++ {
		physicalDomProfileDN := stripQuotes(physicalDomCont.S("imdata").Index(i).S(physicalDomClass, "attributes", "dn").String())
		if filterChildrenDn(physicalDomProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				physicalDomProfileDN,
				physicalDomProfileDN,
				"aci_physical_domain",
				"aci",
				[]string{
					"name_alias",
					"relation_infra_rs_vlan_ns",
					"relation_infra_rs_vlan_ns_def",
					"relation_infra_rs_vip_addr_ns",
					"relation_infra_rs_dom_vxlan_ns_def",
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
