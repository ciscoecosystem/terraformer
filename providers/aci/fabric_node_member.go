package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fabricNodeMemberClass = "fabricNodeIdentP"

type FabricNodeMemberGenerator struct {
	ACIService
}

func (a *FabricNodeMemberGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fabricNodeMemberClass)

	fabricNodeMemberCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(fabricNodeMemberCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	fabricNodeMemberCount, err := strconv.Atoi(stripQuotes(fabricNodeMemberCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < fabricNodeMemberCount; i++ {
		fabricNodeMemberProfileDN := stripQuotes(fabricNodeMemberCont.S("imdata").Index(i).S(fabricNodeMemberClass, "attributes", "dn").String())
		if filterChildrenDn(fabricNodeMemberProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				fabricNodeMemberProfileDN,
				resourceNamefromDn(fabricNodeMemberClass, (fabricNodeMemberProfileDN), i),
				"aci_fabric_node_member",
				"aci",
				[]string{
					"name",
					"ext_pool_id",
					"fabric_id",
					"name_alias",
					"node_id",
					"node_type",
					"pod_id",
					"role",
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
