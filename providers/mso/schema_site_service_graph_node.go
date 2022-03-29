package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteServiceGraphNodeGenerator struct {
	MSOService
}

var nodeType = map[string]string{
	"0000ffff0000000000000051": "firewall",
	"0000ffff0000000000000052": "load-balancer",
}

func (a *SchemaSiteServiceGraphNodeGenerator) InitResources() error {
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
		templateLen := 0
		if schemaCont.Exists("templates") {
			templateLen = len(schemaCont.S("templates").Data().([]interface{}))
		}
		for j := 0; j < templateLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := models.G(templateCont, "name")
			graphLen := 0
			if templateCont.Exists("serviceGraphs") {
				graphLen = len(templateCont.S("serviceGraphs").Data().([]interface{}))
			}
			for k := 0; k < graphLen; k++ {
				graphCont := templateCont.S("serviceGraphs").Index(k)
				graphName := models.G(graphCont, "name")
				serviceNodeLen := 0
				if graphCont.Exists("serviceNodes") {
					serviceNodeLen = len(graphCont.S("serviceNodes").Data().([]interface{}))
				}
				for l := 1; l < serviceNodeLen; l++ {
					serviceNodeCont := graphCont.S("serviceNodes").Index(l)
					serviceNodeName := models.G(serviceNodeCont, "name")
					serviceNodeTypeHash := models.G(serviceNodeCont, "serviceNodeTypeId")
					serviceNodeType := ""
					if val, ok := nodeType[serviceNodeTypeHash]; !ok {
						serviceNodeType = "other"
					} else {
						serviceNodeType = val
					}
					name := schemaId + "_" + templateName + "_" + graphName + "_" + serviceNodeName
					resource := terraformutils.NewResource(
						serviceNodeName,
						name,
						"mso_schema_site_service_graph_node",
						"mso",
						map[string]string{
							"schema_id":          schemaId,
							"template_name":      templateName,
							"service_graph_name": graphName,
							"service_node_type":  serviceNodeType,
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
