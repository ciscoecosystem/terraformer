package mso

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteAnpEpgStaticLeaf struct {
	MSOService
}

func (a *SchemaSiteAnpEpgStaticLeaf) InitResources() error {
	mso := a.getClient().(*client.Client)
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
						resourceName := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k) + "_" + strconv.Itoa(m) + "_" + strconv.Itoa(n)
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
