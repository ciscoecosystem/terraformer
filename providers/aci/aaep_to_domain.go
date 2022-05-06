package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const domainClassName = "infraRsDomP"

type DomainGenerator struct {
	ACIService
}

func (a *DomainGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, domainClassName)

	DomainCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	DomainCount, err := strconv.Atoi(stripQuotes(DomainCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < DomainCount; i++ {
		DomainAttr := DomainCont.S("imdata").Index(i).S(domainClassName, "attributes")
		DomainDN := G(DomainAttr, "dn")
		tDn := G(DomainAttr, "tDn")
		fmt.Printf("ParentDN: %v\n", GetParentDn(DomainDN, fmt.Sprintf("/rsdomP-[%s]", tDn)))
		if filterChildrenDn(DomainDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				DomainDN,
				resourceNamefromDn(domainClassName, DomainDN, i),
				"aci_aaep_to_domain",
				"aci",
				map[string]string{
					"attachable_access_entity_profile_dn": GetParentDn(DomainDN, fmt.Sprintf("/rsdomP-[%s]", tDn)),
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
