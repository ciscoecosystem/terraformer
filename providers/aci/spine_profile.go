package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const spinePClass = "infraSpineP"

type SpinePGenerator struct {
	ACIService
}

func (a *SpinePGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, spinePClass)

	spinePCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(spinePCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	spinePCount, err := strconv.Atoi(stripQuotes(spinePCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < spinePCount; i++ {
		spinePProfileDN := stripQuotes(spinePCont.S("imdata").Index(i).S(spinePClass, "attributes", "dn").String())
		if filterChildrenDn(spinePProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				spinePProfileDN,
				resourceNamefromDn(spinePClass, (spinePProfileDN), i),
				"aci_spine_profile",
				"aci",
				[]string{
					"name_alias",
					"relation_infra_rs_sp_acc_port_p",
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
