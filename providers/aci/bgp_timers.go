package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BgpTimersClass = "bgpCtxPol"

type BgpTimersGenerator struct {
	ACIService
}

func (a *BgpTimersGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BgpTimersClass)

	BgpTimersCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BgpTimersCount, err := strconv.Atoi(stripQuotes(BgpTimersCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BgpTimersCount; i++ {
		BgpTimersDN := stripQuotes(BgpTimersCont.S("imdata").Index(i).S(BgpTimersClass, "attributes", "dn").String())
		if filterChildrenDn(BgpTimersDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				BgpTimersDN,
				fmt.Sprintf("%s_%s_%d", BgpTimersClass, GetMOName(BgpTimersDN), i),
				"aci_bgp_timers",
				"aci",
				[]string{
					"name_alias",
					"gr_ctrl",
					"annotation",
					"description",
					"hold_intvl",
					"ka_intvl",
					"max_as_limit",
					"stale_intvl",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
