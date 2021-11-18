package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const firmwareDownloadTaskClassName = "firmwareOSource"

type FirmwareDownloadTaskGenerator struct {
	ACIService
}

func (a *FirmwareDownloadTaskGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, firmwareDownloadTaskClassName)

	FirmwareDownloadTaskCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FirmwareDownloadTaskCount, err := strconv.Atoi(stripQuotes(FirmwareDownloadTaskCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FirmwareDownloadTaskCount; i++ {
		FirmwareDownloadTaskAttr := FirmwareDownloadTaskCont.S("imdata").Index(i).S(firmwareDownloadTaskClassName, "attributes")
		FirmwareDownloadTaskDN := stripQuotes(G(FirmwareDownloadTaskAttr, "dn"))

		if filterChildrenDn(FirmwareDownloadTaskDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				FirmwareDownloadTaskDN,
				FirmwareDownloadTaskDN,
				"aci_firmware_download_task",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"auth_pass",
					"auth_type",
					"dnld_task_flip",
					"identity_private_key_contents",
					"identity_private_key_passphrase",
					"identity_public_key_contents",
					"load_catalog_if_exists_and_newer",
					"name_alias",
					"password",
					"polling_interval",
					"proto",
					"url",
					"user",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
