package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cloudSubnetClass = "cloudSubnet"

type CloudSubnetGenerator struct {
	ACIService
}

func (a *CloudSubnetGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, cloudSubnetClass)

	cloudSubnetCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(cloudSubnetCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	cloudSubnetCount, err := strconv.Atoi(stripQuotes(cloudSubnetCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < cloudSubnetCount; i++ {
		cloudSubnetProfileDN := stripQuotes(cloudSubnetCont.S("imdata").Index(i).S(cloudSubnetClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			cloudSubnetProfileDN,
			cloudSubnetProfileDN,
			"aci_cloud_subnet",
			"aci",
			[]string{
				"name",
				"scope",
				"usage",
				"zone",
				"relation_cloud_rs_subnet_to_flow_log",
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
