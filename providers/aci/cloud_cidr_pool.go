package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudCidrPoolClass = "cloudCidr"

type CloudCidrPoolGenerator struct {
	ACIService
}

func (a *CloudCidrPoolGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudCidrPoolClass)
	CloudCidrPoolCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudCidrPoolsCount, err := strconv.Atoi(stripQuotes(CloudCidrPoolCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudCidrPoolsCount; i++ {
		CloudCidrPoolDN := stripQuotes(CloudCidrPoolCont.S("imdata").Index(i).S(CloudCidrPoolClass, "attributes", "dn").String())
		if filterChildrenDn(CloudCidrPoolDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudCidrPoolDN,
				resourceNamefromDn(CloudCidrPoolClass, (CloudCidrPoolDN), i),
				"aci_cloud_cidr_pool",
				"aci",
				[]string{
					"name_alias",
					"primary",
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
