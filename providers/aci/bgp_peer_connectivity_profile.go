package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BGPPeerConnectivityProClass = "bgpPeerP"

type BGPPeerConnectivityProGenerator struct {
	ACIService
}

func (a *BGPPeerConnectivityProGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BGPPeerConnectivityProClass)

	BGPPeerConnectivityProCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BGPPeerConnectivityProCount, err := strconv.Atoi(stripQuotes(BGPPeerConnectivityProCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BGPPeerConnectivityProCount; i++ {
		BGPPeerConnectivityProDN := stripQuotes(BGPPeerConnectivityProCont.S("imdata").Index(i).S(BGPPeerConnectivityProClass, "attributes", "dn").String())
		if filterChildrenDn(BGPPeerConnectivityProDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				BGPPeerConnectivityProDN,
				BGPPeerConnectivityProDN,
				"aci_bgp_peer_connectivity_profile",
				"aci",
				[]string{
					"description",
					"annotation",
					"addr_t_ctrl",
					"allowed_self_as_cnt",
					"ctrl",
					"name_alias",
					"password",
					"peer_ctrl",
					"private_a_sctrl",
					"ttl",
					"weight",
					"as_number",
					"local_asn",
					"local_asn_propagate",
					"relation_bgp_rs_peer_pfx_pol",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
