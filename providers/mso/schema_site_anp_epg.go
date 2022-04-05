package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteAnpEpg struct {
	MSOService
}

func (a *SchemaSiteAnpEpg) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}

	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())
		sitesCount := 0
		if schemaCont.Exists("sites") {
			sitesCount = len(schemaCont.S("sites").Data().([]interface{}))
		}

		for j := 0; j < sitesCount; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := stripQuotes(siteCont.S("siteId").String())
			templateName := stripQuotes(schemaCont.S("templateName").String())
			anpCount := 0
			if siteCont.Exists("anps") {
				anpCount = len(siteCont.S("anps").Data().([]interface{}))
			}

			for k := 0; k < anpCount; k++ {
				anpCont := siteCont.S("anps").Index(k)
				anpRef := stripQuotes(anpCont.S("anpRef").String())
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				anpName := match[3]
				epgCount := 0
				if anpCont.Exists("epgs") {
					epgCount = len(anpCont.S("epgs").Data().([]interface{}))
				}

				for l := 0; l < epgCount; l++ {
					epgCont := anpCont.S("epgs").Index(l)
					epgRef := stripQuotes(epgCont.S("epgRef").String())
					epgRe := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)/epgs/(.*)")
					epgMatch := epgRe.FindStringSubmatch(epgRef)
					epgName := epgMatch[4]
					resourceName := schemaId + "_" + siteId + "_" + templateName + "_" + anpName + "_" + epgName
					resource := terraformutils.NewResource(
						epgName,
						resourceName,
						"mso_schema_site_anp_epg",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"epg_name":      epgName,
							"site_id":       siteId,
							"anp_name":      anpName,
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
