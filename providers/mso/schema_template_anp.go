package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateAnpGenerator struct {
	MSOService
}

func (a *SchemaTemplateAnpGenerator) InitResources() error {
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
		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			anpsLen := 0
			if templateCont.Exists("anps") {
				anpsLen = len(templateCont.S("anps").Data().([]interface{}))
			}
			for k := 0; k < anpsLen; k++ {
				anps := templateCont.S("anps").Index(k)
				name := models.G(anps, "name")
				displayName := models.G(anps, "displayName")
				anpRef := models.G(anps, "anpRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				resourceName := match[1] + "_" + match[2] + "_" + name
				resource := terraformutils.NewResource(
					name,
					resourceName,
					"mso_schema_template_anp",
					"mso",
					map[string]string{
						"schema_id":    match[1],
						"template":     match[2],
						"name":         name,
						"display_name": displayName,
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
