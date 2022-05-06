package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const oOBManagedNodesZoneClassName = "mgmtOoBZone"
const inBManagedNodesZoneClassName = "mgmtInBZone"

type OOBManagedNodesZoneGenerator struct {
	ACIService
}

func (a *OOBManagedNodesZoneGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	OobDnURL := fmt.Sprintf("%s/%s.json", baseURL, oOBManagedNodesZoneClassName)
	IndDnURL := fmt.Sprintf("%s/%s.json", baseURL, inBManagedNodesZoneClassName)
	OOBManagedNodesZoneCont, err := client.GetViaURL(OobDnURL)
	if err != nil {
		return err
	}

	INBManagedNodesZoneCont, err := client.GetViaURL(IndDnURL)
	if err != nil {
		return err
	}

	OOBManagedNodesZoneCount, err := strconv.Atoi(stripQuotes(OOBManagedNodesZoneCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	INBManagedNodesZoneCount, err := strconv.Atoi(stripQuotes(INBManagedNodesZoneCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < OOBManagedNodesZoneCount; i++ {
		OOBManagedNodesZoneAttr := OOBManagedNodesZoneCont.S("imdata").Index(i).S(oOBManagedNodesZoneClassName, "attributes")
		OOBManagedNodesZoneDN := G(OOBManagedNodesZoneAttr, "dn")
		nameMgmtZone := G(OOBManagedNodesZoneAttr, "name")
		type_band := "out_of_band"
		if filterChildrenDn(OOBManagedNodesZoneDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				OOBManagedNodesZoneDN,
				resourceNamefromDn(oOBManagedNodesZoneClassName, OOBManagedNodesZoneDN, i),
				"aci_mgmt_zone",
				"aci",
				map[string]string{
					"managed_node_connectivity_group_dn": GetParentDn(OOBManagedNodesZoneDN, fmt.Sprintf("/oobzone")),
					"type":                               type_band,
					"name":                               nameMgmtZone,
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	for i := 0; i < INBManagedNodesZoneCount; i++ {
		INBManagedNodesZoneAttr := INBManagedNodesZoneCont.S("imdata").Index(i).S(inBManagedNodesZoneClassName, "attributes")
		INBManagedNodesZoneDN := G(INBManagedNodesZoneAttr, "dn")
		nameMgmtZone := G(INBManagedNodesZoneAttr, "name")
		type_band := "in_band"
		if filterChildrenDn(INBManagedNodesZoneDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				INBManagedNodesZoneDN,
				resourceNamefromDn(inBManagedNodesZoneClassName, INBManagedNodesZoneDN, i),
				"aci_mgmt_zone",
				"aci",
				map[string]string{
					"managed_node_connectivity_group_dn": GetParentDn(INBManagedNodesZoneDN, fmt.Sprintf("/inbzone")),
					"type":                               type_band,
					"name":                               nameMgmtZone,
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
