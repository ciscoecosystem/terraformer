package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const defaultAuthenticationMethodforallLoginsClassName = "aaaDefaultAuth"

type DefaultAuthenticationMethodforallLoginsGenerator struct {
	ACIService
}

func (a *DefaultAuthenticationMethodforallLoginsGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, defaultAuthenticationMethodforallLoginsClassName)

	DefaultAuthenticationMethodforallLoginsCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	DefaultAuthenticationMethodforallLoginsCount, err := strconv.Atoi(stripQuotes(DefaultAuthenticationMethodforallLoginsCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < DefaultAuthenticationMethodforallLoginsCount; i++ {
		DefaultAuthenticationMethodforallLoginsAttr := DefaultAuthenticationMethodforallLoginsCont.S("imdata").Index(i).S(defaultAuthenticationMethodforallLoginsClassName, "attributes")
		DefaultAuthenticationMethodforallLoginsDN := G(DefaultAuthenticationMethodforallLoginsAttr, "dn")
		if filterChildrenDn(DefaultAuthenticationMethodforallLoginsDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				DefaultAuthenticationMethodforallLoginsDN,
				resourceNamefromDn(defaultAuthenticationMethodforallLoginsClassName, DefaultAuthenticationMethodforallLoginsDN, i),
				"aci_default_authentication",
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
