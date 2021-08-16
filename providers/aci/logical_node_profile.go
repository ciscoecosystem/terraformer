package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const logicalNodeProfileClassName = "l3extLNodeP"

type LogicalNodeProfileGenerator struct {
	ACIService
}

func (a *LogicalNodeProfileGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, logicalNodeProfileClassName)

	LogicalNodeProfilesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LogicalNodeProfileCount, err := strconv.Atoi(stripQuotes(LogicalNodeProfilesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LogicalNodeProfileCount; i++ {
		LogicalNodeProfileDN := LogicalNodeProfilesCont.S("imdata").Index(i).S(logicalNodeProfileClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(LogicalNodeProfileDN),
			stripQuotes(LogicalNodeProfileDN),
			"aci_logical_node_profile",
			"aci",
			[]string{
				"name_alias",
				"config_issues",
				"tag",
				"target_dscp",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
