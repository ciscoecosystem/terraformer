package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const anyClassName = "vzAny"

type AnyGenerator struct {
	ACIService
}

func (a *AnyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, anyClassName)

	AnyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	AnyCount, err := strconv.Atoi(stripQuotes(AnyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < AnyCount; i++ {
		AnyDN := stripQuotes(AnyCont.S("imdata").Index(i).S(anyClassName, "attributes", "dn").String())
		if filterChildrenDn(AnyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(AnyDN),
				stripQuotes(AnyDN),
				"aci_any",
				"aci",
				[]string{
					"name_alias",
					"match_t",
					"pref_gr_memb",
					"realtion_vz_rs_any_to_cons",
					"realtion_vz_rs_any_to_cons_if",
					"realtion_vz_rs_any_to_prov",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
