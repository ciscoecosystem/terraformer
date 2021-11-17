package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const firmwareGroupClassName = "firmwareFwGrp"

type FirmwareGroupGenerator struct {
	ACIService
}

func (a *FirmwareGroupGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, firmwareGroupClassName)

	FirmwareGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FirmwareGroupCount, err := strconv.Atoi(stripQuotes(FirmwareGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FirmwareGroupCount; i++ {
		FirmwareGroupAttr := FirmwareGroupCont.S("imdata").Index(i).S(firmwareGroupClassName, "attributes")
		FirmwareGroupDN := stripQuotes(G(FirmwareGroupAttr, "dn"))
		if filterChildrenDn(FirmwareGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				FirmwareGroupDN,
				FirmwareGroupDN,
				"aci_firmware_group",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"name_alias",
					"firmware_group_type",
					"relation_firmware_rs_fwgrpp",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
