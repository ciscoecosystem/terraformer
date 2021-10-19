package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EPGUsingFunctionClass = "infraRsFuncToEpg"

type EPGUsingFunctionGenerator struct {
	ACIService
}

func (a *EPGUsingFunctionGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EPGUsingFunctionClass)

	EPGUsingFunctionCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGUsingFunctionCount, err := strconv.Atoi(stripQuotes(EPGUsingFunctionCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < EPGUsingFunctionCount; i++ {
		EPGUsingFunctionDN := stripQuotes(EPGUsingFunctionCont.S("imdata").Index(i).S(EPGUsingFunctionClass, "attributes", "dn").String())
		if filterChildrenDn(EPGUsingFunctionDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				EPGUsingFunctionDN,
				EPGUsingFunctionDN,
				"aci_epgs_using_function",
				"aci",
				[]string{
					"instr_imedcy",
					"mode",
					"primary_encap",
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
