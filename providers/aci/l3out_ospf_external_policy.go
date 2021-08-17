package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3outOspfExternalPolicyClass = "ospfExtP"

type L3outOspfExternalPolicyGenerator struct {
	ACIService
}

func (a *L3outOspfExternalPolicyGenerator) InitResources() error {

	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3outOspfExternalPolicyClass)
	L3outOsfpExternalPolicyCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3outOspfExternalPoliciesCount, err := strconv.Atoi(stripQuotes(L3outOsfpExternalPolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3outOspfExternalPoliciesCount; i++ {
		L3outOspfExternalPolicyDN := stripQuotes(L3outOsfpExternalPolicyCont.S("imdata").Index(i).S(L3outOspfExternalPolicyClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			L3outOspfExternalPolicyDN,
			L3outOspfExternalPolicyDN,
			"aci_l3out_ospf_external_policy",
			"aci",
			[]string{
				"area_cost",
				"area_ctrl",
				"area_id",
				"area_type",
				"multipod_internal",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
