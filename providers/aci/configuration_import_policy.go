package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const configImportPolicyClassName = "configImportP"

type ConfigImportPolicyGenerator struct {
	ACIService
}

func (a *ConfigImportPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, configImportPolicyClassName)

	configImportPolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	configImportPolicyCount, err := strconv.Atoi(stripQuotes(configImportPolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < configImportPolicyCount; i++ {
		configImportPolicyDN := stripQuotes(configImportPolicysCont.S("imdata").Index(i).S(configImportPolicyClassName, "attributes", "dn").String())
		if filterChildrenDn(configImportPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				configImportPolicyDN,
				resourceNamefromDn(configImportPolicyClassName, (configImportPolicyDN), i),
				"aci_configuration_import_policy",
				"aci",
				[]string{
					"admin_st",
					"fail_on_decrypt_errors",
					"file_name",
					"import_mode",
					"import_type",
					"name_alias",
					"snapshot",
					"relation_config_rs_import_source",
					"relation_trig_rs_triggerable",
					"relation_config_rs_remote_path",
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
