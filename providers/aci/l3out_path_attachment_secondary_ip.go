package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3outPathAttachmentSecondaryIPClass = "l3extIp"

type L3outPathAttachmentSecondaryIPGenerator struct {
	ACIService
}

func (a *L3outPathAttachmentSecondaryIPGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3outPathAttachmentSecondaryIPClass)
	L3outPathAttachmentSecondaryIPCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3outPathAttachmentSecondaryIPsCount, err := strconv.Atoi(stripQuotes(L3outPathAttachmentSecondaryIPCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3outPathAttachmentSecondaryIPsCount; i++ {
		L3outPathAttachmentSecondaryIPDN := stripQuotes(L3outPathAttachmentSecondaryIPCont.S("imdata").Index(i).S(L3outPathAttachmentSecondaryIPClass, "attributes", "dn").String())
		if filterChildrenDn(L3outPathAttachmentSecondaryIPDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				L3outPathAttachmentSecondaryIPDN,
				resourceNamefromDn(L3outPathAttachmentSecondaryIPClass, (L3outPathAttachmentSecondaryIPDN), i),
				"aci_l3out_path_attachment_secondary_ip",
				"aci",
				[]string{
					"name_alias",
					"ipv6_dad",
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
