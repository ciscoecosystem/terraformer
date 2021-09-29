package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cloudContextPClass = "cloudCtxProfile"

type CloudContextPGenerator struct {
	ACIService
}

func (a *CloudContextPGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, cloudContextPClass)

	cloudContextPCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(cloudContextPCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	cloudContextPCount, err := strconv.Atoi(stripQuotes(cloudContextPCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < cloudContextPCount; i++ {
		cloudContextPProfileDN := stripQuotes(cloudContextPCont.S("imdata").Index(i).S(cloudContextPClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			cloudContextPProfileDN,
			cloudContextPProfileDN,
			"aci_cloud_context_profile",
			"aci",
			[]string{
				"type",
				"name_alias",
				"relation_cloud_rs_ctx_to_flow_log",
				"relation_cloud_rs_to_ctx",
				"relation_cloud_rs_ctx_profile_to_region",
				"hub_network",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}