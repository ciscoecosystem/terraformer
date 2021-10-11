package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const leafSelectorClass = "infraLeafS"

type LeafSelectorGenerator struct {
	ACIService
}

func (a *LeafSelectorGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, leafSelectorClass)

	leafSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(leafSelectorCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	leafSelectorCount, err := strconv.Atoi(stripQuotes(leafSelectorCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < leafSelectorCount; i++ {
		leafSelectorProfileDN := stripQuotes(leafSelectorCont.S("imdata").Index(i).S(leafSelectorClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			leafSelectorProfileDN,
			leafSelectorProfileDN,
			"aci_leaf_selector",
			"aci",
			[]string{
				"name_alias",
				"relation_infra_rs_acc_node_p_grp",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}