package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tenantClassName = "fvTenant"

type TenantGenerator struct {
	ACIService
}

func (a *TenantGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client:= clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tenantClassName)

	tenantsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	tenantCount, err := strconv.Atoi(stripQuotes(tenantsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < tenantCount; i++ {
		tenantDN := tenantsCont.S("imdata").Index(i).S(tenantClassName, "attributes", "dn").String()

		resource := terraformutils.NewSimpleResource(
			stripQuotes(tenantDN),
			GetMOName(tenantDN),
			"aci_tenant",
			"aci",
			[]string{
				"name_alias",
				"relation_fv_rs_tn_deny_rule",
				"relation_fv_rs_tenant_mon_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
