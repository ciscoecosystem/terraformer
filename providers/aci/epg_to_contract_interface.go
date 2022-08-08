package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const contractInterfaceClassName = "fvRsConsIf"

type ContractInterfaceGenerator struct {
	ACIService
}

func (a *ContractInterfaceGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, contractInterfaceClassName)

	ContractInterfaceCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ContractInterfaceCount, err := strconv.Atoi(stripQuotes(ContractInterfaceCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ContractInterfaceCount; i++ {
		ContractInterfaceAttr := ContractInterfaceCont.S("imdata").Index(i).S(contractInterfaceClassName, "attributes")
		ContractInterfaceDN := G(ContractInterfaceAttr, "dn")
		tnVzCPIfName := G(ContractInterfaceAttr, "tnVzCPIfName")
		tDn := G(ContractInterfaceAttr, "tDn")
		if filterChildrenDn(ContractInterfaceDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ContractInterfaceDN,
				resourceNamefromDn(contractInterfaceClassName, ContractInterfaceDN, i),
				"aci_epg_to_contract_interface",
				"aci",
				map[string]string{
					"application_epg_dn":    GetParentDn(ContractInterfaceDN, fmt.Sprintf("/rsconsIf-%s", tnVzCPIfName)),
					"contract_interface_dn": tDn,
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
