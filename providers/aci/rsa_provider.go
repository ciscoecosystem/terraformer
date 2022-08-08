package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const rSAProviderClassName = "aaaRsaProvider"

type RSAProviderGenerator struct {
	ACIService
}

func (a *RSAProviderGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, rSAProviderClassName)

	RSAProviderCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RSAProviderCount, err := strconv.Atoi(stripQuotes(RSAProviderCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RSAProviderCount; i++ {
		RSAProviderAttr := RSAProviderCont.S("imdata").Index(i).S(rSAProviderClassName, "attributes")
		RSAProviderDN := G(RSAProviderAttr, "dn")
		if filterChildrenDn(RSAProviderDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RSAProviderDN,
				resourceNamefromDn(rSAProviderClassName, RSAProviderDN, i),
				"aci_rsa_provider",
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
