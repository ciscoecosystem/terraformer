package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const serviceRedirectPolicyClassName = "vnsSvcRedirectPol"

type ServiceRedirectPolicyGenerator struct {
	ACIService
}

func (a *ServiceRedirectPolicyGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, serviceRedirectPolicyClassName)

	ServiceRedirectPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ServiceRedirectPolicyCount, err := strconv.Atoi(stripQuotes(ServiceRedirectPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ServiceRedirectPolicyCount; i++ {
		ServiceRedirectPolicyAttr := ServiceRedirectPolicyCont.S("imdata").Index(i).S(serviceRedirectPolicyClassName, "attributes")
		ServiceRedirectPolicyDN := G(ServiceRedirectPolicyAttr,"dn")
		name := G(ServiceRedirectPolicyAttr,"name")
		if filterChildrenDn(ServiceRedirectPolicyDN, client.parentResource) != "" {
			

			resource := terraformutils.NewResource(
					ServiceRedirectPolicyDN,
					ServiceRedirectPolicyDN,
					"aci_service_redirect_policy",
					"aci",
					map[string]string{
						"tenant_dn": GetParentDn(ServiceRedirectPolicyDN, fmt.Sprintf("/svcRedirectPol-%s", name,)),
					},
					[]string{
						"description",
					},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
		}	
	}
	return nil
}