package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const lDAPProviderClassName = "aaaLdapProvider"

type LDAPProviderGenerator struct {
	ACIService
}

func (a *LDAPProviderGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, lDAPProviderClassName)

	LDAPProviderCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LDAPProviderCount, err := strconv.Atoi(stripQuotes(LDAPProviderCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LDAPProviderCount; i++ {
		LDAPProviderAttr := LDAPProviderCont.S("imdata").Index(i).S(lDAPProviderClassName, "attributes")
		LDAPProviderDN := G(LDAPProviderAttr, "dn")
		if filterChildrenDn(LDAPProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				LDAPProviderDN,
				resourceNamefromDn(lDAPProviderClassName, LDAPProviderDN, i),
				"aci_ldap_provider",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
