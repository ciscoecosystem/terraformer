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
		ServiceRedirectPolicyDN := G(ServiceRedirectPolicyAttr, "dn")
		name := G(ServiceRedirectPolicyAttr, "name")
		if filterChildrenDn(ServiceRedirectPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ServiceRedirectPolicyDN,
				resourceNamefromDn(serviceRedirectPolicyClassName, (ServiceRedirectPolicyDN), i),
				"aci_service_redirect_policy",
				"aci",
				map[string]string{
					"tenant_dn": GetParentDn(ServiceRedirectPolicyDN, fmt.Sprintf("/svcCont/svcRedirectPol-%s", name)),
				},
				[]string{
					"description",
					"anycast_enabled",
					"dest_type",
					"hashing_algorithm",
					"max_threshold_percent",
					"min_threshold_percent",
					"name_alias",
					"annotation",
					"program_local_pod_only",
					"resilient_hash_enabled",
					"threshold_down_action",
					"threshold_enable",
					"relation_vns_rs_ipsla_monitoring_pol",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
