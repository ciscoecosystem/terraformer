package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessGroupClass = "infraRsAccBaseGrp"

type AccessGroupGenerator struct {
	ACIService
}

func (a *AccessGroupGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessGroupClass)

	accessGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(accessGroupCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	accessGroupCount, err := strconv.Atoi(stripQuotes(accessGroupCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < accessGroupCount; i++ {
		accessGroupProfileDN := stripQuotes(accessGroupCont.S("imdata").Index(i).S(accessGroupClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			accessGroupProfileDN,
			accessGroupProfileDN,
			"aci_access_group",
			"aci",
			[]string{
				"fex_id",
				"tdn",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}