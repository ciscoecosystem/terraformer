package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const userDomainClassName = "aaaUserDomain"

type UserDomainGenerator struct {
	ACIService
}

func (a *UserDomainGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, userDomainClassName)

	UserDomainCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	UserDomainCount, err := strconv.Atoi(stripQuotes(UserDomainCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < UserDomainCount; i++ {
		UserDomainAttr := UserDomainCont.S("imdata").Index(i).S(userDomainClassName, "attributes")
		UserDomainDN := G(UserDomainAttr, "dn")
		name := G(UserDomainAttr, "name")
		if filterChildrenDn(UserDomainDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				UserDomainDN,
				resourceNamefromDn(userDomainClassName, UserDomainDN, i),
				"aci_user_security_domain",
				"aci",
				map[string]string{
					"local_user_dn": GetParentDn(UserDomainDN, fmt.Sprintf("/userdomain-%s", name)),
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
