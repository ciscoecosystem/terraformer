package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudEndpointSelectorForExternalEpgsClass = "cloudExtEPSelector"

type CloudEndpointSelectorForExternalEpgsGenerator struct {
	ACIService
}

func (a *CloudEndpointSelectorForExternalEpgsGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudEndpointSelectorForExternalEpgsClass)
	CloudEndpointSelectorForExternalEpgsCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudEndpointSelectorForExternalEpgssCount, err := strconv.Atoi(stripQuotes(CloudEndpointSelectorForExternalEpgsCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudEndpointSelectorForExternalEpgssCount; i++ {
		CloudEndpointSelectorForExternalEpgsDN := stripQuotes(CloudEndpointSelectorForExternalEpgsCont.S("imdata").Index(i).S(CloudEndpointSelectorForExternalEpgsClass, "attributes", "dn").String())
		if filterChildrenDn(CloudEndpointSelectorForExternalEpgsDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudEndpointSelectorForExternalEpgsDN,
				CloudEndpointSelectorForExternalEpgsDN,
				"aci_cloud_endpoint_selectorfor_external_epgs",
				"aci",
				[]string{
					"is_shared",
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
