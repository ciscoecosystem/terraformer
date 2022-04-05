package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteVRF struct {
	MSOService
}

func (a *SchemaSiteVRF) InitResources() error {
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

			for k := 0; k < len(siteCont.S("vrfs").Data().([]interface{})); k++ {
				vrfCont := siteCont.S("vrfs").Index(k)
				vrfRef := stripQuotes(vrfCont.S("vrfRef").String())
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
				match := re.FindStringSubmatch(vrfRef)
				vrfName := match[3]
				resourceName := schemaId + "_" + siteId + "_" + templateName + "_" + vrfName
				resource := terraformutils.NewResource(
					vrfName,
					resourceName,
					"mso_schema_site_vrf",
					"mso",
					map[string]string{
						"schema_id":     schemaId,
						"template_name": templateName,
						"vrf_name":      vrfName,
						"site_id":       siteId,
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
