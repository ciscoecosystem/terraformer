package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const contractClassName = "vzBrCP"

type ContractGenerator struct {
	ACIService
}

func (a *ContractGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, contractClassName)

	contractCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	contractCount, err := strconv.Atoi(stripQuotes(contractCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < contractCount; i++ {
		contractDN := stripQuotes(contractCont.S("imdata").Index(i).S(contractClassName, "attributes", "dn").String())

		resource := terraformutils.NewSimpleResource(
			contractDN,
			contractDN,
			"aci_contract",
			"aci",
			[]string{
				"name_alias",
				"prio",
				"scope",
				"target_dscp",
				"relation_vz_rs_graph_att",
				"filter",
				"filter_ids",
				"filter_entry_ids",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
