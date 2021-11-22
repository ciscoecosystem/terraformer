package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudEndpointSelectorClass = "cloudEPSelector"

type CloudEndpointSelectorGenerator struct {
	ACIService
}

func (a *CloudEndpointSelectorGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudEndpointSelectorClass)
	CloudEndpointSelectorCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudEndpointSelectorsCount, err := strconv.Atoi(stripQuotes(CloudEndpointSelectorCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudEndpointSelectorsCount; i++ {
		CloudEndpointSelectorDN := stripQuotes(CloudEndpointSelectorCont.S("imdata").Index(i).S(CloudEndpointSelectorClass, "attributes", "dn").String())
		if filterChildrenDn(CloudEndpointSelectorDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudEndpointSelectorDN,
				fmt.Sprintf("%s_%s_%d", CloudEndpointSelectorDN, GetMOName(CloudEndpointSelectorDN), i),
				"aci_cloud_endpoint_selector",
				"aci",
				[]string{
					"match_expression",
					"name_alias",
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
