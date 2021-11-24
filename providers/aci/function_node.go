package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const functionNodeClassName = "vnsAbsNode"

type FunctionNodeGenerator struct {
	ACIService
}

func (a *FunctionNodeGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, functionNodeClassName)

	FunctionNodeCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FunctionNodeCount, err := strconv.Atoi(stripQuotes(FunctionNodeCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FunctionNodeCount; i++ {
		FunctionNodeAttr := FunctionNodeCont.S("imdata").Index(i).S(functionNodeClassName, "attributes")
		FunctionNodeDN := G(FunctionNodeAttr, "dn")
		name := G(FunctionNodeAttr, "name")
		if filterChildrenDn(FunctionNodeDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				FunctionNodeDN,
				resourceNamefromDn(functionNodeClassName, (FunctionNodeDN), i),
				"aci_function_node",
				"aci",
				map[string]string{
					"l4_l7_service_graph_template_dn": GetParentDn(FunctionNodeDN, fmt.Sprintf("/AbsNode-%s", name)),
					"conn_consumer_dn":                FunctionNodeDN + "/AbsFConn-consumer",
					"conn_provider_dn":                FunctionNodeDN + "/AbsFConn-provider",
				},
				[]string{
					"description",
					"annotation",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
