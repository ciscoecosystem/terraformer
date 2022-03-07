package mso

import (
	"math/rand"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteGenerator struct {
	MSOService
}

func (a *SchemaSiteGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := models.G(schemaCont, "id")
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}
		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := models.G(siteCont, "siteId")
			templateName := models.G(siteCont, "templateName")
			schemaSiteName := siteId + "_" + templateName + "_" + strconv.Itoa(rand.Intn(1000))
			resource := terraformutils.NewResource(
				schemaId,
				schemaSiteName,
				"mso_schema_site",
				"mso",
				map[string]string{
					"schema_id":     schemaId,
					"template_name": templateName,
					"site_id":       siteId,
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
