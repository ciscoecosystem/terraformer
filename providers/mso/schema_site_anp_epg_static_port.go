package mso

import (
	"regexp"
	"strconv"

	"math/rand"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteAnpEpgStaticPort struct {
	MSOService
}

func (a *SchemaSiteAnpEpgStaticPort) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas/")
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
			anpLen := 0
			if siteCont.Exists("anps") {
				anpLen = len(siteCont.S("anps").Data().([]interface{}))
			}
			for k := 0; k < anpLen; k++ {
				anpCont := siteCont.S("anps").Index(k)
				anpRef := models.G(anpCont, "anpRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/anps/(.*)")
				match := re.FindStringSubmatch(anpRef)
				anpName := match[3]
				epgLen := 0
				if anpCont.Exists("epgs") {
					epgLen = len(anpCont.S("epgs").Data().([]interface{}))
				}
				for l := 0; l < epgLen; l++ {
					epgCont := anpCont.S("epgs").Index(l)
					epgRef := models.G(epgCont, "epgRef")
					re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/epgs/(.*)")
					match := re.FindStringSubmatch(epgRef)
					epgName := match[3]
					staticPortLen := 0
					if epgCont.Exists("staticPorts") {
						staticPortLen = len(epgCont.S("staticPorts").Data().([]interface{}))
					}
					for m := 0; m < staticPortLen; m++ {
						staticPortCont := epgCont.S("staticPorts").Index(m)
						staticPortType := models.G(staticPortCont, "type")
						staticPortPod := ""
						staticPortLeaf := ""
						staticPortPath := ""
						fex := ""
						staticPath := models.G(staticPortCont, "path")
						reForPort := regexp.MustCompile(`topology\/(.*)\/paths-(.*)\/extpaths-(.*)\/pathep-\[(.*)\]`)
						reForVpc := regexp.MustCompile(`topology\/(.*)\/protpaths-(.*)\/pathep-\[(.*)\]`)
						reDefault := regexp.MustCompile(`topology\/(.*)\/paths-(.*)\/pathep-\[(.*)\]`)
						if match := reForPort.FindStringSubmatch(staticPath); len(match) != 0 {
							staticPortPod = match[1]
							staticPortLeaf = match[2]
							staticPortPath = match[4]
							fex = match[3]
						} else if match := reForVpc.FindStringSubmatch(staticPath); len(match) != 0 {
							staticPortPod = match[1]
							staticPortLeaf = match[2]
							staticPortPath = match[3]
						} else if match := reDefault.FindStringSubmatch(staticPath); len(match) != 0 {
							staticPortPod = match[1]
							staticPortLeaf = match[2]
							staticPortPath = match[3]
						}
						name := schemaId + "_" + siteId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + "port_" + strconv.Itoa(rand.Intn(1000))
						resource := terraformutils.NewResource(
							staticPath,
							name,
							"mso_schema_site_anp_epg_static_port",
							"mso",
							map[string]string{
								"schema_id":     schemaId,
								"template_name": templateName,
								"site_id":       siteId,
								"anp_name":      anpName,
								"epg_name":      epgName,
								"path_type":     staticPortType,
								"pod":           staticPortPod,
								"leaf":          staticPortLeaf,
								"path":          staticPortPath,
								"fex":           fex,
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
