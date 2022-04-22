package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const firmwarePolicyClassName = "firmwareFwP"

type FirmwarePolicyGenerator struct {
	ACIService
}

func (a *FirmwarePolicyGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, firmwarePolicyClassName)

	FirmwarePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FirmwarePolicyCount, err := strconv.Atoi(stripQuotes(FirmwarePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FirmwarePolicyCount; i++ {
		FirmwarePolicyAttr := FirmwarePolicyCont.S("imdata").Index(i).S(firmwarePolicyClassName, "attributes")
		FirmwarePolicyDN := G(FirmwarePolicyAttr, "dn")

		if filterChildrenDn(FirmwarePolicyDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				FirmwarePolicyDN,
				resourceNamefromDn(firmwarePolicyClassName, (FirmwarePolicyDN), i),
				"aci_firmware_policy",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"effective_on_reboot",
					"ignore_compat",
					"internal_label",
					"name_alias",
					"version",
					"version_check_override",
					"annotation",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
