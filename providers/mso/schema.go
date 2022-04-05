package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/container"
)

type SchemaGenerator struct {
	MSOService
}

var globalSchemaCont *container.Container

func (a *SchemaGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
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
		resource.SlowQueryRequired = SlowQueryRequired
		a.Resources = append(a.Resources, resource)
	}
	return nil
}

func getSchemaContainer(mso *client.Client) (*container.Container, error) {
	if globalSchemaCont != nil {
		return globalSchemaCont, nil
	}
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return &container.Container{}, err
	}
	globalSchemaCont = con
	return globalSchemaCont, nil
}
