package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BDDHCPLabelClass = "dhcpLbl"

type BDDHCPLabelGenerator struct {
	ACIService
}

func (a *BDDHCPLabelGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BDDHCPLabelClass)

	BDDHCPLblCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BDDHCPLblCount, err := strconv.Atoi(stripQuotes(BDDHCPLblCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BDDHCPLblCount; i++ {
		BDDHCPLblDN := stripQuotes(BDDHCPLblCont.S("imdata").Index(i).S(BDDHCPLabelClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			BDDHCPLblDN,
			BDDHCPLblDN,
			"aci_bd_dhcp_label",
			"aci",
			[]string{
				"name_alias",
				"owner",
				"tag",
				"relation_dhcp_rs_dhcp_option_pol",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
