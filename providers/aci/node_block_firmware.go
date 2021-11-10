package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const NodeBlockFirmWareClassName = "fabricNodeBlk"

type NodeBlockFirmWareGenerator struct {
	ACIService
}

func (a *NodeBlockFirmWareGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, NodeBlockFirmWareClassName)

	NodeBlockFirmWaresCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	NodeBlockFirmWareCount, err := strconv.Atoi(stripQuotes(NodeBlockFirmWaresCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < NodeBlockFirmWareCount; i++ {
		NodeBlockFirmWareDN := NodeBlockFirmWaresCont.S("imdata").Index(i).S(NodeBlockFirmWareClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(NodeBlockFirmWareDN),
			stripQuotes(NodeBlockFirmWareDN),
			"aci_node_block_firmware",
			"aci",
			[]string{
				"from_",
				"to_",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}