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
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

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
		SubnetDN := stripQuotes(SubnetsCont.S("imdata").Index(i).S(subnetClassName, "attributes", "dn").String())
		nameAlias := stripQuotes(SubnetsCont.S("imdata").Index(i).S(subnetClassName, "attributes", "nameAlias").String())
		if filterChildrenDn(SubnetDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				SubnetDN,
				fmt.Sprintf("%s_%s_%d", subnetClassName, nameAlias, i),
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
	}
	return nil
}
