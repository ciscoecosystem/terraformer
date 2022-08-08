package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const iSISDomainPolicyClassName = "isisDomPol"

type ISISDomainPolicyGenerator struct {
	ACIService
}

func (a *ISISDomainPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, iSISDomainPolicyClassName)

	ISISDomainPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ISISDomainPolicyCount, err := strconv.Atoi(stripQuotes(ISISDomainPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ISISDomainPolicyCount; i++ {
		ISISDomainPolicyAttr := ISISDomainPolicyCont.S("imdata").Index(i).S(iSISDomainPolicyClassName, "attributes")
		ISISDomainPolicyDN := G(ISISDomainPolicyAttr, "dn")
		if filterChildrenDn(ISISDomainPolicyDN, client.parentResource) != "" && ISISDomainPolicyDN == "uni/fabric/isisDomP-default" {
			resource := terraformutils.NewResource(
				ISISDomainPolicyDN,
				resourceNamefromDn(iSISDomainPolicyClassName, ISISDomainPolicyDN, i),
				"aci_isis_domain_policy",
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
