package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const outofServiceFabricPathClassName = "fabricRsOosPath"

type OutofServiceFabricPathGenerator struct {
	ACIService
}

func (a *OutofServiceFabricPathGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, outofServiceFabricPathClassName)

	OutofServiceFabricPathCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	OutofServiceFabricPathCount, err := strconv.Atoi(stripQuotes(OutofServiceFabricPathCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < OutofServiceFabricPathCount; i++ {
		OutofServiceFabricPathAttr := OutofServiceFabricPathCont.S("imdata").Index(i).S(outofServiceFabricPathClassName, "attributes")
		OutofServiceFabricPathDN := G(OutofServiceFabricPathAttr, "dn")
		if filterChildrenDn(OutofServiceFabricPathDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				OutofServiceFabricPathDN,
				resourceNamefromDn(outofServiceFabricPathClassName, OutofServiceFabricPathDN, i),
				"aci_interface_blacklist",
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
