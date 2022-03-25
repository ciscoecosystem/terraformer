package mso

import (
	"regexp"

	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteL3outGenerator struct {
	MSOService
}

func (a *SchemaSiteL3outGenerator) InitResources() error {
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

			l3outLen := 0
			if siteCont.Exists("intersiteL3outs") {
				l3outLen = len(siteCont.S("intersiteL3outs").Data().([]interface{}))
			}
			for k := 0; k < l3outLen; k++ {
				l3outCont := siteCont.S("intersiteL3outs").Index(k)
				if l3outCont.Exists("l3outRef") && l3outCont.Exists("vrfRef") {
					l3outRef := models.G(l3outCont, "l3outRef")
					vrfRef := models.G(l3outCont, "vrfRef")
					reVrf := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
					matchVrf := reVrf.FindStringSubmatch(vrfRef)
					vrfRefName := matchVrf[3]
					rel3out := regexp.MustCompile("/schemas/(.*)/templates/(.*)/l3outs/(.*)")
					matchl3out := rel3out.FindStringSubmatch(l3outRef)
					l3outName := matchl3out[3]
					id := fmt.Sprintf("%s/site/%s/template/%s/vrf/%s/l3out/%s", schemaId, siteId, templateName, vrfRefName, l3outName)
					resource := terraformutils.NewResource(
						id,
						id,
						"mso_schema_site_l3out",
						"mso",
						map[string]string{},
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
