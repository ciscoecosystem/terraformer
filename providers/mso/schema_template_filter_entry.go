package mso

import (
	"math/rand"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateFilterEntryGenerator struct {
	MSOService
}

func (a *SchemaTemplateFilterEntryGenerator) InitResources() error {
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
			filterLen := 0
			if templateCont.Exists("filters") {
				filterLen = len(templateCont.S("filters").Data().([]interface{}))
			}
			for k := 0; k < filterLen; k++ {
				filterCont := templateCont.S("filters").Index(k)
				filterName := models.G(filterCont, "name")
				filterDisplayName := models.G(filterCont, "displayName")
				entryLen := 0
				if filterCont.Exists("entries") {
					entryLen = len(filterCont.S("entries").Data().([]interface{}))
				}
				for l := 0; l < entryLen; l++ {
					entryCont := filterCont.S("entries").Index(l)
					entryName := models.G(entryCont, "name")
					entryDisplayName := models.G(entryCont, "displayName")
					name := schemaId + "_" + templateName + "_" + filterName + "_" + entryName + "_" + strconv.Itoa(rand.Intn(1000))
					resource := terraformutils.NewResource(
						schemaId,
						name,
						"mso_schema_template_filter_entry",
						"mso",
						map[string]string{
							"schema_id":          schemaId,
							"template_name":      templateName,
							"name":               filterName,
							"display_name":       filterDisplayName,
							"entry_name":         entryName,
							"entry_display_name": entryDisplayName,
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
