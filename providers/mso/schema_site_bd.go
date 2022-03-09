package mso

import (
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	// "github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteBdGenerator struct {
	MSOService
}

func (a *SchemaSiteBdGenerator) InitResources() error {
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
		sitesCon := schemaCon.Index(i).S("sites")
		sitesCount := 0

		if schemaCon.Index(i).Exists("sites") {
			sitesCount = len(schemaCon.Index(i).S("sites").Data().([]interface{}))
		}

		for j := 0; j < sitesCount; j++ {
			siteId := models.G(sitesCon, "siteId")
			templateName := models.G(sitesCon, "templateName")
			hostBasedRouting := models.G(sitesCon, "hostBasedRouting")

			bdCount := 0
			bdCont := sitesCon.Index(i).S("bds")

			if sitesCon.Index(i).Exists("bds") {
				bdCount = len(sitesCon.Index(j).S("bds").Data().([]interface{}))
			}

			for k := 0; k < bdCount; k++ {
				bdRef := models.G(bdCont, "bdRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
				match := re.FindStringSubmatch(bdRef)

				bdName := match[3]

				name := schemaId + "_" + siteId + "_" + templateName + "_" + bdName + "_" + hostBasedRouting

				fmt.Printf("name: %v\n", name)
				resource := terraformutils.NewResource(
					bdName,
					name,
					"mso_schema_site_bd",
					"mso",
					map[string]string{
						"site_id":       siteId,
						"template_name": templateName,
						"schema_id":     schemaId,
						"bd_name":       bdName,
						"host_route":    hostBasedRouting,
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
