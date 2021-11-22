package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const aaaDomClass = "aaaDomain"

type AaaDomGenerator struct {
	ACIService
}

func (a *AaaDomGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, aaaDomClass)

	aaaDomCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(aaaDomCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	aaaDomCount, err := strconv.Atoi(stripQuotes(aaaDomCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < aaaDomCount; i++ {
		aaaDomProfileDN := stripQuotes(aaaDomCont.S("imdata").Index(i).S(aaaDomClass, "attributes", "dn").String())
		if filterChildrenDn(aaaDomProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				aaaDomProfileDN,
				fmt.Sprintf("%s_%s_%d", aaaDomClass, GetMOName(aaaDomProfileDN), i),
				"aci_aaa_domain",
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
