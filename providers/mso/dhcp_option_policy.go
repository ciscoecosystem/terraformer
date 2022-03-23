package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type DhcpOptionPolicyGenerator struct {
	MSOService
}

var globalDhcpCont *container.Container

func (a *DhcpOptionPolicyGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getDhcpContainer(mso)
	if err != nil {
		return err
	}
	dhcpLength := len(con.S("DhcpRelayPolicies").Data().([]interface{}))
	for i := 0; i < dhcpLength; i++ {
		dhcpCon := con.S("DhcpRelayPolicies").Index(i)
		dhcpId := models.G(dhcpCon, "id")
		resource := terraformutils.NewResource(
			dhcpId,
			dhcpId,
			"mso_dhcp_option_policy",
			"mso",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}

func getDhcpContainer(mso *client.Client) (*container.Container, error) {
	if globalDhcpCont != nil {
		return globalDhcpCont, nil
	}
	con, err := mso.GetViaURL("api/v1/policies/dhcp/option")
	if err != nil {
		return nil, err
	}
	globalDhcpCont = con
	return globalDhcpCont, nil
}
