package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const OSPFTimersClassName = "ospfCtxPol"

type OSPFTimersGenerator struct {
	ACIService
}

func (a *OSPFTimersGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, OSPFTimersClassName)

	OSPFTimersCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	OSPFTimersCount, err := strconv.Atoi(stripQuotes(OSPFTimersCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < OSPFTimersCount; i++ {
		OSPFTimersDN := stripQuotes(OSPFTimersCont.S("imdata").Index(i).S(OSPFTimersClassName, "attributes", "dn").String())
		if filterChildrenDn(OSPFTimersDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				OSPFTimersDN,
				resourceNamefromDn(OSPFTimersClassName, (OSPFTimersDN), i),
				"aci_ospf_timers",
				"aci",
				[]string{
					"bw_ref",
					"ctrl",
					"dist",
					"gr_ctrl",
					"lsa_arrival_intvl",
					"lsa_gp_pacing_intvl",
					"lsa_hold_intvl",
					"lsa_max_intvl",
					"lsa_start_intvl",
					"max_ecmp",
					"max_lsa_action",
					"max_lsa_num",
					"max_lsa_reset_intvl",
					"max_lsa_sleep_cnt",
					"max_lsa_sleep_intvl",
					"max_lsa_thresh",
					"name_alias",
					"spf_hold_intvl",
					"spf_init_intvl",
					"spf_max_intvl",
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
