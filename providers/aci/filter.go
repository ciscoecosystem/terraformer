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
	client, err := a.createClient()
	if err != nil {
		return err
	}

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
		FilterDN := FiltersCont.S("imdata").Index(i).S(filterClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(FilterDN),
			stripQuotes(FilterDN),
			"aci_filter",
			"aci",
			[]string{
				"name_alias",
				"relation_vz_rs_filt_graph_att",
				"relation_vz_rs_fwd_r_flt_p_att",
				"relation_vz_rs_rev_r_flt_p_att",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}