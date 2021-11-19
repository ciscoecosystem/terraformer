package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpanSourceGroupClass = "spanSrcGrp"

type SpanSourceGroupGenerator struct {
	ACIService
}

func (a *SpanSourceGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpanSourceGroupClass)

	SpanSourceGroupCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	SpanSourceGroupCount, err := strconv.Atoi(stripQuotes(SpanSourceGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < SpanSourceGroupCount; i++ {
		SpanSourceGroupDN := stripQuotes(SpanSourceGroupCont.S("imdata").Index(i).S(SpanSourceGroupClass, "attributes", "dn").String())
		if filterChildrenDn(SpanSourceGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(SpanSourceGroupDN),
				stripQuotes(SpanSourceGroupDN),
				"aci_span_source_group",
				"aci",
				[]string{
					"admin_st",
					"relation_span_rs_src_grp_to_filter_grp",
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
