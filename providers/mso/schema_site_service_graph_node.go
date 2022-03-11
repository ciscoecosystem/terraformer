package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteServiceGraphNodeGenerator struct {
	MSOService
}

func (a *SchemaSiteServiceGraphNodeGenerator) InitResources() error {
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
		schemaId := models.G(schemaCont, "id")
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}
		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := models.G(siteCont, "siteId")
			templateName := models.G(siteCont, "templateName")
			anpLen := 0
			if siteCont.Exists("anps") {
				anpLen = len(siteCont.S("anps").Data().([]interface{}))
			}
			for k := 0; k < anpLen; k++ {
				anpCont := siteCont.S("anps").Index(k)
				anpRef := models.G(anpCont, "anpRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				anpName := match[3]
				name := schemaId + "_" + templateName + "_" + siteId + "_" + anpName
				resource := terraformutils.NewResource(
					anpName,
					name,
					"mso_schema_site_anp",
					"mso",
					map[string]string{
						"schema_id":     schemaId,
						"template_name": templateName,
						"site_id":       siteId,
						"anp_name":      anpName,
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
