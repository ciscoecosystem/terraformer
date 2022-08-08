package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const lDAPGroupMaprulerefClassName = "aaaLdapGroupMapRuleRef"

type LDAPGroupMaprulerefGenerator struct {
	ACIService
}

func (a *LDAPGroupMaprulerefGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, lDAPGroupMaprulerefClassName)

	LDAPGroupMaprulerefCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LDAPGroupMaprulerefCount, err := strconv.Atoi(stripQuotes(LDAPGroupMaprulerefCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LDAPGroupMaprulerefCount; i++ {
		LDAPGroupMaprulerefAttr := LDAPGroupMaprulerefCont.S("imdata").Index(i).S(lDAPGroupMaprulerefClassName, "attributes")
		LDAPGroupMaprulerefDN := G(LDAPGroupMaprulerefAttr, "dn")
		name := G(LDAPGroupMaprulerefAttr, "name")
		if filterChildrenDn(LDAPGroupMaprulerefDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				LDAPGroupMaprulerefDN,
				resourceNamefromDn(lDAPGroupMaprulerefClassName, LDAPGroupMaprulerefDN, i),
				"aci_ldap_group_map_rule_to_group_map",
				"aci",
				map[string]string{
					"ldap_group_map_dn": GetParentDn(LDAPGroupMaprulerefDN, fmt.Sprintf("/ldapgroupmapruleref-%s", name)),
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
