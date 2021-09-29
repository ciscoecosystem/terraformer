package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cloudDomainPClass = "cloudDomP"

type CloudDomainPGenerator struct {
	ACIService
}

func (a *CloudDomainPGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, cloudDomainPClass)

	cloudDomainPCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(cloudDomainPCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	cloudDomainPCount, err := strconv.Atoi(stripQuotes(cloudDomainPCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < cloudDomainPCount; i++ {
		cloudDomainPProfileDN := stripQuotes(cloudDomainPCont.S("imdata").Index(i).S(cloudDomainPClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			cloudDomainPProfileDN,
			cloudDomainPProfileDN,
			"aci_cloud_domain_profile",
			"aci",
			[]string{
				"site_id",
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