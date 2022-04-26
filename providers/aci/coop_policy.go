package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cOOPGroupPolicyClassName = "coopPol"

type COOPGroupPolicyGenerator struct {
	ACIService
}

func (a *COOPGroupPolicyGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, cOOPGroupPolicyClassName)

	COOPGroupPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	COOPGroupPolicyCount, err := strconv.Atoi(stripQuotes(COOPGroupPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < COOPGroupPolicyCount; i++ {
		COOPGroupPolicyAttr := COOPGroupPolicyCont.S("imdata").Index(i).S(cOOPGroupPolicyClassName, "attributes")
		COOPGroupPolicyDN := G(COOPGroupPolicyAttr, "dn")
		name := G(COOPGroupPolicyAttr, "name")
		if filterChildrenDn(COOPGroupPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				COOPGroupPolicyDN,
				name,
				"aci_coop_policy",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
