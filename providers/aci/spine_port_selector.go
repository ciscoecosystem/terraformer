package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpinePortSelectorClassName = "infraRsSpAccPortP"

type SpinePortSelectorGenerator struct {
	ACIService
}

func (a *SpinePortSelectorGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpinePortSelectorClassName)

	SpinePortSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SpinePortSelectorCount, err := strconv.Atoi(stripQuotes(SpinePortSelectorCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SpinePortSelectorCount; i++ {
		SpinePortSelectorDN := SpinePortSelectorCont.S("imdata").Index(i).S(SpinePortSelectorClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(SpinePortSelectorDN),
			stripQuotes(SpinePortSelectorDN),
			"aci_spine_port_selector",
			"aci",
			[]string{
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
