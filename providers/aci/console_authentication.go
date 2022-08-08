package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const consoleAuthenticationMethodClassName = "aaaConsoleAuth"

type ConsoleAuthenticationMethodGenerator struct {
	ACIService
}

func (a *ConsoleAuthenticationMethodGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, consoleAuthenticationMethodClassName)

	ConsoleAuthenticationMethodCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ConsoleAuthenticationMethodCount, err := strconv.Atoi(stripQuotes(ConsoleAuthenticationMethodCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < ConsoleAuthenticationMethodCount; i++ {
		ConsoleAuthenticationMethodAttr := ConsoleAuthenticationMethodCont.S("imdata").Index(i).S(consoleAuthenticationMethodClassName, "attributes")
		ConsoleAuthenticationMethodDN := G(ConsoleAuthenticationMethodAttr, "dn")
		if filterChildrenDn(ConsoleAuthenticationMethodDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ConsoleAuthenticationMethodDN,
				resourceNamefromDn(consoleAuthenticationMethodClassName, ConsoleAuthenticationMethodDN, i),
				"aci_console_authentication",
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
