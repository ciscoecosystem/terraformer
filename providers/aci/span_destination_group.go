package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpanDestinationGroupClass = "spanDestGrp"

type SpanDestinationGroupGenerator struct {
	ACIService
}

func (a *SpanDestinationGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpanDestinationGroupClass)

	SpanDestinationGroupCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	SpanDestinationGroupCount, err := strconv.Atoi(stripQuotes(SpanDestinationGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < SpanDestinationGroupCount; i++ {
		SpanDestinationGroupDN := stripQuotes(SpanDestinationGroupCont.S("imdata").Index(i).S(SpanDestinationGroupClass, "attributes", "dn").String())
		if filterChildrenDn(SpanDestinationGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(SpanDestinationGroupDN),
				stripQuotes(SpanDestinationGroupDN),
				"aci_span_destination_group",
				"aci",
				[]string{
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
