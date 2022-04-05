package mso

import (
	"regexp"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaTemplateContractServiceGraph struct {
	MSOService
}

func (a *SchemaTemplateContractServiceGraph) InitResources() error {
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
		schemaId := models.G(schemaCont, "id")
		siteLen := 0
		siteId := ""
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}
		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId = stripQuotes(siteCont.S("id").String())
			templateName := stripQuotes(siteCont.S("name").String())
			contractsLen := 0
			if siteCont.Exists("contracts") {
				contractsLen = len(siteCont.S("contracts").Data().([]interface{}))
			}
			for k := 0; k < contractsLen; k++ {
				contractCont := siteCont.S("contracts").Index(k)
				contractName := stripQuotes(contractCont.S("name").String())
				serviceGraphLen := 0
				if contractCont.Exists("serviceGraphRelationship") {
					serviceGraphLen = len(contractCont.S("serviceGraphRelationship").Data().([]interface{}))
				}
				for l := 0; l < serviceGraphLen; l++ {
					serviceGraphCont := contractCont.S("serviceGraphRelationship").Index(l)
					serviceGraphName := ""
					if serviceGraphCont.Exists("serviceGraphRef") {
						serviceGraphRef := stripQuotes(serviceGraphCont.S("serviceGraphRef").String())
						reServiceGraph := regexp.MustCompile("/schemas/(.*)/templates/(.*)/serviceGraphs/(.*)")
						matchServiceGraph := reServiceGraph.FindStringSubmatch(serviceGraphRef)
						serviceGraphName = matchServiceGraph[3]
					}
					serviceNodeLen := 0
					if serviceGraphCont.Exists("serviceNodesRelationship") {
						serviceNodeLen = len(serviceGraphCont.S("serviceNodesRelationship").Data().([]interface{}))
					}
					nodeRelationships := make([]map[string]string, 0, 1)
					for m := 0; m < serviceNodeLen; m++ {
						serviceNodeCont := serviceGraphCont.S("serviceNodesRelationship").Index(m)
						providerbdName := ""
						consumerbdName := ""
						if serviceNodeCont.Exists("providerConnector") {

							providerbdRefCont := serviceNodeCont.Index(m).S("providerConnector")
							providerbdRefMap := providerbdRefCont.Data().(map[string]interface{})
							bdRefName := providerbdRefMap["bdRef"].(string)
							reproviderbdRef := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
							matchproviderbdRef := reproviderbdRef.FindStringSubmatch(bdRefName)
							providerbdName = matchproviderbdRef[3]
						}
						if serviceNodeCont.Exists("consumerConnector") {
							consumerbdRefCont := serviceNodeCont.Index(m).S("consumerConnector")
							consumerbdRefMap := consumerbdRefCont.Data().(map[string]interface{})
							bdRefName := consumerbdRefMap["bdRef"].(string)
							reconsumerbdRef := regexp.MustCompile("/schemas/(.*)/templates/(.*)/bds/(.*)")
							matchconsumerbdRef := reconsumerbdRef.FindStringSubmatch(bdRefName)
							consumerbdName = matchconsumerbdRef[3]
						}
						nodeRelationship := make(map[string]string, 0)
						nodeRelationship["provider_connector_bd_name"] = providerbdName
						nodeRelationship["consumer_connector_bd_name"] = consumerbdName
						nodeRelationships = append(nodeRelationships, nodeRelationship)
					}
					name := schemaId + "_" + siteId + "_" + templateName + "_" + contractName + "_" + serviceGraphName
					resource := terraformutils.NewResource(
						serviceGraphName,
						name,
						"mso_schema_template_contract_service_graph",
						"mso",
						map[string]string{
							"site_id":                siteId,
							"template_name":          templateName,
							"schema_id":              schemaId,
							"contract_name":          contractName,
							"service_graph_ref_name": serviceGraphName,
						},
						[]string{},
						map[string]interface{}{
							"node_relationship": nodeRelationships,
						},
					)
					resource.SlowQueryRequired = true
					a.Resources = append(a.Resources, resource)
				}

			}

		}

	}
	return nil
}