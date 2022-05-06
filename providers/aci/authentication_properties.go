package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const aAAAuthenticationClassName = "aaaAuthRealm"
const aAAAuthenticationPingEpClassName = "aaaPingEp"

type AAAAuthenticationGenerator struct {
	ACIService
}

func (a *AAAAuthenticationGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, aAAAuthenticationClassName)
	dnPingEpURL := fmt.Sprintf("%s/%s.json", baseURL, aAAAuthenticationPingEpClassName)

	AAAAuthenticationCont, err := client.GetViaURL(dnURL)
	AAAAuthenticationPingEpCont, err := client.GetViaURL(dnPingEpURL)
	if err != nil {
		return err
	}
	AAAAuthenticationCount, err := strconv.Atoi(stripQuotes(AAAAuthenticationCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	AAAAuthenticationPingEpCount, err := strconv.Atoi(stripQuotes(AAAAuthenticationPingEpCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < AAAAuthenticationCount; i++ {
		AAAAuthenticationAttr := AAAAuthenticationCont.S("imdata").Index(i).S(aAAAuthenticationClassName, "attributes")
		AAAAuthenticationDN := G(AAAAuthenticationAttr,"dn")
		if filterChildrenDn(AAAAuthenticationDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					AAAAuthenticationDN,
					resourceNamefromDn(aAAAuthenticationClassName,AAAAuthenticationDN,i),
					"aci_authentication_properties",
					"aci",
					map[string]string{
					},
					[]string{},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
		}	
	}
	for i := 0; i < AAAAuthenticationPingEpCount; i++ {
		AAAAuthenticationAttr := AAAAuthenticationPingEpCont.S("imdata").Index(i).S(aAAAuthenticationPingEpClassName, "attributes")
		AAAAuthenticationDN := G(AAAAuthenticationAttr,"dn")
		if filterChildrenDn(AAAAuthenticationDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					AAAAuthenticationDN,
					resourceNamefromDn(aAAAuthenticationPingEpClassName,AAAAuthenticationDN,i),
					"aci_authentication_properties",
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