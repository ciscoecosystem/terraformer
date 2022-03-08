package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaGenerator struct {
	MSOService
}

func (a *SchemaGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaId := stripQuotes(con.S("schemas").Index(i).S("id").String())
		templateName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(0).S("name").String())
		tenantId := stripQuotes(con.S("schemas").Index(i).S("templates").Index(0).S("tenantId").String())
		schemaName := schemaId + "_" + templateName + "_" + tenantId
		resource := terraformutils.NewResource(
			schemaId,
			schemaName,
			"mso_schema",
			"mso",
			map[string]string{
				"template_name": templateName,
				"tenant_id":     tenantId,
			},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
