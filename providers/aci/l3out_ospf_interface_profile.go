package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3outOspfInterfaceProfileClass = "ospfIfP"

type L3outOspfInterfaceProfileGenerator struct {
	ACIService
}

func (a *L3outOspfInterfaceProfileGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3outOspfInterfaceProfileClass)
	L3outOspfInterfaceProfileCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3outOspfInterfaceProfilesCount, err := strconv.Atoi(stripQuotes(L3outOspfInterfaceProfileCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3outOspfInterfaceProfilesCount; i++ {
		L3outOspfInterfaceProfileDN := stripQuotes(L3outOspfInterfaceProfileCont.S("imdata").Index(i).S(L3outOspfInterfaceProfileClass, "attributes", "dn").String())
		if filterChildrenDn(L3outOspfInterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3outOspfInterfaceProfileDN,
				L3outOspfInterfaceProfileDN,
				"aci_l3out_ospf_interface_profile",
				"aci",
				[]string{
					"auth_key_id",
					"auth_type",
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
