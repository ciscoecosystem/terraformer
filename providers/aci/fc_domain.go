package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const fCDomainClassName = "fcDomP"

type FCDomainGenerator struct {
	ACIService
}

func (a *FCDomainGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, fCDomainClassName)

	FCDomainCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	FCDomainCount, err := strconv.Atoi(stripQuotes(FCDomainCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < FCDomainCount; i++ {
		FCDomainAttr := FCDomainCont.S("imdata").Index(i).S(fCDomainClassName, "attributes")
		FCDomainDN := G(FCDomainAttr, "dn")
		if filterChildrenDn(FCDomainDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				FCDomainDN,
				FCDomainDN,
				"aci_fc_domain",
				"aci",
				map[string]string{},
				[]string{
					"annotation",
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
