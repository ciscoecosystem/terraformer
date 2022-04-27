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
	client, err := a.createClient()
	if err != nil {
		return err
	}
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
		ManagedNodeConnectivityGroupDN := G(ManagedNodeConnectivityGroupAttr,"dn")
		name := G(ManagedNodeConnectivityGroupAttr,"name")
		if filterChildrenDn(ManagedNodeConnectivityGroupDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					ManagedNodeConnectivityGroupDN,
					name,
					"aci_managed_node_connectivity_group",
					"aci",
					map[string]string{
					},
					[]string{
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