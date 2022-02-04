package mso

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type SchemaGenerator struct {
	MSOService
}

func (a *SchemaGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaId := stripQuotes(con.S("schemas").Index(i).S("id").String())
		resource := terraformutils.NewResource(
			schemaId,
			strconv.Itoa(i),
			"mso_schema",
			"mso",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
