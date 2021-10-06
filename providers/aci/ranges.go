package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const rangesClass = "fvnsEncapBlk"

type RangesGenerator struct {
	ACIService
}

func (a *RangesGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, rangesClass)

	rangesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(rangesCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	rangesCount, err := strconv.Atoi(stripQuotes(rangesCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < rangesCount; i++ {
		rangesProfileDN := stripQuotes(rangesCont.S("imdata").Index(i).S(rangesClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			rangesProfileDN,
			rangesProfileDN,
			"aci_ranges",
			"aci",
			[]string{
				"alloc_mode",
				"role",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}