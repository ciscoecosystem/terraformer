package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteAnpEpgSelector struct {
	MSOService
}

func (a *SchemaSiteAnpEpgSelector) InitResources() error {
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
		schemaId := stripQuotes(schemaCont.S("id").String())
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}
		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := models.G(siteCont, "siteId")
			templateName := models.G(siteCont, "templateName")
			anpsLen := 0
			if siteCont.Exists("anps") {
				anpsLen = len(siteCont.S("anps").Data().([]interface{}))
			}
			for k := 0; k < anpsLen; k++ {
				anpCont := siteCont.S("anps").Index(k)
				anpRef := models.G(anpCont, "anpRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				anpName := match[3]
				epgsLen := 0
				if anpCont.Exists("epgs") {
					epgsLen = len(anpCont.S("epgs").Data().([]interface{}))
				}
				for m := 0; m < epgsLen; m++ {
					epgCont := anpCont.S("epgs").Index(m)
					epgRef := models.G(epgCont, "epgRef")
					re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)/epgs/(.*)")
					epgMatch := re.FindStringSubmatch(epgRef)
					epgName := epgMatch[4]
					selectorLen := 0
					if epgCont.Exists("selectors") {
						selectorLen = len(epgCont.S("selectors").Data().([]interface{}))
					}
					for n := 0; n < selectorLen; n++ {
						selectorCont := epgCont.S("selectors").Index(n)
						selectorName := models.G(selectorCont, "name")
						name := schemaId + "_" + siteId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + selectorName
						resource := terraformutils.NewResource(
							selectorName,
							name,
							"mso_schema_site_anp_epg_selector",
							"mso",
							map[string]string{
								"schema_id":     schemaId,
								"site_id":       siteId,
								"anp_name":      anpName,
								"epg_name":      epgName,
								"name":          selectorName,
								"template_name": templateName,
							},
							[]string{},
							map[string]interface{}{},
						)
						resource.SlowQueryRequired = SlowQueryRequired
						a.Resources = append(a.Resources, resource)
					}
				}
			}
		}
	}
	return nil
}
