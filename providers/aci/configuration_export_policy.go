package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const ConfigExportPolicyClassName = "configExportP"

type ConfigExportPolicyGenerator struct {
	ACIService
}

func (a *ConfigExportPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, ConfigExportPolicyClassName)

	ConfigExportPolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ConfigExportPolicyCount, err := strconv.Atoi(stripQuotes(ConfigExportPolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ConfigExportPolicyCount; i++ {
		ConfigExportPolicyDN := stripQuotes(ConfigExportPolicysCont.S("imdata").Index(i).S(ConfigExportPolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(ConfigExportPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				ConfigExportPolicyDN,
				fmt.Sprintf("%s_%s_%d", ConfigExportPolicyClassName, GetMOName(ConfigExportPolicyDN), i),
				"aci_configuration_export_policy",
				"aci",
				[]string{
					"admin_st",
					"format",
					"include_secure_fields",
					"max_snapshot_count",
					"snapshot",
					"target_dn",
					"relation_config_rs_export_destination",
					"relation_trig_rs_triggerable",
					"relation_config_rs_remote_path",
					"relation_config_rs_export_scheduler",
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
