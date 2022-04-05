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
	con, err := getSchemaContainer(mso)
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
				var filterType string
				filtersLen := 0
				if contractCont.Exists("filterRelationships") {
					filtersLen = len(contractCont.S("filterRelationships").Data().([]interface{}))
				}

				for m := 0; m < filtersLen; m++ {
					filterCont := contractCont.S("filterRelationships").Index(m)
					filterType = "bothWay"
					var fmatch []string
					if filterCont.Exists("filterRef") {
						filRef := models.G(filterCont, "filterRef")
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/filters/(.*)")
						fmatch = re.FindStringSubmatch(filRef)
					}
					resourceName := schemaId + "_" + templateName + "_" + match[3] + "_" + filterType + "_" + fmatch[3]
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
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}

				if contractCont.Exists("filterRelationshipsProviderToConsumer") {
					filtersLen = len(contractCont.S("filterRelationshipsProviderToConsumer").Data().([]interface{}))
				}

				for m := 0; m < filtersLen; m++ {
					filterCont := contractCont.S("filterRelationshipsProviderToConsumer").Index(m)
					filterType = "provider_to_consumer"
					var fmatch []string
					if filterCont.Exists("filterRef") {
						filRef := models.G(filterCont, "filterRef")
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/filters/(.*)")
						fmatch = re.FindStringSubmatch(filRef)
					}
					resourceName := schemaId + "_" + templateName + "_" + match[3] + "_" + filterType + "_" + fmatch[3]
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
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}

				if contractCont.Exists("filterRelationshipsConsumerToProvider") {
					filtersLen = len(contractCont.S("filterRelationshipsConsumerToProvider").Data().([]interface{}))
				}

				for m := 0; m < filtersLen; m++ {
					filterCont := contractCont.S("filterRelationshipsConsumerToProvider").Index(m)
					filterType = "consumer_to_provider"
					var fmatch []string
					if filterCont.Exists("filterRef") {
						filRef := models.G(filterCont, "filterRef")
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/filters/(.*)")
						fmatch = re.FindStringSubmatch(filRef)
					}
					resourceName := schemaId + "_" + templateName + "_" + match[3] + "_" + filterType + "_" + fmatch[3]
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
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}
			}
		}
	}
	return nil
}
