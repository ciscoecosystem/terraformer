package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tACACSMonitoringDestinationGroupClassName = "tacacsGroup"

type TACACSMonitoringDestinationGroupGenerator struct {
	ACIService
}

func (a *TACACSMonitoringDestinationGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tACACSMonitoringDestinationGroupClassName)

	TACACSMonitoringDestinationGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	TACACSMonitoringDestinationGroupCount, err := strconv.Atoi(stripQuotes(TACACSMonitoringDestinationGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < TACACSMonitoringDestinationGroupCount; i++ {
		TACACSMonitoringDestinationGroupAttr := TACACSMonitoringDestinationGroupCont.S("imdata").Index(i).S(tACACSMonitoringDestinationGroupClassName, "attributes")
		TACACSMonitoringDestinationGroupDN := G(TACACSMonitoringDestinationGroupAttr, "dn")
		if filterChildrenDn(TACACSMonitoringDestinationGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				TACACSMonitoringDestinationGroupDN,
				resourceNamefromDn(tACACSMonitoringDestinationGroupClassName, TACACSMonitoringDestinationGroupDN, i),
				"aci_tacacs_accounting",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
