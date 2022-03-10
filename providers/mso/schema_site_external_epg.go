package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteExternalEpgGenerator struct {
	MSOService
}

func (a *SchemaSiteExternalEpgGenerator) InitResources() error {
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
			externalEpgLen := 0
			if siteCont.Exists("externalEpgs") {
				externalEpgLen = len(siteCont.S("externalEpgs").Data().([]interface{}))
			}
			for k := 0; k < externalEpgLen; k++ {
				externalEpgCont := siteCont.S("externalEpgs").Index(k)
				externalEpgRef := models.G(externalEpgCont, "externalEpgRef")
				re := regexp.MustCompile("/schemas/(.*?)/templates/(.*?)/externalEpgs/(.*)")
				match := re.FindStringSubmatch(externalEpgRef)
				externalEpgName := match[3]
				name := schemaId + "_" + templateName + "_" + siteId + "_" + externalEpgName
				resource := terraformutils.NewResource(
					externalEpgName,
					name,
					"mso_schema_site_external_epg",
					"mso",
					map[string]string{
						"schema_id":         schemaId,
						"template_name":     templateName,
						"site_id":           siteId,
						"external_epg_name": externalEpgName,
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
