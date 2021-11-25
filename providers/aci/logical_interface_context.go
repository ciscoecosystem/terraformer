package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const logicalInterfaceContextClassName = "vnsLIfCtx"

type LogicalInterfaceContextGenerator struct {
	ACIService
}

func (a *LogicalInterfaceContextGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, logicalInterfaceContextClassName)

	LogicalInterfaceContextCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LogicalInterfaceContextCount, err := strconv.Atoi(stripQuotes(LogicalInterfaceContextCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LogicalInterfaceContextCount; i++ {
		LogicalInterfaceContextAttr := LogicalInterfaceContextCont.S("imdata").Index(i).S(logicalInterfaceContextClassName, "attributes")
		LogicalInterfaceContextDN := G(LogicalInterfaceContextAttr, "dn")
		connNameOrLbl := G(LogicalInterfaceContextAttr, "connNameOrLbl")
		if filterChildrenDn(LogicalInterfaceContextDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				LogicalInterfaceContextDN,
				resourceNamefromDn(logicalInterfaceContextClassName, (LogicalInterfaceContextDN), i),
				"aci_logical_interface_context",
				"aci",
				map[string]string{
					"logical_device_context_dn": GetParentDn(LogicalInterfaceContextDN, fmt.Sprintf("/lIfCtx-c-%s", connNameOrLbl)),
				},
				[]string{
					"description",
					"l3_dest",
					"name_alias",
					"permit_log",
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
