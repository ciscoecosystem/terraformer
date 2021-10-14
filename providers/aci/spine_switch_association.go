package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const spineSwitchAssClass = "infraSpineS"

type SpineSwitchAssGenerator struct {
	ACIService
}

func (a *SpineSwitchAssGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, spineSwitchAssClass)

	spineSwitchAssCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(spineSwitchAssCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	spineSwitchAssCount, err := strconv.Atoi(stripQuotes(spineSwitchAssCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < spineSwitchAssCount; i++ {
		spineSwitchAssProfileDN := stripQuotes(spineSwitchAssCont.S("imdata").Index(i).S(spineSwitchAssClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			spineSwitchAssProfileDN,
			spineSwitchAssProfileDN,
			"aci_spine_switch_association",
			"aci",
			[]string{
				"name_alias",
				"relation_infra_rs_spine_acc_node_p_grp",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}