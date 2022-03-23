package mso

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateServiceGraph struct {
	MSOService
}

func (a *SchemaTemplateServiceGraph) InitResources() error {
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

		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}

		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			serviceGraphsLen := 0
			if templateCont.Exists("serviceGraphs") {
				serviceGraphsLen = len(templateCont.S("serviceGraphs").Data().([]interface{}))
			}

			for k := 0; k < serviceGraphsLen; k++ {
				serviceGraphCont := templateCont.S("serviceGraphs").Index(k)
				serviceGraphName := models.G(serviceGraphCont, "name")

				serviceNodesLen := 0
				if serviceGraphCont.Exists("serviceNodes") {
					serviceNodesLen = len(serviceGraphCont.S("serviceNodes").Data().([]interface{}))
				}

				serviceNodeType := ""
				siteParams := make([]interface{}, 0)
				for p := 0; p < serviceNodesLen; p++ {
					serviceNodeCont := serviceGraphCont.S("serviceNodes").Index(p)
					serviceNodeType = models.G(serviceNodeCont, "serviceNodeTypeId")
					serviceNodeName := models.G(serviceNodeCont, "name")
					if serviceNodeType == "0000ffff0000000000000051" {
						serviceNodeType = "firewall"
					} else if serviceNodeType == "0000ffff0000000000000052" {
						serviceNodeType = "load-balancer"
					} else {
						serviceNodeType = "other"
					}

					siteLen := 0
					if schemaCont.Exists("sites") {
						siteLen = len(schemaCont.S("sites").Data().([]interface{}))
					}
					for m := 0; m < siteLen; m++ {
						siteCont := schemaCont.S("sites").Index(m)
						siteServiceGraphsLen := 0
						if siteCont.Exists("serviceGraphs") {
							siteServiceGraphsLen = len(siteCont.S("serviceGraphs").Data().([]interface{}))
						}
						for n := 0; n < siteServiceGraphsLen; n++ {
							siteServiceGraphCont := siteCont.S("serviceGraphs").Index(n)
							siteServiceGraphRef := models.G(siteServiceGraphCont, "serviceGraphRef")
							re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/serviceGraphs/(.*)")
							match := re.FindStringSubmatch(siteServiceGraphRef)
							if match[1] == schemaId && match[2] == templateName && match[3] == serviceGraphName {
								siteServiceNodesLen := 0
								if siteServiceGraphCont.Exists("serviceNodes") {
									siteServiceNodesLen = len(siteServiceGraphCont.S("serviceNodes").Data().([]interface{}))
								}
								siteMap := make(map[string]interface{})
								for p := 0; p < siteServiceNodesLen; p++ {
									siteServiceNodeCont := siteServiceGraphCont.S("serviceNodes").Index(p)
									siteServiceNodeRef := models.G(siteServiceNodeCont, "serviceNodeRef")
									re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/serviceGraphs/(.*)/serviceNodes/(.*)")
									nmatch := re.FindStringSubmatch(siteServiceNodeRef)
									if schemaId == nmatch[1] && templateName == nmatch[2] && serviceGraphName == nmatch[3] && serviceNodeName == nmatch[4] {
										deviceDn := models.StripQuotes(siteServiceNodeCont.S("device", "dn").String())
										dnSplit := strings.Split(deviceDn, "/")
										siteMap["tenant_name"] = strings.Join(strings.Split(dnSplit[1], "-")[1:], "-")
										siteMap["node_name"] = strings.Join(strings.Split(dnSplit[2], "-")[1:], "-")
										siteMap["site_id"] = models.G(siteCont, "siteId")
										siteParams = append(siteParams, siteMap)
										break
									}
								}
							}
						}
					}
				}
				extraAttr := make(map[string]interface{})
				if len(siteParams) != 0 {
					extraAttr["site_nodes"] = siteParams
				}
				resourceName := schemaId + "_" + templateName + "_" + serviceGraphName
				resource := terraformutils.NewResource(
					serviceGraphName,
					resourceName,
					"mso_schema_template_service_graph",
					"mso",
					map[string]string{
						"schema_id":          schemaId,
						"template_name":      templateName,
						"service_graph_name": serviceGraphName,
						"service_node_type":  serviceNodeType,
					},
					[]string{},
					extraAttr,
				)
				resource.SlowQueryRequired = SlowQueryRequired
				a.Resources = append(a.Resources, resource)
			}
		}
	}
	return nil
}
