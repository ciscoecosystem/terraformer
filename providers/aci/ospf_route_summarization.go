package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const OSPFRouteSumClass = "ospfRtSummPol"

type OSPFRouteSumGenerator struct {
	ACIService
}

func (a *OSPFRouteSumGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, OSPFRouteSumClass)

	OSPFRouteSumCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	OSPFRouteSumConut, err := strconv.Atoi(stripQuotes(OSPFRouteSumCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < OSPFRouteSumConut; i++ {
		OSPFRouteSumDN := stripQuotes(OSPFRouteSumCont.S("imdata").Index(i).S(OSPFRouteSumClass, "attributes", "dn").String())
		if filterChildrenDn(OSPFRouteSumDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				OSPFRouteSumDN,
				resourceNamefromDn(OSPFRouteSumClass, (OSPFRouteSumDN), i),
				"aci_ospf_route_summarization",
				"aci",
				[]string{
					"description",
					"annotation",
					"cost",
					"inter_area_enabled",
					"name_alias",
					"tag",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
