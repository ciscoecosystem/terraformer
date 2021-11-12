package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l4L7ServiceGraphTemplateClassName = "vnsAbsGraph"

type L4L7ServiceGraphTemplateGenerator struct {
	ACIService
}

func (a *L4L7ServiceGraphTemplateGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l4L7ServiceGraphTemplateClassName)

	L4L7ServiceGraphTemplateCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	L4L7ServiceGraphTemplateCount, err := strconv.Atoi(stripQuotes(L4L7ServiceGraphTemplateCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < L4L7ServiceGraphTemplateCount; i++ {
		L4L7ServiceGraphTemplateAttr := L4L7ServiceGraphTemplateCont.S("imdata").Index(i).S(l4L7ServiceGraphTemplateClassName, "attributes")
		L4L7ServiceGraphTemplateDN := G(L4L7ServiceGraphTemplateAttr,"dn")
		name := G(L4L7ServiceGraphTemplateAttr,"name")
		if filterChildrenDn(L4L7ServiceGraphTemplateDN, client.parentResource) != "" {
			

			resource := terraformutils.NewResource(
					L4L7ServiceGraphTemplateDN,
					L4L7ServiceGraphTemplateDN,
					"aci_l4_l7_service_graph_template",
					"aci",
					map[string]string{
						"tenant_dn": GetParentDn(L4L7ServiceGraphTemplateDN, fmt.Sprintf("/AbsGraph-%s", name,)),
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