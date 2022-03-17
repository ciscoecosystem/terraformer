package mso

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SlowQueryRequired = false

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
		"schema": {
			"tenant":          []string{"tenant_id", "id"},
			"schema_template": []string{"template_name", "name"},
		},
		"schema_site": {
			"schema":          []string{"schema_id", "id"},
			"site":            []string{"site_id", "id"},
			"schema_template": []string{"template_name", "name"},
		},
		"schema_template": {
			"schema": []string{"schema_id", "id"},
			"tenant": []string{"tenant_id", "id"},
		},
		"schema_template_anp_epg": {
			"schema":              []string{"schema_id", "id"},
			"schema_template":     []string{"template_name", "name"},
			"schema_template_anp": []string{"anp_name", "name"},
		},
		"tenant": {
			"site": []string{"site_associations.site_id", "id"},
		},
		"schema_template_bd": {
			"schema":              []string{"schema_id", "id", "vrf_schema_id", "id"},
			"schema_template_vrf": []string{"vrf_name", "name"},
			"schema_template":     []string{"template_name", "name", "vrf_template_name", "name"},
			// "dhcp_relay_policy":  []string{"name", "name"},
			// "dhcp_option_policy": []string{"dhcp_option_policy_name", "name"},
		},
		"schema_template_bd_subnet": {
			"schema":             []string{"schema_id", "id"},
			"schema_template":    []string{"template_name", "name"},
			"schema_template_bd": []string{"bd_name", "name"},
		},
		"schema_template_anp": {
			"schema":          []string{"schema_id", "id"},
			"schema_template": []string{"template", "name"},
		},
		"schema_template_anp_epg_subnet": {
			"schema":                  []string{"schema_id", "id"},
			"schema_template":         []string{"template", "name"},
			"schema_template_anp":     []string{"anp_name", "name"},
			"schema_template_anp_epg": []string{"epg_name", "name"},
		},
		"schema_template_vrf": {
			"schema":          []string{"schema_id", "id"},
			"schema_template": []string{"template", "name"},
		},
		"schema_template_external_epg_contract": {
			"schema":                       []string{"schema_id", "id"},
			"schema_template":              []string{"template_name", "name"},
			"schema_template_external_epg": []string{"external_epg_name", "external_epg_name"},
		},
		"schema_template_anp_epg_contract": {
			"schema":                  []string{"schema_id", "id", "contract_schema_id", "id"},
			"schema_template":         []string{"template_name", "name", "contract_template_name", "name"},
			"schema_template_anp":     []string{"anp_name", "name"},
			"schema_template_anp_epg": []string{"epg_name", "name"},
		},
		"schema_site_anp_epg_domain": {
			"schema":              []string{"schema_id", "id"},
			"schema_template":     []string{"template_name", "name"},
			"schema_site_anp":     []string{"anp_name", "anp_name"},
			"schema_site_anp_epg": []string{"epg_name", "epg_name"},
			"site":                []string{"site_id", "id"},
		},
		"schema_template_l3out": {
			"schema":              []string{"schema_id", "id", "vrf_schema_id", "id"},
			"schema_template":     []string{"template_name", "name", "vrf_template_name", "name"},
			"schema_template_vrf": []string{"vrf_name", "name"},
		},
		"schema_site_vrf_region_cidr": {
			"schema":                 []string{"schema_id", "id"},
			"schema_template":        []string{"template_name", "name"},
			"site":                   []string{"site_id", "id"},
			"schema_site_vrf":        []string{"vrf_name", "vrf_name"},
			"schema_site_vrf_region": []string{"region_name", "region_name"},
		},
		"schema_template_filter_entry": {
			"schema":                   []string{"schema_id", "id"},
			"schema_template":          []string{"template_name", "name"},
			"schema_template_contract": []string{"name", "contract_name"},
		},
		"schema_site_anp_epg_static_port": {
			"schema_site_anp":     []string{"anp_name", "anp_name"},
			"schema_site_anp_epg": []string{"epg_name", "epg_name"},
			"schema_site":         []string{"site_id", "site_id"},
			"schema_template":     []string{"template_name", "name"},
			"schema":              []string{"schema_id", "id"},
		},
		"schema_template_contract_filter": {
			"schema":                       []string{"schema_id", "id", "filter_schema_id", "id"},
			"schema_template":              []string{"template_name", "name", "filter_template_name", "name"},
			"schema_template_filter_entry": []string{"filter_name", "name"},
			// "schema_template_contract": []string{"contract_name", "contract_name"},
		},
		"schema_site_anp_epg_static_leaf": {
			"schema":              []string{"schema_id", "id"},
			"schema_template":     []string{"template_name", "name"},
			"schema_site":         []string{"site_id", "site_id"},
			"schema_site_anp":     []string{"anp_name", "anp_name"},
			"schema_site_anp_epg": []string{"epg_name", "epg_name"},
		},
		"schema_site_vrf_region": {
			"schema":          []string{"schema_id", "id"},
			"schema_template": []string{"template_name", "name"},
			"schema_site":     []string{"site_id", "site_id"},
			// "schema_site_vrf_region_hub_network": []string{
			// 	"hub_network.name", "name",
			// 	"hub_network.tenant_name", "tenant_name",
			// },
			"schema_site_vrf_region_cidr": []string{
				"cidr.cidr_ip", "ip",
				"cidr.primary", "primary",
			},
			"schema_site_vrf_region_cidr_subnet": []string{
				"cidr.subnet.ip", "ip",
				"cidr.subnet.zone", "zone",
				"cidr.subnet.usage", "usage",
			},
			"schema_site_vrf": []string{"vrf_name", "vrf_name"},
		},
		"schema_site_vrf_region_cidr_subnet": {
			"schema":          []string{"schema_id", "id"},
			"schema_template": []string{"template_name", "name"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_site_vrf_region_cidr": []string{
				"cidr_ip", "ip",
			},
			"schema_site_vrf_region": []string{
				"region_name", "region_name",
			},
			"schema_site_vrf": []string{"vrf_name", "vrf_name"},
		},
		"schema_template_external_epg": {
			"schema":                                []string{"schema_id", "id", "vrf_schema_id", "id", "anp_schema_id", "id", "l3out_schema_id", "id"},
			"schema_template":                       []string{"template_name", "name", "vrf_template_name", "name", "anp_template_name", "name", "l3out_template_name", "name"},
			"schema_template_vrf":                   []string{"vrf_name", "name"},
			"schema_template_anp":                   []string{"anp_name", "name"},
			"schema_template_l3out":                 []string{"l3out_name", "l3out_name"},
			"schema_template_external_epg_selector": []string{"selector_name", "name"},
		},
		"schema_site_bd_subnet": {
			"schema":          []string{"schema_id", "id"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_template": []string{"template_name", "name"},
			// "schema_site_bd":  []string{"bd_name", "bd_name"},
		},
		"schema_template_external_epg_subnet": {
			"schema":                       []string{"schema_id", "id"},
			"schema_template":              []string{"template_name", "name"},
			"schema_template_external_epg": []string{"external_epg_name", "external_epg_name"},
		},
		"schema_site_vrf": {
			"schema":          []string{"schema_id", "id"},
			"schema_template": []string{"template_name", "name"},
			"schema_site":     []string{"site_id", "site_id"},
		},
		"schema_site_bd": {
			"schema":          []string{"schema_id", "id"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_template": []string{"template_name", "name"},
		},
		"schema_site_anp_epg_subnet": {
			"schema":                  []string{"schema_id", "id"},
			"schema_site":             []string{"site_id", "site_id"},
			"schema_template":         []string{"template_name", "name"},
			"schema_template_anp":     []string{"anp_name", "name"},
			"schema_template_anp_epg": []string{"epg_name", "name"},
		},
		"schema_template_anp_epg_selector": {
			"schema":                  []string{"schema_id", "id"},
			"schema_template":         []string{"template_name", "name"},
			"schema_template_anp":     []string{"anp_name", "name"},
			"schema_template_anp_epg": []string{"epg_name", "name"},
		},
		"schema_site_anp_epg": {
			"schema":          []string{"schema_id", "id"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_template": []string{"template_name", "name"},
			"schema_site_anp": []string{"anp_name", "anp_name"},
		},
		"schema_template_external_epg_selector": {
			"schema":                       []string{"schema_id", "id"},
			"schema_template":              []string{"template_name", "name"},
			"schema_template_external_epg": []string{"external_epg_name", "external_epg_name"},
		},
		"schema_site_bd_l3out": {
			"schema":          []string{"schema_id", "id"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_template": []string{"template_name", "name"},
			"schema_site_bd":  []string{"bd_name", "bd_name"},
		},
		"schema_site_anp": {
			"schema":          []string{"schema_id", "id"},
			"schema_site":     []string{"site_id", "site_id"},
			"schema_template": []string{"template_name", "name"},
		},
		"schema_site_external_epg": {
			"schema":            []string{"schema_id", "id"},
			"schema_site":       []string{"site_id", "site_id"},
			"schema_template":   []string{"template_name", "name"},
			"schema_site_l3out": []string{"l3out_name", "l3out_name"},
		},
		"schema_template_vrf_contract": {
			"schema":              []string{"schema_id", "id", "contract_schema_id", "id"},
			"schema_template":     []string{"template_name", "name", "contract_template_name", "name"},
			"schema_template_vrf": []string{"vrf_name", "name"},
		},
		"schema_template_anp_epg_useg_attr": {
			"schema":                  []string{"schema_id", "id"},
			"schema_template":         []string{"template_name", "name"},
			"schema_template_anp":     []string{"anp_name", "name"},
			"schema_template_anp_epg": []string{"epg_name", "name"},
		},
		"schema_site_anp_epg_selector": {
			"schema":              []string{"schema_id", "id"},
			"schema_site":         []string{"site_id", "site_id"},
			"schema_template":     []string{"template_name", "name"},
			"schema_site_anp":     []string{"anp_name", "anp_name"},
			"schema_site_anp_epg": []string{"epg_name", "epg_name"},
		},
		"schema_site_external_epg_selector": {
			"schema":                   []string{"schema_id", "id"},
			"schema_site":              []string{"site_id", "site_id"},
			"schema_template":          []string{"template_name", "name"},
			"schema_site_external_epg": []string{"external_epg_name", "external_epg_name"},
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
		"schema":                                &SchemaGenerator{},
		"schema_site":                           &SchemaSiteGenerator{},
		"label":                                 &LabelGenerator{},
		"schema_template_anp_epg":               &SchemaTemplateAnpEpgGenerator{},
		"site":                                  &SiteGenerator{},
		"tenant":                                &TenantGenerator{},
		"schema_template_bd":                    &SchemaTemplateBdGenerator{},
		"schema_template":                       &SchemaTemplateGenerator{},
		"schema_template_bd_subnet":             &SchemaTemplateBDSubnet{},
		"schema_template_anp":                   &SchemaTemplateAnpGenerator{},
		"schema_template_anp_epg_subnet":        &SchemaTemplateAnpEpgSubnet{},
		"schema_template_vrf":                   &SchemaTemplateVRFGenerator{},
		"schema_template_external_epg_contract": &SchemaTemplateExternalEpgContract{},
		"schema_template_anp_epg_contract":      &SchemaTemplateANPEPGContractGenerator{},
		"schema_site_anp_epg_domain":            &SchemaSiteAnpEpgDomain{},
		"schema_template_l3out":                 &SchemaTemplateL3OutGenerator{},
		"schema_site_vrf_region_cidr":           &SchemaSiteVrfRegionCidr{},
		"schema_template_filter_entry":          &SchemaTemplateFilterEntryGenerator{},
		"schema_site_anp_epg_static_port":       &SchemaSiteAnpEpgStaticPort{},
		"schema_template_contract_filter":       &SchemaTemplateContractFilter{},
		"schema_site_anp_epg_static_leaf":       &SchemaSiteAnpEpgStaticLeaf{},
		"schema_site_vrf_region":                &SchemaSiteVrfRegion{},
		"schema_template_external_epg":          &SchemaTemplateExternalEPG{},
		"schema_site_vrf_region_cidr_subnet":    &SchemaSiteVrfRegionCidrSubnet{},
		"schema_site_bd_subnet":                 &SchemaSiteBDSubnetGenerator{},
		"schema_template_external_epg_subnet":   &SchemaTemplateExternalEPGSubnet{},
		"schema_site_vrf":                       &SchemaSiteVRF{},
		"schema_site_bd":                        &SchemaSiteBdGenerator{},
		"schema_site_anp_epg_subnet":            &SchemaSiteAnpEpgSubnetGenerator{},
		"schema_site_bd_l3out":                  &SchemaSitel3OutsGenerator{},
		"schema_template_anp_epg_selector":      &SchemaTemplateAnpEpgSelector{},
		"schema_site_anp_epg":                   &SchemaSiteAnpEpg{},
		"schema_template_external_epg_selector": &SchemaTemplateExternalEPGSelector{},
		"schema_site_anp":                       &SchemaSiteAnpGenerator{},
		"schema_site_external_epg":              &SchemaSiteExternalEpgGenerator{},
		"schema_template_vrf_contract":          &SchemaTemplateVrfContractGenerator{},
		"schema_template_anp_epg_useg_attr":     &SchemaTemplateAnpEpgUsegAttr{},
		"schema_site_anp_epg_selector":          &SchemaSiteAnpEpgSelector{},
		"schema_site_external_epg_selector":     &SchemaSiteExternalEPGSelector{},
		"service_node_type":                     &ServiceNodeType{},
	}
}
