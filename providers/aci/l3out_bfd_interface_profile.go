package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutBFDInterfaceProfileClass = "bfdIfP"

type L3OutBFDInterfaceProfileGenerator struct {
	ACIService
}

func (a *L3OutBFDInterfaceProfileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutBFDInterfaceProfileClass)

	L3OutBFDInterfaceProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutBFDInterfaceProfileCount, err := strconv.Atoi(stripQuotes(L3OutBFDInterfaceProfileCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3OutBFDInterfaceProfileCount; i++ {
		L3OutBFDInterfaceProfileDN := stripQuotes(L3OutBFDInterfaceProfileCont.S("imdata").Index(i).S(L3OutBFDInterfaceProfileClass, "attributes", "dn").String())
		if filterChildrenDn(L3OutBFDInterfaceProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutBFDInterfaceProfileDN,
				resourceNamefromDn(L3OutBFDInterfaceProfileClass, (L3OutBFDInterfaceProfileDN), i),
				"aci_l3out_bfd_interface_profile",
				"aci",
				[]string{
					"key",
					"key_id",
					"name_alias",
					"interface_profile_type",
					"relation_bfd_rs_if_pol",
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
