package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const providerGroupMemberClassName = "aaaProviderRef"

type ProviderGroupMemberGenerator struct {
	ACIService
}

func (a *ProviderGroupMemberGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, providerGroupMemberClassName)

	ProviderGroupMemberCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ProviderGroupMemberCount, err := strconv.Atoi(stripQuotes(ProviderGroupMemberCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ProviderGroupMemberCount; i++ {
		ProviderGroupMemberAttr := ProviderGroupMemberCont.S("imdata").Index(i).S(providerGroupMemberClassName, "attributes")
		ProviderGroupMemberDN := G(ProviderGroupMemberAttr, "dn")
		name := G(ProviderGroupMemberAttr, "name")
		if filterChildrenDn(ProviderGroupMemberDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ProviderGroupMemberDN,
				resourceNamefromDn(providerGroupMemberClassName, ProviderGroupMemberDN, i),
				"aci_login_domain_provider",
				"aci",
				map[string]string{
					"parent_dn": GetParentDn(ProviderGroupMemberDN, fmt.Sprintf("/providerref-%s", name)),
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
