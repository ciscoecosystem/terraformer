package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const iPAgingPolicyClassName = "epIpAgingP"

type IPAgingPolicyGenerator struct {
	ACIService
}

func (a *IPAgingPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, iPAgingPolicyClassName)

	IPAgingPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	IPAgingPolicyCount, err := strconv.Atoi(stripQuotes(IPAgingPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < IPAgingPolicyCount; i++ {
		IPAgingPolicyAttr := IPAgingPolicyCont.S("imdata").Index(i).S(iPAgingPolicyClassName, "attributes")
		IPAgingPolicyDN := G(IPAgingPolicyAttr, "dn")
		if filterChildrenDn(IPAgingPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				IPAgingPolicyDN,
				resourceNamefromDn(iPAgingPolicyClassName, IPAgingPolicyDN, i),
				"aci_endpoint_ip_aging_profile",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
