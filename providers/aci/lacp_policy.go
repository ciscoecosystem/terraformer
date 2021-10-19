package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const LacpPolicyClassName = "lacpLagPol"

type LacpPolicyGenerator struct {
	ACIService
}

func (a *LacpPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, LacpPolicyClassName)

	LacpPolicysCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LacpPolicyCount, err := strconv.Atoi(stripQuotes(LacpPolicysCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LacpPolicyCount; i++ {
		LacpPolicyDN := LacpPolicysCont.S("imdata").Index(i).S(LacpPolicyClassName, "attributes", "dn").String()
		if filterChildrenDn(LacpPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(LacpPolicyDN),
				stripQuotes(LacpPolicyDN),
				"aci_lacp_policy",
				"aci",
				[]string{
					"ctrl",
					"max_links",
					"min_links",
					"mode",
					"name_alias",
					"annotation",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
