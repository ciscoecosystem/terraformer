package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const sNMPContextProfileClassName = "snmpCtxP"

type SNMPContextProfileGenerator struct {
	ACIService
}

func (a *SNMPContextProfileGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, sNMPContextProfileClassName)

	SNMPContextProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SNMPContextProfileCount, err := strconv.Atoi(stripQuotes(SNMPContextProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SNMPContextProfileCount; i++ {
		SNMPContextProfileAttr := SNMPContextProfileCont.S("imdata").Index(i).S(sNMPContextProfileClassName, "attributes")
		SNMPContextProfileDN := G(SNMPContextProfileAttr, "dn")
		if filterChildrenDn(SNMPContextProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				SNMPContextProfileDN,
				SNMPContextProfileDN,
				"aci_vrf_snmp_context",
				"aci",
				map[string]string{
					"vrf_dn": GetParentDn(SNMPContextProfileDN, fmt.Sprint("/snmpctx")),
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
