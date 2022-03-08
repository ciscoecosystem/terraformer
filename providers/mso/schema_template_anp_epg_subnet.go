package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateAnpEpgSubnet struct {
	MSOService
}

func (a *SchemaTemplateAnpEpgSubnet) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())
		templateLen := len(schemaCont.S("templates").Data().([]interface{}))

		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			anpsLen := 0
			if templateCont.Exists("anps") {
				anpsLen = len(templateCont.S("anps").Data().([]interface{}))
			}

			for k := 0; k < anpsLen; k++ {
				anpCont := templateCont.S("anps").Index(k)
				anpName := models.G(anpCont, "name")

				epgsLen := 0
				if anpCont.Exists("epgs") {
					epgsLen = len(anpCont.S("epgs").Data().([]interface{}))
				}

				for m := 0; m < epgsLen; m++ {
					epgCont := anpCont.S("epgs").Index(m)
					epgName := models.G(epgCont, "name")
					subnetLen := 0
					if epgCont.Exists("subnets") {
						subnetLen = len(epgCont.S("subnets").Data().([]interface{}))
					}
					for n := 0; n < subnetLen; n++ {
						subnetCont := epgCont.S("subnets").Index(n)
						ip := models.G(subnetCont, "ip")
						resourceName := schemaId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + ip
						resource := terraformutils.NewResource(
							ip,
							resourceName,
							"mso_schema_template_anp_epg_subnet",
							"mso",
							map[string]string{
								"schema_id": schemaId,
								"template":  templateName,
								"anp_name":  anpName,
								"epg_name":  epgName,
								"ip":        ip,
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
	}
	return nil
}
