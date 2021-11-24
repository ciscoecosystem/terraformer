package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const podMaintenanceGroupClassName = "maintMaintGrp"

type PodMaintenanceGroupGenerator struct {
	ACIService
}

func (a *PodMaintenanceGroupGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, podMaintenanceGroupClassName)

	PodMaintenanceGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	PodMaintenanceGroupCount, err := strconv.Atoi(stripQuotes(PodMaintenanceGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < PodMaintenanceGroupCount; i++ {
		PodMaintenanceGroupAttr := PodMaintenanceGroupCont.S("imdata").Index(i).S(podMaintenanceGroupClassName, "attributes")
		PodMaintenanceGroupDN := G(PodMaintenanceGroupAttr, "dn")

		if filterChildrenDn(PodMaintenanceGroupDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				PodMaintenanceGroupDN,
				resourceNamefromDn(podMaintenanceGroupClassName, PodMaintenanceGroupDN, i),
				"aci_pod_maintenance_group",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"fwtype",
					"name_alias",
					"pod_maintenance_group_type",
					"relation_maint_rs_mgrpp",
					"annotation",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
