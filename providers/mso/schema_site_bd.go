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
	con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}
	schemaLength := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLength; i++ {

		schemaCon := con.S("schemas").Index(i)
		schemaId := models.G(schemaCon, "id")
		sitesCount := 0

		if schemaCon.Exists("sites") {
			sitesCount = len(schemaCon.S("sites").Data().([]interface{}))
		}

		for j := 0; j < sitesCount; j++ {
			sitesCon := schemaCon.S("sites").Index(j)
			siteId := models.G(sitesCon, "siteId")
			templateName := models.G(sitesCon, "templateName")
			bdCount := 0
			if sitesCon.Exists("bds") {
				bdCount = len(sitesCon.S("bds").Data().([]interface{}))
			}

			for k := 0; k < bdCount; k++ {
				bdCont := sitesCon.S("bds").Index(k)
				bdRef := models.G(bdCont, "bdRef")
				hostBasedRouting := bdCont.S("hostBasedRouting").Data().(bool)
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
				match := re.FindStringSubmatch(bdRef)

				bdName := match[3]

				name := schemaId + "_" + siteId + "_" + templateName + "_" + bdName

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
					},
					[]string{},
					map[string]interface{}{
						"host_route": hostBasedRouting,
					},
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
