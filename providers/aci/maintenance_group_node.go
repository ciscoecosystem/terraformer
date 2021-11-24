package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const MaintenanceGroupNodeClassName = "fabricNodeBlk"

type MaintenanceGroupNodeGenerator struct {
	ACIService
}

func (a *MaintenanceGroupNodeGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, MaintenanceGroupNodeClassName)

	MaintenanceGroupNodesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MaintenanceGroupNodeCount, err := strconv.Atoi(stripQuotes(MaintenanceGroupNodesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MaintenanceGroupNodeCount; i++ {
		MaintenanceGroupNodeDN := stripQuotes(MaintenanceGroupNodesCont.S("imdata").Index(i).S(MaintenanceGroupNodeClassName, "attributes", "dn").String())
		if filterChildrenDn(MaintenanceGroupNodeDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				MaintenanceGroupNodeDN,
				resourceNamefromDn(MaintenanceGroupNodeClassName, (MaintenanceGroupNodeDN), i),
				"aci_maintenance_group_node",
				"aci",
				[]string{
					"from_",
					"to_",
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
