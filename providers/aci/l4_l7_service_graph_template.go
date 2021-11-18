package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/Jeffail/gabs"
)

const L4L7ServiceGraphTemplateClassName = "vnsAbsGraph"

type L4L7ServiceGraphTemplateGenerator struct {
	ACIService
}

func (a *L4L7ServiceGraphTemplateGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, L4L7ServiceGraphTemplateClassName)

	l4L7ServiceGraphTemplateCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	l4L7ServiceGraphTemplateCount, err := strconv.Atoi(stripQuotes(l4L7ServiceGraphTemplateCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < l4L7ServiceGraphTemplateCount; i++ {
		l4L7ServiceGraphTemplateAttr := l4L7ServiceGraphTemplateCont.S("imdata").Index(i).S(L4L7ServiceGraphTemplateClassName, "attributes")
		l4L7ServiceGraphTemplateDN := G(l4L7ServiceGraphTemplateAttr, "dn")
		name := G(l4L7ServiceGraphTemplateAttr, "name")
		if filterChildrenDn(l4L7ServiceGraphTemplateDN, client.parentResource) != "" {
			vnsAbsTermNodeConAttr, _ := readConnectionAttributes(client, l4L7ServiceGraphTemplateDN, "vnsAbsTermNodeCon")
			vnsAbsTermNodeProvAttr, _ := readConnectionAttributes(client, l4L7ServiceGraphTemplateDN, "vnsAbsTermNodeProv")
			vnsAbsTermConnAttr, _ := readConnectionAttributes(client, l4L7ServiceGraphTemplateDN, "vnsAbsTermConn")

			resource := terraformutils.NewResource(
				l4L7ServiceGraphTemplateDN,
				l4L7ServiceGraphTemplateDN,
				"aci_l4_l7_service_graph_template",
				"aci",
				map[string]string{
					"tenant_dn":         GetParentDn(l4L7ServiceGraphTemplateDN, fmt.Sprintf("/AbsGraph-%s", name)),
					"term_cons_name":    G(vnsAbsTermNodeConAttr, "name"),
					"term_prov_name":    G(vnsAbsTermNodeProvAttr, "name"),
					"term_node_cons_dn": G(vnsAbsTermNodeConAttr, "dn"),
					"term_node_prov_dn": G(vnsAbsTermNodeProvAttr, "dn"),
					"term_cons_dn":      G(vnsAbsTermConnAttr, "dn"),
					"term_prov_dn":      G(vnsAbsTermConnAttr, "dn"),
				},
				[]string{
					"description",
					"annotation",
					"name_alias",
					"ui_template_type",
					"l4_l7_service_graph_template_type",
					"name_alias",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}

func readConnectionAttributes(client *ACIClient, parentDn, class string) (*gabs.Container, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, class)
	cont, err := client.GetViaURL(dnUrl)
	if err != nil {
		return nil, err
	}
	return cont.S("imdata").Index(0).S(class, "attributes"), nil
}
