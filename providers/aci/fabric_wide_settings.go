package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fabricWideSettingsPolicyClassName = "infraSetPol"

type fabricWideSettingsPolicyGenerator struct {
	ACIService
}

func (a *fabricWideSettingsPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fabricWideSettingsPolicyClassName)

	fabricWideSettingsPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	fabricWideSettingsPolicyCount, err := strconv.Atoi(stripQuotes(fabricWideSettingsPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < fabricWideSettingsPolicyCount; i++ {
		fabricWideSettingsPolicyAttr := fabricWideSettingsPolicyCont.S("imdata").Index(i).S(fabricWideSettingsPolicyClassName, "attributes")
		fabricWideSettingsPolicyDN := G(fabricWideSettingsPolicyAttr, "dn")
		if filterChildrenDn(fabricWideSettingsPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				fabricWideSettingsPolicyDN,
				resourceNamefromDn(fabricWideSettingsPolicyClassName, fabricWideSettingsPolicyDN, i),
				"aci_fabric_wide_settings",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
