package mso

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type MSOProvider struct {
	terraformutils.Provider
	baseURL  string
	username string
	password string
	insecure bool
	domain   string
	platform string
}

func (p MSOProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		// "schema":{
		// 	"tenant":[]string{"tenant_id","id"},
		// },
		"schema_site": {
			"schema": []string{"schema_id", "id"},
			// "site":   []string{"site_id", "id"},
		},
		"schema_template_anp_epg": {
			"schema": []string{"schema_id", "id"},
			// "schema_template": []string{"template_name", "name"},
			// "schema_template_anp": []string{"anp_name", "name"},
		},
		"tenant": {
			"site": []string{"site_associations.site_id", "id"},
		},
		"schema_template_bd": {
			"schema": []string{"schema_id", "id"},
			// "schema_template_vrf": []string{
			// 	"vrf_name", "name",
			// 	"vrf_schema_id", "schema_id",
			// 	"vrf_template_name", "template",
			// },
			// "dhcp_relay_policy":  []string{"name", "name"},
			// "dhcp_option_policy": []string{"dhcp_option_policy_name", "name"},
		},
	}
}

func (p MSOProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"mso": map[string]interface{}{
				"username": p.username,
				"password": p.password,
				"url":      p.baseURL,
				"domain":   p.domain,
				"platform": p.platform,
			},
		},
	}
}

func (p *MSOProvider) Init(args []string) error {
	p.baseURL = args[0]
	p.username = args[1]
	p.password = args[2]
	p.insecure = true
	p.domain = args[3]
	p.platform = args[4]
	os.Setenv("MSO_URL", p.baseURL)
	os.Setenv("MSO_USERNAME", p.username)
	os.Setenv("MSO_PASSWORD", p.password)
	os.Setenv("MSO_DOMAIN", p.domain)
	os.Setenv("MSO_PLATFORM", p.platform)
	return nil
}

func (p *MSOProvider) GetName() string {
	return "mso"
}

func (p *MSOProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username": p.username,
		"password": p.password,
		"base_url": p.baseURL,
		"insecure": p.insecure,
		"domain":   p.domain,
		"platform": p.platform,
	})
	return nil
}

func (p *MSOProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"schema":                  &SchemaGenerator{},
		"schema_site":             &SchemaSiteGenerator{},
		"label":                   &LabelGenerator{},
		"schema_template_anp_epg": &SchemaTemplateAnpEpgGenerator{},
		"site":                    &SiteGenerator{},
		"tenant":                  &TenantGenerator{},
		"schema_template_bd":      &SchemaTemplateBdGenerator{},
		"schema_template":         &SchemaTemlateGenerator{},
	}
}
