package aci

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const lDAPGroupMapClassName = "aaaLdapGroupMap"

type LDAPGroupMapGenerator struct {
	ACIService
}

func (a *LDAPGroupMapGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, lDAPGroupMapClassName)

	LDAPGroupMapCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LDAPGroupMapCount, err := strconv.Atoi(stripQuotes(LDAPGroupMapCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LDAPGroupMapCount; i++ {
		LDAPGroupMapAttr := LDAPGroupMapCont.S("imdata").Index(i).S(lDAPGroupMapClassName, "attributes")
		LDAPGroupMapDN := G(LDAPGroupMapAttr, "dn")
		re := regexp.MustCompile("uni/userext/ldapext/ldapgroupmap-(.)+")
		match := re.FindStringSubmatch(LDAPGroupMapDN)
		LDAPGroupMapType := "duo"
		if len(match) > 0 {
			LDAPGroupMapType = "ldap"
		}
		if filterChildrenDn(LDAPGroupMapDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				LDAPGroupMapDN,
				resourceNamefromDn(lDAPGroupMapClassName, LDAPGroupMapDN, i),
				"aci_ldap_group_map",
				"aci",
				map[string]string{
					"type": LDAPGroupMapType,
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
