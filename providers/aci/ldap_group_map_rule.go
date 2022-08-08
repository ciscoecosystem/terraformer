package aci

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const lDAPGroupMapRuleClassName = "aaaLdapGroupMapRule"

type LDAPGroupMapRuleGenerator struct {
	ACIService
}

func (a *LDAPGroupMapRuleGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, lDAPGroupMapRuleClassName)

	LDAPGroupMapRuleCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LDAPGroupMapRuleCount, err := strconv.Atoi(stripQuotes(LDAPGroupMapRuleCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LDAPGroupMapRuleCount; i++ {
		LDAPGroupMapRuleAttr := LDAPGroupMapRuleCont.S("imdata").Index(i).S(lDAPGroupMapRuleClassName, "attributes")
		LDAPGroupMapRuleDN := G(LDAPGroupMapRuleAttr, "dn")
		arr := strings.Split(LDAPGroupMapRuleDN, "/")
		type_grp := ""
		if arr[2] == "duoext" {
			type_grp = "duo"
		} else {
			type_grp = "ldap"
		}
		if filterChildrenDn(LDAPGroupMapRuleDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				LDAPGroupMapRuleDN,
				resourceNamefromDn(lDAPGroupMapRuleClassName, LDAPGroupMapRuleDN, i),
				"aci_ldap_group_map_rule",
				"aci",
				map[string]string{
					"type": type_grp,
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
