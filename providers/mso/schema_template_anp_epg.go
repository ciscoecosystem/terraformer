package mso

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type SchemaTemplateAnpEpgGenerator struct {
	MSOService
}

func (a *SchemaTemplateAnpEpgGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLength := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLength; i++ {
		temPcount := len(con.S("schemas").Index(i).S("templates").Data().([]interface{}))
		fmt.Printf("temPcount: %v\n", temPcount)
		schemaId := stripQuotes(con.S("schemas").Index(i).S("id").String())

		for j := 0; j < temPcount; j++ {
			templateName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("name").String())
			fmt.Printf("templateName: %v\n", templateName)
			anpCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Data().([]interface{}))
			fmt.Printf("anpCount: %v\n", anpCount)

			for k := 0; k < anpCount; k++ {
				anpNname := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("name").String())
				fmt.Printf("anpNname: %v\n", anpNname)
				epgCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Data().([]interface{}))
				fmt.Printf("epgCount: %v\n", epgCount)

				for l := 0; l < epgCount; l++ {
					epgName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("name").String())
					fmt.Printf("epgName: %v\n", epgName)
					epgID := fmt.Sprintf("/schemas/%s/templates/%s/anps/%s/epgs/%s", schemaId, templateName, anpNname, epgName)
					fmt.Printf("epgID: %v\n", epgID)
					name := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k) + "_" + strconv.Itoa(l)
					resource := terraformutils.NewResource(
						epgID,
						name,
						"mso_schema_template_anp_epg",
						"mso",
						map[string]string{
							"template_name": templateName,
							"schema_id":     schemaId,
							"anp_name":      anpNname,
							"name":          epgName,
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
