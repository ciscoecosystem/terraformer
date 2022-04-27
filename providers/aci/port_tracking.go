package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const portTrackingClassName = "infraPortTrackPol"

type PortTrackingGenerator struct {
	ACIService
}

func (a *PortTrackingGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, portTrackingClassName)

	PortTrackingCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	PortTrackingCount, err := strconv.Atoi(stripQuotes(PortTrackingCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < PortTrackingCount; i++ {
		PortTrackingAttr := PortTrackingCont.S("imdata").Index(i).S(portTrackingClassName, "attributes")
		PortTrackingDN := G(PortTrackingAttr, "dn")
		if filterChildrenDn(PortTrackingDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				PortTrackingDN,
				resourceNamefromDn(portTrackingClassName, PortTrackingDN, i),
				"aci_port_tracking",
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
