package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const filterEntryClassName = "vzEntry"

type FilterEntryGenerator struct {
	ACIService
}

func (a *FilterEntryGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, filterEntryClassName)

	FilterEntriesCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FilterEntryCount, err := strconv.Atoi(stripQuotes(FilterEntriesCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FilterEntryCount; i++ {
		FilterEntryDN := FilterEntriesCont.S("imdata").Index(i).S(filterEntryClassName, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(FilterEntryDN),
			stripQuotes(FilterEntryDN),
			"aci_filter_entry",
			"aci",
			[]string{
				"name_alias",
				"apply_to_frag",
				"arp_opc",
				"d_from_port",
				"d_to_port",
				"ether_t",
				"icmpv4_t",
				"icmpv6_t",
				"match_dscp",
				"prot",
				"s_from_port",
				"s_to_port",
				"stateful",
				"tcp_rules",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
