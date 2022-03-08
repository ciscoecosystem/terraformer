package mso

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateAnpEpgGenerator struct {
	MSOService
}

func (a *SchemaTemplateAnpEpgGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
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
			anpCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Data().([]interface{}))

			for k := 0; k < anpCount; k++ {
				anpNname := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("name").String())
				epgCount := len(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Data().([]interface{}))

				for l := 0; l < epgCount; l++ {
					epgName := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("name").String())
					epgID := fmt.Sprintf("/schemas/%s/templates/%s/anps/%s/epgs/%s", schemaId, templateName, anpNname, epgName)
					name := schemaId + "_" + templateName + "_" + anpNname + "_" + epgName
					if stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("vrfRef").String()) != "" && stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("bdRef").String()) != "" {
						vrfRef := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("vrfRef").String())
						bdRef := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("bdRef").String())
						vrfRefSplitted := strings.Split(vrfRef, "/")
						bdRefSplitted := strings.Split(bdRef, "/")

						resource := terraformutils.NewResource(
							epgID,
							name,
							"mso_schema_template_anp_epg",
							"mso",
							map[string]string{
								"template_name":     templateName,
								"schema_id":         schemaId,
								"anp_name":          anpNname,
								"name":              epgName,
								"bd_name":           bdRefSplitted[6],
								"bd_schema_id":      bdRefSplitted[2],
								"bd_template_name":  bdRefSplitted[4],
								"vrf_name":          vrfRefSplitted[6],
								"vrf_schema_id":     vrfRefSplitted[2],
								"vrf_template_name": vrfRefSplitted[4],
							},
							[]string{},
							map[string]interface{}{},
						)
						resource.SlowQueryRequired = true
						a.Resources = append(a.Resources, resource)

					} else if stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("vrfRef").String()) != "" {
						vrfRef := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("vrfRef").String())
						vrfRefSplitted := strings.Split(vrfRef, "/")
						resource := terraformutils.NewResource(
							epgID,
							name,
							"mso_schema_template_anp_epg",
							"mso",
							map[string]string{
								"template_name":     templateName,
								"schema_id":         schemaId,
								"anp_name":          anpNname,
								"name":              epgName,
								"vrf_name":          vrfRefSplitted[6],
								"vrf_schema_id":     vrfRefSplitted[2],
								"vrf_template_name": vrfRefSplitted[4],
							},
							[]string{},
							map[string]interface{}{},
						)
						resource.SlowQueryRequired = true
						a.Resources = append(a.Resources, resource)

					} else if stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("bdRef").String()) != "" {
						bdRef := stripQuotes(con.S("schemas").Index(i).S("templates").Index(j).S("anps").Index(k).S("epgs").Index(l).S("bdRef").String())
						bdRefSplitted := strings.Split(bdRef, "/")
						resource := terraformutils.NewResource(
							epgID,
							name,
							"mso_schema_template_anp_epg",
							"mso",
							map[string]string{
								"template_name":    templateName,
								"schema_id":        schemaId,
								"anp_name":         anpNname,
								"name":             epgName,
								"bd_name":          bdRefSplitted[6],
								"bd_schema_id":     bdRefSplitted[2],
								"bd_template_name": bdRefSplitted[4],
							},
							[]string{},
							map[string]interface{}{},
						)
						resource.SlowQueryRequired = true
						a.Resources = append(a.Resources, resource)
					} else {
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
	}
	return nil
}
