package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const LogicalNodeToFabricNodeClass = "l3extRsNodeL3OutAtt"

type LogicalNodeToFabricNodeGenerator struct {
	ACIService
}

func (a *LogicalNodeToFabricNodeGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, LogicalNodeToFabricNodeClass)
	LogicalNodeToFabricNodeCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	LogicalNodeToFabricNodesCount, err := strconv.Atoi(stripQuotes(LogicalNodeToFabricNodeCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < LogicalNodeToFabricNodesCount; i++ {
		LogicalNodeToFabricNodeDN := stripQuotes(LogicalNodeToFabricNodeCont.S("imdata").Index(i).S(LogicalNodeToFabricNodeClass, "attributes", "dn").String())
		if filterChildrenDn(LogicalNodeToFabricNodeDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				LogicalNodeToFabricNodeDN,
				LogicalNodeToFabricNodeDN,
				"aci_logical_node_to_fabric_node",
				"aci",
				[]string{
					"config_issues",
					"rtr_id",
					"rtr_id_loop_back",
					"annotation",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
