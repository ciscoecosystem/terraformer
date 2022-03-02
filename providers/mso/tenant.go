package mso

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type TenantGenerator struct {
	MSOService
}

func (a *TenantGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/tenants")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("tenants").Data().([]interface{})); i++ {
		tenantCont := con.S("tenants").Index(i)
		tenantId := stripQuotes(tenantCont.S("id").String())
		name := stripQuotes(tenantCont.S("name").String())
		displayName := stripQuotes(tenantCont.S("displayName").String())
		resource := terraformutils.NewResource(
			tenantId,
			strconv.Itoa(i),
			"mso_tenant",
			"mso",
			map[string]string{
				"name":         name,
				"display_name": displayName,
			},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
