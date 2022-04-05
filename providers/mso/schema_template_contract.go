package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateContractGenerator struct {
	MSOService
}

func (a *SchemaTemplateContractGenerator) InitResources() error {
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
		schemaId := stripQuotes(schemaCont.S("id").String())

		for j := 0; j < len(schemaCont.S("templates").Data().([]interface{})); j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			for k := 0; k < len(templateCont.S("contracts").Data().([]interface{})); k++ {
				contractCont := templateCont.S("contracts").Index(k)
				contractName := stripQuotes(contractCont.S("name").String())
				resourceName := schemaId + "_" + templateName + "_" + contractName
				filterRelationship := make([]map[string]string, 0)
				var resource terraformutils.Resource
				if contractCont.Exists("filterRelationships") {
					for m := 0; m < len(contractCont.S("filterRelationships").Data().([]interface{})); m++ {
						filterRef := stripQuotes(contractCont.S("filterRelationships").Index(m).S("filterRef").String())
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/filters/(.*)")
						match := re.FindStringSubmatch(filterRef)
						filterName := match[3]
						filterSchemaId := match[1]
						filterTemplateName := match[2]
						filterMap := map[string]string{
							"filter_name":          filterName,
							"filter_schema_id":     filterSchemaId,
							"filter_template_name": filterTemplateName,
						}
						filterRelationship = append(filterRelationship, filterMap)
					}
					resource = terraformutils.NewResource(
						contractName,
						resourceName,
						"mso_schema_template_contract",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"contract_name": contractName,
						},
						[]string{},
						map[string]interface{}{
							"filter_relationship": filterRelationship,
						},
					)
				} else {
					resource = terraformutils.NewResource(
						contractName,
						resourceName,
						"mso_schema_template_contract",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"contract_name": contractName,
						},
						[]string{},
						map[string]interface{}{},
					)
				}
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
