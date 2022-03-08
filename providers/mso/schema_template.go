package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemlateGenerator struct {
	MSOService
}

func (a *SchemaTemlateGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := models.G(schemaCont, "id")
		templateCount := 0
		if schemaCont.Exists("templates") {
			templateCount = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateCount; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			tenantId := models.G(templateCont, "tenantId")
			templateName := models.G(templateCont, "name")
			displayName := models.G(templateCont, "displayName")
			name := schemaId + "_" + templateName
			resource := terraformutils.NewResource(
				templateName,
				name,
				"mso_schema_template",
				"mso",
				map[string]string{
					"schema_id":    schemaId,
					"tenant_id":    tenantId,
					"name":         templateName,
					"display_name": displayName,
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
