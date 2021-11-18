package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tabooContractClass = "vzTaboo"

type TabooContractGenerator struct {
	ACIService
}

func (a *TabooContractGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tabooContractClass)

	tabooContractCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(tabooContractCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	tabooContractCount, err := strconv.Atoi(stripQuotes(tabooContractCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < tabooContractCount; i++ {
		tabooContractProfileDN := stripQuotes(tabooContractCont.S("imdata").Index(i).S(tabooContractClass, "attributes", "dn").String())
		if filterChildrenDn(tabooContractProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				tabooContractProfileDN,
				tabooContractProfileDN,
				"aci_taboo_contract",
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
	}

	return nil
}
