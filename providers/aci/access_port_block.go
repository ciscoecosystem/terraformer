package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessPortBlkClass = "infraPortBlk"

type AccessPortBlkGenerator struct {
	ACIService
}

func (a *AccessPortBlkGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessPortBlkClass)

	accessPortBlkCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(accessPortBlkCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	accessPortBlkCount, err := strconv.Atoi(stripQuotes(accessPortBlkCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < accessPortBlkCount; i++ {
		accessPortBlkProfileDN := stripQuotes(accessPortBlkCont.S("imdata").Index(i).S(accessPortBlkClass, "attributes", "dn").String())
		if filterChildrenDn(accessPortBlkProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				accessPortBlkProfileDN,
				fmt.Sprintf("%s_%s_%d", accessPortBlkClass, GetMOName(accessPortBlkProfileDN), i),
				"aci_access_port_block",
				"aci",
				[]string{
					"name_alias",
					"from_card",
					"from_port",
					"to_card",
					"to_port",
					"relation_infra_rs_acc_bndl_subgrp",
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
