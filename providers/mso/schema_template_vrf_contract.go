package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateVrfContractGenerator struct {
	MSOService
}

func (a *SchemaTemplateVrfContractGenerator) InitResources() error {
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
				contractName := ""
				contractSchemaId := ""
				contractTemplateName := ""
				contractType := ""
				contractLen := 0
				if vrfCont.Exists("vzAnyProviderContracts") {
					contractLen = len(vrfCont.S("vzAnyProviderContracts").Data().([]interface{}))
				}
				for l := 0; l < contractLen; l++ {
					contractCont := vrfCont.S("vzAnyProviderContracts").Index(l)
					contractRef := models.G(contractCont, "contractRef")
					re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/contracts/(.*)")
					split := re.FindStringSubmatch(contractRef)
					contractName = split[3]
					contractSchemaId = split[1]
					contractTemplateName = split[2]
					contractType = "provider"
					name := schemaId + "_" + templateName + "_" + vrfName + "_" + contractName + "_" + contractType
					resource := terraformutils.NewResource(
						contractName,
						name,
						"mso_schema_template_vrf_contract",
						"mso",
						map[string]string{
							"schema_id":              schemaId,
							"template_name":          templateName,
							"vrf_name":               vrfName,
							"relationship_type":      contractType,
							"contract_name":          contractName,
							"contract_schema_id":     contractSchemaId,
							"contract_template_name": contractTemplateName,
						},
						[]string{},
						map[string]interface{}{},
					)
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)

				}
				contractLen = 0
				if vrfCont.Exists("vzAnyConsumerContracts") {
					contractLen = len(vrfCont.S("vzAnyConsumerContracts").Data().([]interface{}))
				}
				for l := 0; l < contractLen; l++ {
					contractCont := vrfCont.S("vzAnyConsumerContracts").Index(l)
					contractRef := models.G(contractCont, "contractRef")
					re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/contracts/(.*)")
					split := re.FindStringSubmatch(contractRef)
					contractName = split[3]
					contractSchemaId = split[1]
					contractTemplateName = split[2]
					contractType = "consumer"
					name := schemaId + "_" + templateName + "_" + vrfName + "_" + contractName + "_" + contractType
					resource := terraformutils.NewResource(
						contractName,
						name,
						"mso_schema_template_vrf_contract",
						"mso",
						map[string]string{
							"schema_id":              schemaId,
							"template_name":          templateName,
							"vrf_name":               vrfName,
							"relationship_type":      contractType,
							"contract_name":          contractName,
							"contract_schema_id":     contractSchemaId,
							"contract_template_name": contractTemplateName,
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
	return nil
}
