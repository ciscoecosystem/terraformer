package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const applicationProfileClass = "fvAp"

type ApplicationProfileGenerator struct {
	ACIService
}

func (a *ApplicationProfileGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, applicationProfileClass)

	apCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	apCount, err := strconv.Atoi(stripQuotes(apCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < apCount; i++ {
		apDN := stripQuotes(apCont.S("imdata").Index(i).S(applicationProfileClass, "attributes", "dn").String())
		if filterChildrenDn(apDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				apDN,
				resourceNamefromDn(applicationProfileClass, apDN, i),
				"aci_application_profile",
				"aci",
				[]string{
					"name_alias",
					"prio",
					"relation_fv_rs_ap_mon_pol",
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
