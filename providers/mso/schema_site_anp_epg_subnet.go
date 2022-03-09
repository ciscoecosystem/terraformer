package mso

import (
	"regexp"
	"strconv"

	"math/rand"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteAnpEpgSubnetGenerator struct {
	MSOService
}

func (a *SchemaSiteAnpEpgSubnetGenerator) InitResources() error {
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
					subnetsLen := 0
					if epgCont.Exists("subnets") {
						subnetsLen = len(epgCont.S("subnets").Data().([]interface{}))
					}
					for m := 0; m < subnetsLen; m++ {
						staticPortCont := epgCont.S("subnets").Index(m)
						subnetIP := models.G(staticPortCont, "ip")
						subnetScope := models.G(staticPortCont, "scope")
						subnetShared := models.G(staticPortCont, "shared")
						subnetDescription := models.G(staticPortCont, "description")
						noDefaultGateway := models.G(staticPortCont, "noDefaultGateway")
						querier := models.G(staticPortCont, "querier")

						name := schemaId + "_" + siteId + "_" + templateName + "_" + anpName + "_" + epgName + "_" + "subnet_" + strconv.Itoa(rand.Intn(1000))
						resource := terraformutils.NewResource(
							subnetIP,
							name,
							"mso_schema_site_anp_epg_subnet",
							"mso",
							map[string]string{
								"schema_id":          schemaId,
								"template_name":      templateName,
								"site_id":            siteId,
								"anp_name":           anpName,
								"epg_name":           epgName,
								"ip":                 subnetIP,
								"scope":              subnetScope,
								"shared":             subnetShared,
								"description":        subnetDescription,
								"no_default_gateway": noDefaultGateway,
								"querier":            querier,
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
