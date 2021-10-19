package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const applicationEndpointSecurityGroupClass = "fvESg"

type ApplicationEndpointSecurityGroupGenerator struct {
	ACIService
}

func (a *ApplicationEndpointSecurityGroupGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, applicationEndpointSecurityGroupClass)

	endpointSecurityGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	endpointSecurityGroupCount, err := strconv.Atoi(stripQuotes(endpointSecurityGroupCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < endpointSecurityGroupCount; i++ {
		endpointSecurityGroupProfileDN := stripQuotes(endpointSecurityGroupCont.S("imdata").Index(i).S(applicationEndpointSecurityGroupClass, "attributes", "dn").String())
		if filterChildrenDn(endpointSecurityGroupProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				endpointSecurityGroupProfileDN,
				endpointSecurityGroupProfileDN,
				"aci_endpoint_security_group",
				"aci",
				[]string{
					"flood_on_encap",
					"match_t",
					"pc_enf_pref",
					"pref_gr_memb",
					"prio",
					"relation_fv_rs_cons",
					"relation_fv_rs_cons_if",
					"relation_fv_rs_cust_qos_pol",
					"relation_fv_rs_intra_epg",
					"relation_fv_rs_prot_by",
					"relation_fv_rs_prov",
					"relation_fv_rs_scope",
					"relation_fv_rs_sec_inherited",
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
