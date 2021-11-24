package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutLoopbackInterfaceProClass = "l3extLoopBackIfP"

type L3OutLoopbackInterfaceProGenerator struct {
	ACIService
}

func (a *L3OutLoopbackInterfaceProGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutLoopbackInterfaceProClass)

	L3OutLoopbackInterfaceProCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3OutLoopbackInterfaceProCount, err := strconv.Atoi(stripQuotes(L3OutLoopbackInterfaceProCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3OutLoopbackInterfaceProCount; i++ {
		L3OutLoopbackInterfaceProDN := stripQuotes(L3OutLoopbackInterfaceProCont.S("imdata").Index(i).S(L3OutLoopbackInterfaceProClass, "attributes", "dn").String())
		if filterChildrenDn(L3OutLoopbackInterfaceProDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3OutLoopbackInterfaceProDN,
				resourceNamefromDn(L3OutLoopbackInterfaceProClass, L3OutLoopbackInterfaceProDN, i),
				"aci_l3out_loopback_interface_profile",
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
	}
	return nil
}
