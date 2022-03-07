package mso

import (
	"math/rand"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateVRFGenerator struct {
	MSOService
}

func (a *SchemaTemplateVRFGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := models.G(schemaCont, "id")
		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := models.G(templateCont, "name")
			vrfLen := 0
			if templateCont.Exists("vrfs") {
				vrfLen = len(templateCont.S("vrfs").Data().([]interface{}))
			}
			for k := 0; k < vrfLen; k++ {
				vrfCont := templateCont.S("vrfs").Index(k)
				vrfName := models.G(vrfCont, "name")
				name := schemaId + "_" + templateName + "_" + vrfName + "_" + strconv.Itoa(rand.Intn(1000))
				resource := terraformutils.NewResource(
					vrfName,
					name,
					"mso_schema_template_vrf",
					"mso",
					map[string]string{
						"schema_id": schemaId,
						"template":  templateName,
						"name":      vrfName,
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
