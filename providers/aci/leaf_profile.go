package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const leafProfileClassName = "infraNodeP"

type LeafProfileGenerator struct {
	ACIService
}

func (a *LeafProfileGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, leafProfileClassName)

	LeafProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LeafProfileCount, err := strconv.Atoi(stripQuotes(LeafProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LeafProfileCount; i++ {
		LeafProfileAttr := LeafProfileCont.S("imdata").Index(i).S(leafProfileClassName, "attributes")
		LeafProfileDN := G(LeafProfileAttr, "dn")
		if filterChildrenDn(LeafProfileDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				LeafProfileDN,
				resourceNamefromDn(leafProfileClassName, LeafProfileDN, i),
				"aci_leaf_profile",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"name_alias",
					"annotation",
					"relation_infra_rs_acc_port_p",
					"relation_infra_rs_acc_card_p",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
