package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteExternalEPGSelector struct {
	MSOService
}

func (a *SchemaSiteExternalEPGSelector) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}

	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())
		sitesCount := 0
		if schemaCont.Exists("sites") {
			sitesCount = len(schemaCont.S("sites").Data().([]interface{}))
		}

		for j := 0; j < sitesCount; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			templateName := stripQuotes(siteCont.S("templateName").String())
			siteId := stripQuotes(siteCont.S("siteId").String())
			epgCount := 0
			if siteCont.Exists("externalEpgs") {
				epgCount = len(siteCont.S("externalEpgs").Data().([]interface{}))
			}

			for k := 0; k < epgCount; k++ {
				epgCont := siteCont.S("externalEpgs").Index(k)
				epgRef := stripQuotes(epgCont.S("externalEpgRef").String())
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/externalEpgs/(.*)")
				match := re.FindStringSubmatch(epgRef)
				epgName := match[3]
				selectorCount := 0
				if epgCont.Exists("subnets") && epgCont.S("subnets").Data() != nil {
					selectorCount = len(epgCont.S("subnets").Data().([]interface{}))
				}
				for l := 0; l < selectorCount; l++ {
					selectorCont := epgCont.S("subnets").Index(l)
					selectorName := stripQuotes(selectorCont.S("name").String())
					selectorIp := stripQuotes(selectorCont.S("ip").String())
					resourceName := schemaId + "_" + siteId + "_" + templateName + "_" + epgName + "_" + selectorName
					resource := terraformutils.NewResource(
						selectorName,
						resourceName,
						"mso_schema_site_external_epg_selector",
						"mso",
						map[string]string{
							"schema_id":         schemaId,
							"site_id":           siteId,
							"template_name":     templateName,
							"external_epg_name": epgName,
							"name":              selectorName,
							"ip":                selectorIp,
						},
						[]string{},
						map[string]interface{}{},
					)
					resource.SlowQueryRequired = true
					a.Resources = append(a.Resources, resource)
				}
			}
		}

	}
	return nil
}
