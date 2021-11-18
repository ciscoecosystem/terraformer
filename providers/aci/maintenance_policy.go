package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const MaintenancePolicyClassName = "maintMaintP"

type MaintenancePolicyGenerator struct {
	ACIService
}

func (a *MaintenancePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, MaintenancePolicyClassName)

	MaintenancePolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MaintenancePolicyCount, err := strconv.Atoi(stripQuotes(MaintenancePolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MaintenancePolicyCount; i++ {
		MaintenancePolicyDN := MaintenancePolicysCont.S("imdata").Index(i).S(MaintenancePolicyClassName, "attributes", "dn").String()
		if filterChildrenDn(MaintenancePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(MaintenancePolicyDN),
				stripQuotes(MaintenancePolicyDN),
				"aci_maintenance_policy",
				"aci",
				[]string{
					"admin_st",
					"graceful",
					"ignore_compat",
					"internal_label",
					"notif_cond",
					"run_mode",
					"version",
					"version_check_override",
					"relation_maint_rs_pol_scheduler",
					"relation_maint_rs_pol_notif",
					"relation_trig_rs_triggerable",
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
