package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const connectionClassName = "vnsAbsConnection"

type ConnectionGenerator struct {
	ACIService
}

func (a *ConnectionGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, connectionClassName)

	ConnectionCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ConnectionCount, err := strconv.Atoi(stripQuotes(ConnectionCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ConnectionCount; i++ {
		ConnectionAttr := ConnectionCont.S("imdata").Index(i).S(connectionClassName, "attributes")
		ConnectionDN := G(ConnectionAttr, "dn")
		name := G(ConnectionAttr, "name")
		if filterChildrenDn(ConnectionDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				ConnectionDN,
				resourceNamefromDn(connectionClassName, (ConnectionDN), i),
				"aci_connection",
				"aci",
				map[string]string{
					"l4_l7_service_graph_template_dn": GetParentDn(ConnectionDN, fmt.Sprintf("/AbsConnection-%s", name)),
				},
				[]string{
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
