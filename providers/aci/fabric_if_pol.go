package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const FabricIfPolClass = "fabricHIfPol"

type FabricIfPolGenerator struct {
	ACIService
}

func (a *FabricIfPolGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, FabricIfPolClass)

	FabricIfPolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FabricIfPolCount, err := strconv.Atoi(stripQuotes(FabricIfPolCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < FabricIfPolCount; i++ {
		FabricIfPolDN := stripQuotes(FabricIfPolCont.S("imdata").Index(i).S(FabricIfPolClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			FabricIfPolDN,
			FabricIfPolDN,
			"aci_fabric_if_pol",
			"aci",
			[]string{
				"auto_neg",
				"fec_mode",
				"link_debounce",
				"name_alias",
				"speed",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
