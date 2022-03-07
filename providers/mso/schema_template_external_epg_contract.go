package mso

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type SchemaTemplateExternalEpgContract struct {
	MSOService
}

func (a *SchemaTemplateExternalEpgContract) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLength := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLength; i++ {
		temPcount := len(con.S("schemas").Index(i).S("templates").Data().([]interface{}))
		schemaId := stripQuotes(con.S("schemas").Index(i).S("id").String())

		for j := 0; j < temPcount; j++ {
			templateName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("name").String())
			externalEpgsCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("externalEpgs").Data().([]interface{}))

			for k := 0; k < externalEpgsCount; k++ {
				externalEpgname := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("externalEpgs").Index(k).S("name").String())
				contractRelationshipsCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("externalEpgs").Index(k).S("contractRelationships").Data().([]interface{}))

				for l := 0; l < contractRelationshipsCount; l++ {
					contractRelationshipsName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("externalEpgs").Index(k).S("contractRelationships").Index(l).S("contractRef").String())
					contractRelationshipsNameSplit := strings.Split(contractRelationshipsName, "/")
					contractID := contractRelationshipsNameSplit[6]
					name := schemaId + "_" + templateName + "_" + externalEpgname + "_" + contractID + "_" + strconv.Itoa(rand.Intn(1000))
					resource := terraformutils.NewResource(
						contractID,
						name,
						"mso_schema_template_external_epg_contract",
						"mso",
						map[string]string{
							"template_name":     templateName,
							"schema_id":         schemaId,
							"contract_name":     contractID,
							"external_epg_name": externalEpgname,
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
