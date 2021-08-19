package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const subnetClassName = "fvSubnet"

type SubnetGenerator struct {
	ACIService
}

func (a *SubnetGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, subnetClassName)

	SubnetsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SubnetCount, err := strconv.Atoi(stripQuotes(SubnetsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SubnetCount; i++ {
		SubnetDN := SubnetsCont.S("imdata").Index(i).S(subnetClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(SubnetDN),
			stripQuotes(SubnetDN),
			"aci_subnet",
			"aci",
			[]string{
				"ctrl",
				"name_alias",
				"preferred",
				"scope",
				"virtual",
				"relation_fv_rs_bd_subnet_to_out",
				"relation_fv_rs_nd_pfx_pol",
				"relation_fv_rs_bd_subnet_to_profile",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
