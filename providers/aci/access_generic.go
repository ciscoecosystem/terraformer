package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const accessGenericClass = "infraGeneric"

type AccessGenericGenerator struct {
	ACIService
}

func (a *AccessGenericGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, accessGenericClass)

	accessGenericCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(accessGenericCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	accessGenericCount, err := strconv.Atoi(stripQuotes(accessGenericCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < accessGenericCount; i++ {
		accessGenericProfileDN := stripQuotes(accessGenericCont.S("imdata").Index(i).S(accessGenericClass, "attributes", "dn").String())
		if filterChildrenDn(accessGenericProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				accessGenericProfileDN,
				resourceNamefromDn(accessGenericClass,accessGenericProfileDN,i),
				"aci_access_generic",
				"aci",
				[]string{
					"name_alias",
					"annotation",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
