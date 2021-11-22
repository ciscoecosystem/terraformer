package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EndpointRetentionPolicyClass = "fvEpRetPol"

type EndpointRetentionPolicyGenerator struct {
	ACIService
}

func (a *EndpointRetentionPolicyGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EndpointRetentionPolicyClass)
	EndpointRetentionPolicyCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EndpointRetentionPolicysCount, err := strconv.Atoi(stripQuotes(EndpointRetentionPolicyCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < EndpointRetentionPolicysCount; i++ {
		EndpointRetentionPolicyDN := stripQuotes(EndpointRetentionPolicyCont.S("imdata").Index(i).S(EndpointRetentionPolicyClass, "attributes", "dn").String())
		if filterChildrenDn(EndpointRetentionPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				EndpointRetentionPolicyDN,
				fmt.Sprintf("%s_%s_%d", EndpointRetentionPolicyClass, GetMOName(EndpointRetentionPolicyDN), i),
				"aci_end_point_retention_policy",
				"aci",
				[]string{
					"bounce_age_intvl",
					"bounce_trig",
					"hold_intvl",
					"local_ep_age_intvl",
					"move_freq",
					"remote_ep_age_intvl",
					"name_alias",
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
