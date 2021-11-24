package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const filterClassName = "vzFilter"

type FilterGenerator struct {
	ACIService
}

func (a *FilterGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, filterClassName)

	FiltersCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FilterCount, err := strconv.Atoi(stripQuotes(FiltersCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FilterCount; i++ {
		FilterDN := stripQuotes(FiltersCont.S("imdata").Index(i).S(filterClassName, "attributes", "dn").String())
		if filterChildrenDn(FilterDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				FilterDN,
				resourceNamefromDn(filterClassName, (FilterDN), i),
				"aci_filter",
				"aci",
				[]string{
					"name_alias",
					"relation_vz_rs_filt_graph_att",
					"relation_vz_rs_fwd_r_flt_p_att",
					"relation_vz_rs_rev_r_flt_p_att",
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
