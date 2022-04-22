package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudAWSProviderClass = "cloudAwsProvider"

type CloudAWSProviderGenerator struct {
	ACIService
}

func (a *CloudAWSProviderGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudAWSProviderClass)
	CloudAWSProviderCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudAWSProvidersCount, err := strconv.Atoi(stripQuotes(CloudAWSProviderCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudAWSProvidersCount; i++ {
		CloudAWSProviderDN := stripQuotes(CloudAWSProviderCont.S("imdata").Index(i).S(CloudAWSProviderClass, "attributes", "dn").String())
		if filterChildrenDn(CloudAWSProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudAWSProviderDN,
				resourceNamefromDn(CloudAWSProviderClass, (CloudAWSProviderDN), i),
				"aci_cloud_aws_provider",
				"aci",
				[]string{
					"access_key_id",
					"account_id",
					"email",
					"http_proxy",
					"is_account_in_org",
					"is_trusted",
					"name_alias",
					"provider_id",
					"region",
					"secret_access_key",
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
