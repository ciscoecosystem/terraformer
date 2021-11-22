package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudApplicationContainerClass = "cloudApp"

type CloudApplicationContainerGenerator struct {
	ACIService
}

func (a *CloudApplicationContainerGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudApplicationContainerClass)
	CloudApplicationContainerCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudApplicationContainersCount, err := strconv.Atoi(stripQuotes(CloudApplicationContainerCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudApplicationContainersCount; i++ {
		CloudApplicationContainerDN := stripQuotes(CloudApplicationContainerCont.S("imdata").Index(i).S(CloudApplicationContainerClass, "attributes", "dn").String())
		if filterChildrenDn(CloudApplicationContainerDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudApplicationContainerDN,
				fmt.Sprintf("%s_%s_%d", CloudApplicationContainerClass, GetMOName(CloudApplicationContainerDN), i),
				"aci_cloud_applicationcontainer",
				"aci",
				[]string{
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
