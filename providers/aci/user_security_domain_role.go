package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const userRoleClassName = "aaaUserRole"

type UserRoleGenerator struct {
	ACIService
}

func (a *UserRoleGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, userRoleClassName)

	UserRoleCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	UserRoleCount, err := strconv.Atoi(stripQuotes(UserRoleCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < UserRoleCount; i++ {
		UserRoleAttr := UserRoleCont.S("imdata").Index(i).S(userRoleClassName, "attributes")
		UserRoleDN := G(UserRoleAttr, "dn")
		name := G(UserRoleAttr, "name")
		if filterChildrenDn(UserRoleDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				UserRoleDN,
				resourceNamefromDn(userRoleClassName, UserRoleDN, i),
				"aci_user_security_domain_role",
				"aci",
				map[string]string{
					"user_domain_dn": GetParentDn(UserRoleDN, fmt.Sprintf("/role-%s", name)),
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
