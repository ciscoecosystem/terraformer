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
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

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
		if filterChildrenDn(contractDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				contractDN,
				fmt.Sprintf("%s_%s_%d", contractClassName, GetMOName(contractDN), i),
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
	}
	return nil
}
