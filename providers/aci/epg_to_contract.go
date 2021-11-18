package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EPGToContractClass = "fvRsProv"

type EPGToContractGenerator struct {
	ACIService
}

func (a *EPGToContractGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EPGToContractClass)

	EPGToContractCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGToContractCount, err := strconv.Atoi(stripQuotes(EPGToContractCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < EPGToContractCount; i++ {
		EPGToContractDN := stripQuotes(EPGToContractCont.S("imdata").Index(i).S(EPGToContractClass, "attributes", "dn").String())
		if filterChildrenDn(EPGToContractDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				EPGToContractDN,
				EPGToContractDN,
				"aci_epg_to_contract",
				"aci",
				map[string]string{
					"contract_type": "provider",
				},
				[]string{
					"annotation",
					"match_t",
					"prio",
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	// Consumer
	dnURL = fmt.Sprintf("%s/%s.json", baseURL, "fvRsCons")

	EPGToContractCont, err = client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGToContractCount, err = strconv.Atoi(stripQuotes(EPGToContractCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < EPGToContractCount; i++ {
		EPGToContractDN := stripQuotes(EPGToContractCont.S("imdata").Index(i).S("fvRsCons", "attributes", "dn").String())
		if filterChildrenDn(EPGToContractDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				EPGToContractDN,
				EPGToContractDN,
				"aci_epg_to_contract",
				"aci",
				map[string]string{
					"contract_type": "consumer",
				},
				[]string{
					"annotation",
					"match_t",
					"prio",
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
