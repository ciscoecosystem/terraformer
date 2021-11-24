package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const ImportedContractClass = "vzCPIf"

type ImportedContractGenerator struct {
	ACIService
}

func (a *ImportedContractGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, ImportedContractClass)

	importedContractCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	importedContractCount, err := strconv.Atoi(stripQuotes(importedContractCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < importedContractCount; i++ {
		importedContractDN := stripQuotes(importedContractCont.S("imdata").Index(i).S(ImportedContractClass, "attributes", "dn").String())
		if filterChildrenDn(importedContractDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				importedContractDN,
				resourceNamefromDn(ImportedContractClass, (importedContractDN), i),
				"aci_imported_contract",
				"aci",
				[]string{
					"name_alias",
					"relation_vz_rs_if",
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
