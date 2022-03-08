package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateContractFilter struct {
	MSOService
}

func (a *SchemaTemplateContractFilter) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())
		templateLen := len(schemaCont.S("templates").Data().([]interface{}))

		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			contractsLen := 0
			if templateCont.Exists("contracts") {
				contractsLen = len(templateCont.S("contracts").Data().([]interface{}))
			}

			for k := 0; k < contractsLen; k++ {
				contractCont := templateCont.S("contracts").Index(k)
				contractRef := models.G(contractCont, "contractRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/contracts/(.*)")
				match := re.FindStringSubmatch(contractRef)
				filterType := models.G(contractCont, "filterType")
				filtersLen := 0
				if contractCont.Exists("filterRelationships") {
					filtersLen = len(contractCont.S("filterRelationships").Data().([]interface{}))
				}

				for m := 0; m < filtersLen; m++ {
					filterCont := contractCont.S("filterRelationships").Index(m)
					var fmatch []string
					if filterCont.Exists("filterRef") {
						filRef := models.G(filterCont, "filterRef")
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/filters/(.*)")
						fmatch = re.FindStringSubmatch(filRef)
					}
					resourceName := schemaId + "_" + templateName + "_" + match[3] + "_" + fmatch[3]
					resource := terraformutils.NewResource(
						match[3],
						resourceName,
						"mso_schema_template_contract_filter",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"contract_name": match[3],
							"filter_type":   filterType,
							"filter_name":   fmatch[3],
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
	return nil
}
