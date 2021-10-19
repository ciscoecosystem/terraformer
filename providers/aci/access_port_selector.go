package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessPortSelectorClass = "infraHPortS"

type AccessPortSelectorGenerator struct {
	ACIService
}

func (a *AccessPortSelectorGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessPortSelectorClass)

	accessPortSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(accessPortSelectorCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	accessPortSelectorCount, err := strconv.Atoi(stripQuotes(accessPortSelectorCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < accessPortSelectorCount; i++ {
		accessPortSelectorProfileDN := stripQuotes(accessPortSelectorCont.S("imdata").Index(i).S(accessPortSelectorClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			accessPortSelectorProfileDN,
			accessPortSelectorProfileDN,
			"aci_access_port_selector",
			"aci",
			[]string{
				"name_alias",
				"relation_infra_rs_acc_base_grp",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}