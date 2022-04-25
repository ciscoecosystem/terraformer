package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fabricNodeControlClassName = "fabricNodeControl"

type FabricNodeControlGenerator struct {
	ACIService
}

func (a *FabricNodeControlGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fabricNodeControlClassName)

	FabricNodeControlCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FabricNodeControlCount, err := strconv.Atoi(stripQuotes(FabricNodeControlCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FabricNodeControlCount; i++ {
		FabricNodeControlAttr := FabricNodeControlCont.S("imdata").Index(i).S(fabricNodeControlClassName, "attributes")
		FabricNodeControlDN := G(FabricNodeControlAttr, "dn")
		name := G(FabricNodeControlAttr, "name")
		if filterChildrenDn(FabricNodeControlDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				FabricNodeControlDN,
				name,
				"aci_fabric_node_control",
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
