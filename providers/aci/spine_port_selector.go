package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const SpinePortSelectorClassName = "infraRsSpAccPortP"

type SpinePortSelectorGenerator struct {
	ACIService
}

func (a *SpinePortSelectorGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, SpinePortSelectorClassName)

	SpinePortSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SpinePortSelectorCount, err := strconv.Atoi(stripQuotes(SpinePortSelectorCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SpinePortSelectorCount; i++ {
		SpinePortSelectorDN := stripQuotes(SpinePortSelectorCont.S("imdata").Index(i).S(SpinePortSelectorClassName, "attributes", "dn").String())
		SpinePortSelectorAttr := SpinePortSelectorCont.S("imdata").Index(i).S(SpinePortSelectorClassName, "attributes")
		SpinePortSelectortDn := G(SpinePortSelectorAttr, "tDn")
		if filterChildrenDn(SpinePortSelectorDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				SpinePortSelectorDN,
				resourceNamefromDn(SpinePortSelectorClassName, (SpinePortSelectorDN), i),
				"aci_spine_port_selector",
				"aci",
				map[string]string{
					"spine_profile_dn": GetParentDn(SpinePortSelectorDN, fmt.Sprintf("/rsspAccPortP-[%s]", SpinePortSelectortDn)),
				},
				[]string{
					"annotation",
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
