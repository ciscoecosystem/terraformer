package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateAnpEpgSelector struct {
	MSOService
}

func (a *SchemaTemplateAnpEpgSelector) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())
		templateLen := len(schemaCont.S("templates").Data().([]interface{}))

		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			anpsLen := 0
			if templateCont.Exists("anps") {
				anpsLen = len(templateCont.S("anps").Data().([]interface{}))
			}

			for k := 0; k < anpsLen; k++ {
				anpCont := templateCont.S("anps").Index(k)
				anpName := models.G(anpCont, "name")

				epgsLen := 0
				if anpCont.Exists("epgs") {
					epgsLen = len(anpCont.S("epgs").Data().([]interface{}))
				}

				for m := 0; m < epgsLen; m++ {
					epgCont := anpCont.S("epgs").Index(m)
					epgName := models.G(epgCont, "name")
					selectorLen := 0
					if epgCont.Exists("selectors") {
						selectorLen = len(epgCont.S("selectors").Data().([]interface{}))
					}
					for n := 0; n < selectorLen; n++ {
						selectorCont := epgCont.S("selectors").Index(n)
						selectorName := models.G(selectorCont, "name")
						resourceName := schemaId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + selectorName
						resource := terraformutils.NewResource(
							selectorName,
							resourceName,
							"mso_schema_template_anp_epg_selector",
							"mso",
							map[string]string{
								"schema_id":     schemaId,
								"template_name": templateName,
								"anp_name":      anpName,
								"epg_name":      epgName,
							},
							[]string{},
							map[string]interface{}{},
						)
						resource.SlowQueryRequired = SlowQueryRequired
						a.Resources = append(a.Resources, resource)
					}
				}
			}
		}
	}
	return nil
}
