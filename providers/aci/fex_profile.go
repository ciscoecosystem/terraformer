package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fexProfClass = "infraFexP"

type FexProfGenerator struct {
	ACIService
}

func (a *FexProfGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fexProfClass)

	fexProfCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(fexProfCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	fexProfCount, err := strconv.Atoi(stripQuotes(fexProfCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < fexProfCount; i++ {
		fexProfProfileDN := stripQuotes(fexProfCont.S("imdata").Index(i).S(fexProfClass, "attributes", "dn").String())
		if filterChildrenDn(fexProfProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				fexProfProfileDN,
				resourceNamefromDn(fexProfClass, (fexProfProfileDN), i),
				"aci_fex_profile",
				"aci",
				[]string{
					"name",
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
