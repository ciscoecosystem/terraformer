package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateExternalEPGSubnet struct {
	MSOService
}

func (a *SchemaTemplateExternalEPGSubnet) InitResources() error {
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

				for l := 0; l < len(externalEPGCont.S("subnets").Data().([]interface{})); l++ {
					subnetCont := externalEPGCont.S("subnets").Index(l)
					subnetIP := stripQuotes(subnetCont.S("ip").String())

					resourceName := schemaId + "_" + templateName + "_" + externalEPGName + "_" + subnetIP
					resource := terraformutils.NewResource(
						subnetIP,
						resourceName,
						"mso_schema_template_external_epg_subnet",
						"mso",
						map[string]string{
							"schema_id":         schemaId,
							"template_name":     templateName,
							"external_epg_name": externalEPGName,
							"ip":                subnetIP,
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
