package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const spanningTreeInterfacePolicyClass = "stpIfPol"

type SpanningTreeInterfacePolicyGenerator struct {
	ACIService
}

func (a *SpanningTreeInterfacePolicyGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, spanningTreeInterfacePolicyClass)

	spanningTreeInterfacePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(spanningTreeInterfacePolicyCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	spanningTreeInterfacePolicyCount, err := strconv.Atoi(stripQuotes(spanningTreeInterfacePolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < spanningTreeInterfacePolicyCount; i++ {
		spanningTreeInterfacePolicyProfileDN := stripQuotes(spanningTreeInterfacePolicyCont.S("imdata").Index(i).S(spanningTreeInterfacePolicyClass, "attributes", "dn").String())
		if filterChildrenDn(spanningTreeInterfacePolicyProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				spanningTreeInterfacePolicyProfileDN,
				resourceNamefromDn(spanningTreeInterfacePolicyClass, (spanningTreeInterfacePolicyProfileDN), i),
				"aci_spanning_tree_interface_policy",
				"aci",
				[]string{
					"ctrl",
					"annotation",
					"description",
					"name_alias",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
