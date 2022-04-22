package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EPGToDomainClass = "fvRsDomAtt"

type EPGToDomainGenerator struct {
	ACIService
}

func (a *EPGToDomainGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EPGToDomainClass)

	EPGToDomainCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGToDomainCount, err := strconv.Atoi(stripQuotes(EPGToDomainCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < EPGToDomainCount; i++ {
		EPGToDomainDN := stripQuotes(EPGToDomainCont.S("imdata").Index(i).S(EPGToDomainClass, "attributes", "dn").String())
		if filterChildrenDn(EPGToDomainDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				EPGToDomainDN,
				resourceNamefromDn(EPGToDomainClass, EPGToDomainDN, i),
				"aci_epg_to_domain",
				"aci",
				[]string{
					"annotation",
					"binding_type",
					"allow_micro_seg",
					"delimiter",
					"encap",
					"encap_mode",
					"epg_cos",
					"epg_cos_pref",
					"instr_imedcy",
					"lag_policy_name",
					"netflow_dir",
					"netflow_pref",
					"num_ports",
					"port_allocation",
					"primary_encap",
					"primary_encap_inner",
					"res_imedcy",
					"secondary_encap_inner",
					"switching_mode",
					"vmm_allow_promiscuous",
					"vmm_forged_transmits",
					"vmm_mac_changes",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
