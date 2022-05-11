package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const userManagementClassName = "aaaUserEp"

type UserManagementGenerator struct {
	ACIService
}

func (a *UserManagementGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, userManagementClassName)

	UserManagementCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	UserManagementCount, err := strconv.Atoi(stripQuotes(UserManagementCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < UserManagementCount; i++ {
		UserManagementAttr := UserManagementCont.S("imdata").Index(i).S(userManagementClassName, "attributes")
		UserManagementDN := G(UserManagementAttr, "dn")
		if filterChildrenDn(UserManagementDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				UserManagementDN,
				resourceNamefromDn(userManagementClassName, UserManagementDN, i),
				"aci_global_security",
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
