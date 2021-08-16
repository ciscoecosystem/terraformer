package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const subContractClassName = "vzSubj"

type ContractSubjectGenerator struct {
	ACIService
}

func (a *ContractSubjectGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, subContractClassName)

	subContractCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	subContractCount, err := strconv.Atoi(stripQuotes(subContractCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < subContractCount; i++ {
		contractDN := stripQuotes(subContractCont.S("imdata").Index(i).S(subContractClassName, "attributes", "dn").String())

		resource := terraformutils.NewSimpleResource(
			contractDN,
			contractDN,
			"aci_contract_subject",
			"aci",
			[]string{
				"cons_match_t",
				"name_alias",
				"prio",
				"prov_match_t",
				"rev_flt_ports",
				"target_dscp",
				"relation_vz_rs_subj_graph_att",
				"relation_vz_rs_sdwan_pol",
				"relation_vz_rs_subj_filt_att",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}