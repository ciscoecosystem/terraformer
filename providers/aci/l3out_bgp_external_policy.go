package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3OutBGPExtPolClass = "bgpExtP"

type L3OutBGPExtPolGenerator struct {
	ACIService
}

func (a *L3OutBGPExtPolGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client:= clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3OutBGPExtPolClass)

	L3OutBGPExtPolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L3OutBGPExtPolCount, err := strconv.Atoi(stripQuotes(L3OutBGPExtPolCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L3OutBGPExtPolCount; i++ {
		L3OutBGPExtPolDN := stripQuotes(L3OutBGPExtPolCont.S("imdata").Index(i).S(L3OutBGPExtPolClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			L3OutBGPExtPolDN,
			L3OutBGPExtPolDN,
			"aci_l3out_bgp_external_policy",
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
