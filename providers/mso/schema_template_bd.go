package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateBdGenerator struct {
	MSOService
}

func (a *SchemaTemplateBdGenerator) InitResources() error {
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
		schemaId := models.G(schemaCont, "id")
		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := models.G(templateCont, "name")
			bdsLen := 0
			if templateCont.Exists("bds") {
				bdsLen = len(templateCont.S("bds").Data().([]interface{}))
			}
			for k := 0; k < bdsLen; k++ {
				bds := templateCont.S("bds").Index(k)
				name := models.G(bds, "name")
				optimizeWanBandwidth := models.G(bds, "optimizeWanBandwidth")
				l2Stretch := models.G(bds, "l2Stretch")
				l3MCast := models.G(bds, "l3MCast")
				unicastRouting := "false"
				if bds.Exists("unicastRouting") {
					unicastRouting = models.G(bds, "unicastRouting")
				}
				arpFlood := "false"
				if bds.Exists("arpFlood") {
					arpFlood = models.G(bds, "arpFlood")
				}
				resourceName := schemaId + "_" + templateName + "_" + name
				resource := terraformutils.NewResource(
					name,
					resourceName,
					"mso_schema_template_bd",
					"mso",
					map[string]string{
						"template_name":          templateName,
						"name":                   name,
						"schema_id":              schemaId,
						"optimize_wan_bandwidth": optimizeWanBandwidth,
						"layer2_stretch":         l2Stretch,
						"layer3_multicast":       l3MCast,
						"arp_flooding":           arpFlood,
						"unicast_routing":        unicastRouting,
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
