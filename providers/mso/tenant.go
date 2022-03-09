package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TenantGenerator struct {
	MSOService
}

func (a *TenantGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/tenants")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("tenants").Data().([]interface{})); i++ {
		tenantCont := con.S("tenants").Index(i)
		tenantId := stripQuotes(tenantCont.S("id").String())
		name := stripQuotes(tenantCont.S("name").String())
		displayName := stripQuotes(tenantCont.S("displayName").String())
		var description string
		if tenantCont.Exists("description") {
			description = stripQuotes(tenantCont.S("description").String())
		} else {
			description = ""
		}
		tenantName := tenantId + "_" + name
		resource := terraformutils.NewResource(
			tenantId,
			tenantName,
			"mso_tenant",
			"mso",
			map[string]string{
				"name":         name,
				"display_name": displayName,
				"description":  description,
			},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = SlowQueryRequired
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
