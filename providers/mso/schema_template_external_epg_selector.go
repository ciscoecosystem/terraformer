package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateExternalEPGSelector struct {
	MSOService
}

func (a *SchemaTemplateExternalEPGSelector) InitResources() error {
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
		templateCount := 0
		if schemaCont.Exists("templates") {
			templateCount = len(schemaCont.S("templates").Data().([]interface{}))
		}

		for j := 0; j < templateCount; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())
			epgCount := 0
			if templateCont.Exists("externalEpgs") {
				epgCount = len(templateCont.S("externalEpgs").Data().([]interface{}))
			}

			for k := 0; k < epgCount; k++ {
				epgCont := templateCont.S("externalEpgs").Index(k)
				epgName := stripQuotes(epgCont.S("name").String())
				selectorCount := 0
				if epgCont.Exists("selectors") && epgCont.S("selectors").Data() != nil {
					selectorCount = len(epgCont.S("selectors").Data().([]interface{}))
				}
				for l := 0; l < selectorCount; l++ {
					selectorCont := epgCont.S("selectors").Index(l)
					selectorName := stripQuotes(selectorCont.S("name").String())
					resourceName := schemaId + "_" + templateName + "_" + epgName + "_" + selectorName
					resource := terraformutils.NewResource(
						selectorName,
						resourceName,
						"mso_schema_template_external_epg_selector",
						"mso",
						map[string]string{
							"schema_id":         schemaId,
							"template_name":     templateName,
							"external_epg_name": epgName,
							"name":              selectorName,
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
