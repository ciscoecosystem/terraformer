package mso

import (
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	// "github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSitel3OutsGenerator struct {
	MSOService
}

func (a *SchemaSitel3OutsGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLength := len(con.S("schemas").Data().([]interface{}))
	schemaCon := con.S("schemas")
	for i := 0; i < schemaLength; i++ {

		schemaId := models.G(schemaCon, "id")

		sitesCount := 0

		if schemaCon.Index(i).Exists("sites") {
			sitesCount = len(schemaCon.Index(i).S("sites").Data().([]interface{}))
		}

		sitesCon := schemaCon.Index(i).S("sites")

		for j := 0; j < sitesCount; j++ {

			siteId := models.G(sitesCon, "siteId")
			templateName := models.G(sitesCon, "templateName")

			bdCount := 0
			bdCont := sitesCon.Index(j).S("bds")

			if sitesCon.Index(j).Exists("bds") {
				bdCount = len(sitesCon.Index(j).S("bds").Data().([]interface{}))
			}

			for k := 0; k < bdCount; k++ {
				bdRef := models.G(bdCont, "bdRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
				match := re.FindStringSubmatch(bdRef)

				bdName := match[3]

				l3OutsCount := 0

				if bdCont.Index(k).Exists("l3Outs") {
					l3OutsCount = len(bdCont.Index(k).S("l3Outs").Data().([]interface{}))
				}

				for l := 0; l < l3OutsCount; l++ {
					l3OutName := stripQuotes(bdCont.Index(k).S("l3Outs").Index(l).String())

					name := schemaId + "_" + siteId + "_" + templateName + "_" + bdName + "_" + l3OutName

					fmt.Printf("name: %v\n", name)

					resource := terraformutils.NewResource(
						l3OutName,
						name,
						"mso_schema_site_bd_l3out",
						"mso",
						map[string]string{
							"site_id":       siteId,
							"template_name": templateName,
							"schema_id":     schemaId,
							"bd_name":       bdName,
							"l3out_name":    l3OutName,
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
