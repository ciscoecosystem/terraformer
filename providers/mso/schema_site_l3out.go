package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	// "github.com/ciscoecosystem/mso-go-client/client"
	// "github.com/ciscoecosystem/mso-go-client/container"
	// "github.com/ciscoecosystem/mso-go-client/models"
	"fmt"
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
		schemaId := stripQuotes(schemaCont.S("id").String())
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}

		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := stripQuotes(siteCont.S("siteId").String())
			templateName := stripQuotes(siteCont.S("templateName").String())

			l3outLen := 0
			if siteCont.Exists("intersiteL3outs") {
				l3outLen = len(schemaCont.S("sites").Data().([]interface{}))
				// fmt.Printf("len: %d\n", l3outLen)
			}
			for k := 0; k < l3outLen; k++ {
				l3outCont := siteCont.S("intersiteL3outs").Index(k)
				if l3outCont.Exists("l3outRef") && l3outCont.Exists("vrfRef") {
					l3outRef := stripQuotes(l3outCont.S("l3outRef").String())

					vrfRef := stripQuotes(l3outCont.S("vrfRef").String())
					// fmt.Printf("Vrf: %s\n", vrfRef)
					// fmt.Printf("l3out: %s\n", l3outRef)
					reVrf := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
					matchVrf := reVrf.FindStringSubmatch(vrfRef)
					// fmt.Printf("matchVrf: %s\n", matchVrf)
					vrfRefName := matchVrf[3]
					rel3out := regexp.MustCompile("/schemas/(.*)/templates/(.*)/l3outs/(.*)")
					// fmt.Printf("vrfstr: %s\n", reVrf)
					matchl3out := rel3out.FindStringSubmatch(l3outRef)
					l3outName := matchl3out[3]
					id := fmt.Sprintf("%s/site/%s/template/%s/vrf/%s/l3out/%s", schemaId, siteId, templateName, vrfRefName, l3outName)
					// fmt.Printf("ID: %s\n", id)
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
