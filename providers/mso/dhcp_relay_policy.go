package mso

import (
    "github.com/GoogleCloudPlatform/terraformer/terraformutils"
    "github.com/ciscoecosystem/mso-go-client/client"
    "github.com/ciscoecosystem/mso-go-client/container"
    "github.com/ciscoecosystem/mso-go-client/models"
)

type DhcpRelayPolicyGenerator struct {
    MSOService
}

var globalDhcpRelayCont *container.Container

func (a *DhcpRelayPolicyGenerator) InitResources() error {
    mso, err := a.getClient()
    if err != nil {
        return err
    }
    con, err := getDhcpRelayContainer(mso)
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
            "mso_dhcp_relay_policy",
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

func getDhcpRelayContainer(mso *client.Client) (*container.Container, error) {
    if globalDhcpRelayCont != nil {
        return globalDhcpRelayCont, nil
    }
    con, err := mso.GetViaURL("api/v1/policies/dhcp/relay")
    if err != nil {
        return nil, err
    }
    globalDhcpRelayCont = con
    return globalDhcpRelayCont, nil
}