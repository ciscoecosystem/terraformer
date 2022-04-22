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
		cloudSubnetAttr := cloudSubnetCont.S("imdata").Index(i).S(cloudSubnetClass, "attributes")
		cloudSubnetProfileDN := stripQuotes(cloudSubnetCont.S("imdata").Index(i).S(cloudSubnetClass, "attributes", "dn").String())
		ip := G(cloudSubnetAttr, "ip")
		if filterChildrenDn(cloudSubnetProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				cloudSubnetProfileDN,
				resourceNamefromDn(cloudSubnetClass, (cloudSubnetProfileDN), i),
				"aci_cloud_subnet",
				"aci",
				map[string]string{
					"cloud_cidr_pool_dn": GetParentDn(cloudSubnetProfileDN, fmt.Sprintf("/subnet-[%s]", ip)),
				},
				[]string{
					"description",
					"relation_cloud_rs_subnet_to_flow_log",
					"annotation",
					"name_alias",
					"scope",
					"usage",
					"zone",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
