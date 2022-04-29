package aci

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const oOBManagedNodesZoneClassName = "mgmtOoBZone"

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
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, oOBManagedNodesZoneClassName)

	OOBManagedNodesZoneCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	OOBManagedNodesZoneCount, err := strconv.Atoi(stripQuotes(OOBManagedNodesZoneCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < OOBManagedNodesZoneCount; i++ {
		OOBManagedNodesZoneAttr := OOBManagedNodesZoneCont.S("imdata").Index(i).S(oOBManagedNodesZoneClassName, "attributes")
		OOBManagedNodesZoneDN := G(OOBManagedNodesZoneAttr, "dn")
		arr := strings.Split(OOBManagedNodesZoneDN, "/")
		type_band := ""
		if arr[4] == "oobzone" {
			type_band = "out_of_band"
		} else {
			type_band = "in_band"
		}

		if filterChildrenDn(OOBManagedNodesZoneDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				OOBManagedNodesZoneDN,
				resourceNamefromDn(oOBManagedNodesZoneClassName, OOBManagedNodesZoneDN, i),
				"aci_mgmt_zone",
				"aci",
				map[string]string{
					"managed_node_connectivity_group_dn": GetParentDn(OOBManagedNodesZoneDN, fmt.Sprintf("/oobzone")),
					"type":                               type_band,
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
