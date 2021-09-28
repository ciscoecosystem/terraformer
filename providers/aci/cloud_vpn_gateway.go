package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudVPNGatewayClass = "cloudRouterP"

type CloudVPNGatewayGenerator struct {
	ACIService
}

func (a *CloudVPNGatewayGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudVPNGatewayClass)
	CloudVPNGatewayCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudVPNGatewaysCount, err := strconv.Atoi(stripQuotes(CloudVPNGatewayCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudVPNGatewaysCount; i++ {
		CloudVPNGatewayDN := stripQuotes(CloudVPNGatewayCont.S("imdata").Index(i).S(CloudVPNGatewayClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			CloudVPNGatewayDN,
			CloudVPNGatewayDN,
			"aci_cloud_vpn_gateway",
			"aci",
			[]string{
				"name_alias",
				"num_instances",
				"cloud_router_profile_type",
				"relation_cloud_rs_to_vpn_gw_pol",
				"relation_cloud_rs_to_direct_conn_pol",
				"relation_cloud_rs_to_host_router_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
