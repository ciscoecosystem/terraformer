package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessSubPortBlkClass = "infraSubPortBlk"

type AccessSubPortBlkGenerator struct {
	ACIService
}

func (a *AccessSubPortBlkGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessSubPortBlkClass)

	accessSubPortBlkCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(accessSubPortBlkCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	accessSubPortBlkCount, err := strconv.Atoi(stripQuotes(accessSubPortBlkCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < accessSubPortBlkCount; i++ {
		accessSubPortBlkProfileDN := stripQuotes(accessSubPortBlkCont.S("imdata").Index(i).S(accessSubPortBlkClass, "attributes", "dn").String())
		if filterChildrenDn(accessSubPortBlkProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				accessSubPortBlkProfileDN,
				resourceNamefromDn(accessSubPortBlkClass,accessSubPortBlkProfileDN,i),
				"aci_access_sub_port_block",
				"aci",
				[]string{
					"name_alias",
					"from_card",
					"from_port",
					"from_sub_port",
					"to_card",
					"to_port",
					"to_sub_port",
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
