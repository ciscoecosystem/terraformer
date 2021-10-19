package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const bridgeDomainClassName = "fvBD"

type BridgeDomainGenerator struct {
	ACIService
}

func (a *BridgeDomainGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, bridgeDomainClassName)

	BridgeDomainsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	BridgeDomainCount, err := strconv.Atoi(stripQuotes(BridgeDomainsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < BridgeDomainCount; i++ {
		BridgeDomainDN := BridgeDomainsCont.S("imdata").Index(i).S(bridgeDomainClassName, "attributes", "dn").String()
		if filterChildrenDn(BridgeDomainDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(BridgeDomainDN),
				stripQuotes(BridgeDomainDN),
				"aci_bridge_domain",
				"aci",
				[]string{
					"optimize_band_width",
					"arp_flood",
					"ep_clear",
					"host_based_routing",
					"intersite_bum_traffic_allow",
					"intersite_l2_stretch",
					"ip_learning",
					"ipv6_mcast_allow",
					"limit_ip_learn_to_subnets",
					"ll_addr",
					"mac",
					"mcast_allow",
					"multi_dst_pkt_act",
					"name_alias",
					"bridge_domain_type",
					"unicast_route",
					"unk_mac_ucast_act",
					"unk_mcast_act",
					"v6unk_mcast_act",
					"vmac",
					"relation_fv_rs_bd_to_profile",
					"relation_fv_rs_mldsn",
					"relation_fv_rs_abd_pol_mon_pol",
					"relation_fv_rs_bd_to_nd_p",
					"relation_fv_rs_bd_flood_to",
					"relation_fv_rs_bd_to_fhs",
					"relation_fv_rs_bd_to_relay_p",
					"relation_fv_rs_ctx",
					"relation_fv_rs_bd_to_netflow_monitor_pol",
					"relation_fv_rs_igmpsn",
					"relation_fv_rs_bd_to_ep_ret",
					"relation_fv_rs_bd_to_out",
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
