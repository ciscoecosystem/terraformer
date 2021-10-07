package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpanSourceDestGroupMatchClass = "spanSpanLbl"

type SpanSourceDestGroupMatchGenerator struct {
	ACIService
}

func (a *SpanSourceDestGroupMatchGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpanSourceDestGroupMatchClass)

	SpanSourceDestGroupMatchCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	SpanSourceDestGroupMatchCount, err := strconv.Atoi(stripQuotes(SpanSourceDestGroupMatchCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < SpanSourceDestGroupMatchCount; i++ {
		SpanSourceDestGroupMatchDN := SpanSourceDestGroupMatchCont.S("imdata").Index(i).S(SpanSourceDestGroupMatchClass, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(SpanSourceDestGroupMatchDN),
			stripQuotes(SpanSourceDestGroupMatchDN),
			"aci_span_sourcedestination_group_match_label",
			"aci",
			[]string{
				"tag",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
