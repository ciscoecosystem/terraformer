package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const BGPPeerPrefixClass = "bgpPeerPfxPol"

type BGPPeerPrefixGenerator struct {
	ACIService
}

func (a *BGPPeerPrefixGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, BGPPeerPrefixClass)

	BGPPeerPrefixCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BGPPeerPrefixCount, err := strconv.Atoi(stripQuotes(BGPPeerPrefixCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BGPPeerPrefixCount; i++ {
		BGPPeerPrefixDN := stripQuotes(BGPPeerPrefixCont.S("imdata").Index(i).S(BGPPeerPrefixClass, "attributes", "dn").String())
		if filterChildrenDn(BGPPeerPrefixDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				BGPPeerPrefixDN,
				resourceNamefromDn(BGPPeerPrefixClass, (BGPPeerPrefixDN), i),
				"aci_bgp_peer_prefix",
				"aci",
				[]string{
					"description",
					"annotation",
					"action",
					"max_pfx",
					"name_alias",
					"restart_time",
					"thresh",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
