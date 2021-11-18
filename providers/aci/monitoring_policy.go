package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const monPolClass = "monEPGPol"

type MonPolGenerator struct {
	ACIService
}

func (a *MonPolGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, monPolClass)

	monPolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(monPolCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	monPolCount, err := strconv.Atoi(stripQuotes(monPolCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < monPolCount; i++ {
		monPolProfileDN := stripQuotes(monPolCont.S("imdata").Index(i).S(monPolClass, "attributes", "dn").String())
		if filterChildrenDn(monPolProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				monPolProfileDN,
				monPolProfileDN,
				"aci_monitoring_policy",
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
