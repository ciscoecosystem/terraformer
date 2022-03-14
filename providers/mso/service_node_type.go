package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceNodeType struct {
	MSOService
}

func (a *ServiceNodeType) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas/service-node-types")
	if err != nil {
		return err
	}
	serviceNodesCont := con.S("serviceNodeTypes")
	serviceNodeCount := len(serviceNodesCont.Data().([]interface{}))

	for i := 0; i < serviceNodeCount; i++ {
		serviceNodeCont := serviceNodesCont.Index(i)
		serviceNodeID := stripQuotes(serviceNodeCont.S("id").String())
		serviceNodeName := stripQuotes(serviceNodeCont.S("name").String())
		serviceNodeDisplayName := stripQuotes(serviceNodeCont.S("displayName").String())
		resourceName := serviceNodeID + "_" + serviceNodeName
		resource := terraformutils.NewResource(
			serviceNodeID,
			resourceName,
			"mso_service_node_type",
			"mso",
			map[string]string{
				"name":         serviceNodeName,
				"display_name": serviceNodeDisplayName,
			},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
