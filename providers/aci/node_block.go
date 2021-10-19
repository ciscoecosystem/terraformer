package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const nodeBlockClass = "infraNodeBlk"

type NodeBlockGenerator struct {
	ACIService
}

func (a *NodeBlockGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, nodeBlockClass)

	nodeBlockCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(nodeBlockCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	nodeBlockCount, err := strconv.Atoi(stripQuotes(nodeBlockCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < nodeBlockCount; i++ {
		nodeBlockProfileDN := stripQuotes(nodeBlockCont.S("imdata").Index(i).S(nodeBlockClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			nodeBlockProfileDN,
			nodeBlockProfileDN,
			"aci_node_block",
			"aci",
			[]string{
				"name_alias",
				"from_",
				"to_",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}