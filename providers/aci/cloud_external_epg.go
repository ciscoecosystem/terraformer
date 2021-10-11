package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const CloudExternalEPGClass = "cloudExtEPg"

type CloudExternalEPGGenerator struct {
	ACIService
}

func (a *CloudExternalEPGGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, CloudExternalEPGClass)
	CloudExternalEPGCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	CloudExternalEPGsCount, err := strconv.Atoi(stripQuotes(CloudExternalEPGCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < CloudExternalEPGsCount; i++ {
		CloudExternalEPGDN := stripQuotes(CloudExternalEPGCont.S("imdata").Index(i).S(CloudExternalEPGClass, "attributes", "dn").String())
		if filterChildrenDn(CloudExternalEPGDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				CloudExternalEPGDN,
				resourceNamefromDn(CloudExternalEPGClass, (CloudExternalEPGDN), i),
				"aci_cloud_external_epg",
				"aci",
				[]string{
					"match_expression",
					"exception_tag",
					"flood_on_encap",
					"match_t",
					"name_alias",
					"pref_gr_memb",
					"prio",
					"route_reachability",
					"relation_fv_rs_sec_inherited",
					"relation_fv_rs_prov",
					"relation_fv_rs_cons_if",
					"relation_fv_rs_cust_qos_pol",
					"relation_fv_rs_cons",
					"relation_cloud_rs_cloud_epg_ctx",
					"relation_fv_rs_prot_by",
					"relation_fv_rs_intra_epg",
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
