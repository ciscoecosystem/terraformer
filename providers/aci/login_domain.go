package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const loginDomainClassName = "aaaLoginDomain"

type LoginDomainGenerator struct {
	ACIService
}

func (a *LoginDomainGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, loginDomainClassName)

	LoginDomainCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LoginDomainCount, err := strconv.Atoi(stripQuotes(LoginDomainCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LoginDomainCount; i++ {
		LoginDomainAttr := LoginDomainCont.S("imdata").Index(i).S(loginDomainClassName, "attributes")
		LoginDomainDN := G(LoginDomainAttr,"dn")
		name := G(LoginDomainAttr,"name")
		if filterChildrenDn(LoginDomainDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					LoginDomainDN,
					resourceNamefromDn(loginDomainClassName,LoginDomainDN,i),
					"aci_login_domain",
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