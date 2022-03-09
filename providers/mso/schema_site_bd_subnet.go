package mso

import (
	"regexp"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteBDSubnetGenerator struct {
	MSOService
}

func (a *SchemaSiteBDSubnetGenerator) InitResources() error {
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
			bdLen := 0
			if siteCont.Exists("bds") {
				bdLen = len(siteCont.S("bds").Data().([]interface{}))
			}
			for k := 0; k < bdLen; k++ {
				bdCont := siteCont.S("bds").Index(k)
				bdRef := models.G(bdCont, "bdRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
				match := re.FindStringSubmatch(bdRef)
				bdName := match[3]
				subnetLen := 0
				if bdCont.Exists("subnets") {
					subnetLen = len(bdCont.S("subnets").Data().([]interface{}))
				}
				for l := 0; l < subnetLen; l++ {
					subnetCont := bdCont.S("subnets").Index(l)
					ip := models.G(subnetCont, "ip")
					name := schemaId + "_" + templateName + "_" + siteId + "_" + bdName + "_" + strconv.Itoa(int(hash(ip)))
					resource := terraformutils.NewResource(
						ip,
						name,
						"mso_schema_site_bd_subnet",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"site_id":       siteId,
							"bd_name":       bdName,
							"ip":            ip,
						},
						[]string{},
						map[string]interface{}{},
					)
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}
			}
		}
	}
	return nil
}
