package mso

import (
	"regexp"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteAnpEpgStaticLeaf struct {
	MSOService
}

func (a *SchemaSiteAnpEpgStaticLeaf) InitResources() error {
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
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}

		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := models.G(siteCont, "siteId")

			anpsLen := 0
			if siteCont.Exists("anps") {
				anpsLen = len(siteCont.S("anps").Data().([]interface{}))
			}

			for k := 0; k < anpsLen; k++ {
				anpCont := siteCont.S("anps").Index(k)
				anpRef := models.G(anpCont, "anpRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				epgsLen := 0
				if anpCont.Exists("epgs") {
					epgsLen = len(anpCont.S("epgs").Data().([]interface{}))
				}

				for m := 0; m < epgsLen; m++ {
					epgCont := anpCont.S("epgs").Index(m)
					epgRef := models.G(epgCont, "epgRef")
					re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)/epgs/(.*)")
					epgMatch := re.FindStringSubmatch(epgRef)
					staticLeafsLen := 0
					if epgCont.Exists("staticLeafs") {
						staticLeafsLen = len(epgCont.S("staticLeafs").Data().([]interface{}))
					}
					for n := 0; n < staticLeafsLen; n++ {
						staticLeafCont := epgCont.S("staticLeafs").Index(n)
						path := models.G(staticLeafCont, "path")
						port, _ := strconv.Atoi(staticLeafCont.S("portEncapVlan").String())
						resourceName := schemaId + "_" + siteId + "_" + match[3] + "_" + epgMatch[4] + "_" + path
						resource := terraformutils.NewResource(
							path,
							resourceName,
							"mso_schema_site_anp_epg_static_leaf",
							"mso",
							map[string]string{
								"schema_id": schemaId,
								"site_id":   siteId,
								"anp_name":  match[3],
								"epg_name":  epgMatch[4],
								"path":      path,
							},
							[]string{},
							map[string]interface{}{
								"port_encap_vlan": port,
							},
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
