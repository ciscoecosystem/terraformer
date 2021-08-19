package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BGPRouteSumClass = "bgpRtSummPol"

type BGPRouteSumGenerator struct {
	ACIService
}

func (a *BGPRouteSumGenerator) InitResources() error {

	client, err := a.createClient()

	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BGPRouteSumClass)

	BGPRouteSumCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BGPRouteSumCount, err := strconv.Atoi(stripQuotes(BGPRouteSumCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BGPRouteSumCount; i++ {
		BGPRouteSumDN := stripQuotes(BGPRouteSumCont.S("imdata").Index(i).S(BGPRouteSumClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			BGPRouteSumDN,
			BGPRouteSumDN,
			"aci_bgp_route_summarization",
			"aci",
			[]string{
				"attrmap",
				"ctrl",
				"name_alias",
				"description",
				"annotation",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
