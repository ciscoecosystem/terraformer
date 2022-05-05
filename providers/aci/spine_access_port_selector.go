package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const spineAccessPortSelectorClassName = "infraSHPortS"

type SpineAccessPortSelectorGenerator struct {
	ACIService
}

func (a *SpineAccessPortSelectorGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, spineAccessPortSelectorClassName)

	SpineAccessPortSelectorCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SpineAccessPortSelectorCount, err := strconv.Atoi(stripQuotes(SpineAccessPortSelectorCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SpineAccessPortSelectorCount; i++ {
		SpineAccessPortSelectorAttr := SpineAccessPortSelectorCont.S("imdata").Index(i).S(spineAccessPortSelectorClassName, "attributes")
		SpineAccessPortSelectorDN := G(SpineAccessPortSelectorAttr, "dn")
		name := G(SpineAccessPortSelectorAttr, "name")
		spine_access_port_selector_type := G(SpineAccessPortSelectorAttr, "spine_access_port_selector_type")
		if filterChildrenDn(SpineAccessPortSelectorDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				SpineAccessPortSelectorDN,
				resourceNamefromDn(spineAccessPortSelectorClassName, SpineAccessPortSelectorDN, i),
				"aci_spine_access_port_selector",
				"aci",
				map[string]string{
					"spine_interface_profile_dn": GetParentDn(SpineAccessPortSelectorDN, fmt.Sprintf("/shports-%s-typ-%s", name, spine_access_port_selector_type)),
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
