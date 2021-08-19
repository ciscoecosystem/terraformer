package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3ExtSubnetClass = "l3extSubnet"

type L3ExtSubnetGenerator struct {
	ACIService
}

func (a *L3ExtSubnetGenerator) InitResources() error {
	client, err := a.createClient()

	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3ExtSubnetClass)

	l3ExtSubnetCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	l3ExtSubnetCount, err := strconv.Atoi(stripQuotes(l3ExtSubnetCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3ExtSubnetCount; i++ {
		L3ExtSubnetDN := stripQuotes(l3ExtSubnetCont.S("imdata").Index(i).S(L3ExtSubnetClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			L3ExtSubnetDN,
			L3ExtSubnetDN,
			"aci_l3_ext_subnet",
			"aci",
			[]string{
				"aggregate",
				"name_alias",
				"scope",
				"relation_l3ext_rs_subnet_to_profile",
				"relation_l3ext_rs_subnet_to_rt_summ",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil

}
