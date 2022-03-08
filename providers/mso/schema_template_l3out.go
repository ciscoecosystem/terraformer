package mso

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateL3OutGenerator struct {
	MSOService
}

func (a *SchemaTemplateL3OutGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := getSchemaContainer(mso)
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
			l3outLen := 0
			if templateCont.Exists("intersiteL3outs") {
				l3outLen = len(templateCont.S("intersiteL3outs").Data().([]interface{}))
			}
			for k := 0; k < l3outLen; k++ {
				l3outCont := templateCont.S("intersiteL3outs").Index(k)
				l3outName := models.G(l3outCont, "name")
				displayName := models.G(l3outCont, "displayName")
				vrfRef := models.G(l3outCont, "vrfRef")
				vrfRefSplitted := strings.Split(vrfRef, "/")
				vrfName := vrfRefSplitted[len(vrfRefSplitted)-1]
				name := schemaId + "_" + templateName + "_" + vrfName + "_" + l3outName
				resource := terraformutils.NewResource(
					l3outName,
					name,
					"mso_schema_template_l3out",
					"mso",
					map[string]string{
						"schema_id":     schemaId,
						"template_name": templateName,
						"l3out_name":    l3outName,
						"display_name":  displayName,
						"vrf_name":      vrfName,
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
