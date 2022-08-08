package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const sNMPCommunityClassName = "snmpCommunityP"

type SNMPCommunityGenerator struct {
	ACIService
}

func (a *SNMPCommunityGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, sNMPCommunityClassName)

	SNMPCommunityCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	SNMPCommunityCount, err := strconv.Atoi(stripQuotes(SNMPCommunityCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < SNMPCommunityCount; i++ {
		SNMPCommunityAttr := SNMPCommunityCont.S("imdata").Index(i).S(sNMPCommunityClassName, "attributes")
		SNMPCommunityDN := G(SNMPCommunityAttr, "dn")
		name := G(SNMPCommunityAttr, "name")
		if filterChildrenDn(SNMPCommunityDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				SNMPCommunityDN,
				resourceNamefromDn(sNMPCommunityClassName, SNMPCommunityDN, i),
				"aci_vrf_snmp_context_community",
				"aci",
				map[string]string{
					"vrf_snmp_context_dn": GetParentDn(SNMPCommunityDN, fmt.Sprintf("/community-%s", name)),
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
