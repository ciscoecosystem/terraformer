package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const LeafInterfaceProfileClass = "infraAccPortP"

type LeafInterfaceProfileGenerator struct {
	ACIService
}

func (a *LeafInterfaceProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, LeafInterfaceProfileClass)

	LeafInterfaceProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LeafInterfaceProfileCount, err := strconv.Atoi(stripQuotes(LeafInterfaceProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LeafInterfaceProfileCount; i++ {
		LeafInterfaceProfileDN := stripQuotes(LeafInterfaceProfileCont.S("imdata").Index(i).S(LeafInterfaceProfileClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			LeafInterfaceProfileDN,
			LeafInterfaceProfileDN,
			"aci_leaf_interface_profile",
			"aci",
			[]string{

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
