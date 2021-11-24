package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const logicalDeviceContextClassName = "vnsLDevCtx"

type LogicalDeviceContextGenerator struct {
	ACIService
}

func (a *LogicalDeviceContextGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, logicalDeviceContextClassName)

	LogicalDeviceContextCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LogicalDeviceContextCount, err := strconv.Atoi(stripQuotes(LogicalDeviceContextCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LogicalDeviceContextCount; i++ {
		LogicalDeviceContextAttr := LogicalDeviceContextCont.S("imdata").Index(i).S(logicalDeviceContextClassName, "attributes")
		LogicalDeviceContextDN := G(LogicalDeviceContextAttr, "dn")
		ctrctNameOrLbl := G(LogicalDeviceContextAttr, "ctrctNameOrLbl")
		graphNameOrLbl := G(LogicalDeviceContextAttr, "graphNameOrLbl")
		nodeNameOrLbl := G(LogicalDeviceContextAttr, "nodeNameOrLbl")
		if filterChildrenDn(LogicalDeviceContextDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				LogicalDeviceContextDN,
				resourceNamefromDn(logicalDeviceContextClassName, (LogicalDeviceContextDN), i),
				"aci_logical_device_context",
				"aci",
				map[string]string{
					"tenant_dn": GetParentDn(LogicalDeviceContextDN, fmt.Sprintf("/ldevCtx-c-%s-g-%s-n-%s", ctrctNameOrLbl, graphNameOrLbl, nodeNameOrLbl)),
				},
				[]string{
					"description",
					"name_alias",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
