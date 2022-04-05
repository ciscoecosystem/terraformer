package mso

import (
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateExternalEPG struct {
	MSOService
}

func (a *SchemaTemplateExternalEPG) InitResources() error {
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

		for j := 0; j < len(schemaCont.S("templates").Data().([]interface{})); j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			for k := 0; k < len(templateCont.S("externalEpgs").Data().([]interface{})); k++ {
				externalEPGCont := templateCont.S("externalEpgs").Index(k)
				externalEPGName := stripQuotes(externalEPGCont.S("name").String())
				externalEPGId := fmt.Sprintf("/schemas/%s/templates/%s/externalEpgs/%s", schemaId, templateName, externalEPGName)
				externalEPGDisplayName := stripQuotes(externalEPGCont.S("displayName").String())
				vrfRef := stripQuotes(externalEPGCont.S("vrfRef").String())
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
				match := re.FindStringSubmatch(vrfRef)
				vrfRefName := match[3]
				vrfSchemaID := match[1]
				vrfTemplateName := match[2]
				siteID := make([]interface{}, 0, 1)
				externalEPGType := stripQuotes(externalEPGCont.S("extEpgType").String())
				resourceName := schemaId + "_" + templateName + "_" + externalEPGName
				resource := terraformutils.NewResource(
					externalEPGId,
					resourceName,
					"mso_schema_template_external_epg",
					"mso",
					map[string]string{
						"schema_id":         schemaId,
						"template_name":     templateName,
						"external_epg_name": externalEPGName,
						"display_name":      externalEPGDisplayName,
						"vrf_name":          vrfRefName,
						"vrf_schema_id":     vrfSchemaID,
						"vrf_template_name": vrfTemplateName,
						"external_epg_type": externalEPGType,
					},
					[]string{},
					map[string]interface{}{
						"site_id": siteID,
					},
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)

			}
		}
	}
	return nil
}
