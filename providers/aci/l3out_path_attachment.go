package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const L3outPathAttachmentClass = "l3extRsPathL3OutAtt"

type L3outPathAttachmentGenerator struct {
	ACIService
}

func (a *L3outPathAttachmentGenerator) InitResources() error {

	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L3outPathAttachmentClass)
	L3outPathAttachmentCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	L3outPathAttachmentsCount, err := strconv.Atoi(stripQuotes(L3outPathAttachmentCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < L3outPathAttachmentsCount; i++ {
		L3outPathAttachmentDN := stripQuotes(L3outPathAttachmentCont.S("imdata").Index(i).S(L3outPathAttachmentClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			L3outPathAttachmentDN,
			L3outPathAttachmentDN,
			"aci_l3out_path_attachment",
			"aci",
			[]string{
				"addr",
				"autostate",
				"encap",
				"encap_scope",
				"ipv6_dad",
				"ll_addr",
				"mac",
				"mode",
				"mtu",
				"target_dscp",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}
