package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const managedNodeConnectivityGroupClassName = "mgmtGrp"

type ManagedNodeConnectivityGroupGenerator struct {
	ACIService
}

func (a *ManagedNodeConnectivityGroupGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, managedNodeConnectivityGroupClassName)
	ManagedNodeConnectivityGroupCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	ManagedNodeConnectivityGroupCount, err := strconv.Atoi(stripQuotes(ManagedNodeConnectivityGroupCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < ManagedNodeConnectivityGroupCount; i++ {
		ManagedNodeConnectivityGroupAttr := ManagedNodeConnectivityGroupCont.S("imdata").Index(i).S(managedNodeConnectivityGroupClassName, "attributes")
		ManagedNodeConnectivityGroupDN := G(ManagedNodeConnectivityGroupAttr, "dn")
		if filterChildrenDn(ManagedNodeConnectivityGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ManagedNodeConnectivityGroupDN,
				resourceNamefromDn(managedNodeConnectivityGroupClassName, ManagedNodeConnectivityGroupDN, i),
				"aci_managed_node_connectivity_group",
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
