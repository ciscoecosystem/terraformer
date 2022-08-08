package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tACACSSourceClassName = "tacacsSrc"

type TACACSSourceGenerator struct {
	ACIService
}

func (a *TACACSSourceGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tACACSSourceClassName)

	TACACSSourceCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	TACACSSourceCount, err := strconv.Atoi(stripQuotes(TACACSSourceCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < TACACSSourceCount; i++ {
		TACACSSourceAttr := TACACSSourceCont.S("imdata").Index(i).S(tACACSSourceClassName, "attributes")
		TACACSSourceDN := G(TACACSSourceAttr, "dn")
		name := G(TACACSSourceAttr, "name")
		if filterChildrenDn(TACACSSourceDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				TACACSSourceDN,
				resourceNamefromDn(tACACSSourceClassName, TACACSSourceDN, i),
				"aci_tacacs_source",
				"aci",
				map[string]string{
					"parent_dn": GetParentDn(TACACSSourceDN, fmt.Sprintf("/tacacssrc-%s", name)),
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
