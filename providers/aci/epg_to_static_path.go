package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EPGToStaticPathClass = "fvRsPathAtt"

type EPGToStaticPathGenerator struct {
	ACIService
}

func (a *EPGToStaticPathGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EPGToStaticPathClass)

	EGPToStaticPathCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EGPToStaticPathCount, err := strconv.Atoi(stripQuotes(EGPToStaticPathCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < EGPToStaticPathCount; i++ {
		EGPToStaticPathDN := stripQuotes(EGPToStaticPathCont.S("imdata").Index(i).S(EPGToStaticPathClass, "attributes", "dn").String())
		if filterChildrenDn(EGPToStaticPathDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				EGPToStaticPathDN,
				EGPToStaticPathDN,
				"aci_epg_to_static_path",
				"aci",
				[]string{
					"encap",
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
