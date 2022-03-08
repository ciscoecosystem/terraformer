package mso

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateANPEPGContractGenerator struct {
	MSOService
}

func (a *SchemaTemplateANPEPGContractGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := models.G(schemaCont, "id")
		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := models.G(templateCont, "name")
			anpLen := 0
			if templateCont.Exists("anps") {
				anpLen = len(templateCont.S("anps").Data().([]interface{}))
			}
			for k := 0; k < anpLen; k++ {
				anpCont := templateCont.S("anps").Index(k)
				anpName := models.G(anpCont, "name")
				epgLen := 0
				if anpCont.Exists("epgs") {
					epgLen = len(anpCont.S("epgs").Data().([]interface{}))
				}
				for l := 0; l < epgLen; l++ {
					epgCont := anpCont.S("epgs").Index(l)
					epgName := models.G(epgCont, "name")
					contractLen := 0
					if epgCont.Exists("contractRelationships") {
						contractLen = len(epgCont.S("contractRelationships").Data().([]interface{}))
					}
					for m := 0; m < contractLen; m++ {
						contractCont := epgCont.S("contractRelationships").Index(m)
						contractRef := models.G(contractCont, "contractRef")
						contractRefSplitted := strings.Split(contractRef, "/")
						contractName := contractRefSplitted[len(contractRefSplitted)-1]
						relationShipType := models.G(contractCont, "relationshipType")
						name := schemaId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + contractName
						resource := terraformutils.NewResource(
							schemaId,
							name,
							"mso_schema_template_anp_epg_contract",
							"mso",
							map[string]string{
								"schema_id":         schemaId,
								"template_name":     templateName,
								"anp_name":          anpName,
								"epg_name":          epgName,
								"contract_name":     contractName,
								"relationship_type": relationShipType,
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
	}
	return nil
}
