package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const destinationOfRedirectedTrafficClassName = "vnsRedirectDest"

type DestinationOfRedirectedTrafficGenerator struct {
	ACIService
}

func (a *DestinationOfRedirectedTrafficGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, destinationOfRedirectedTrafficClassName)

	DestinationOfRedirectedTrafficCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	DestinationOfRedirectedTrafficCount, err := strconv.Atoi(stripQuotes(DestinationOfRedirectedTrafficCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < DestinationOfRedirectedTrafficCount; i++ {
		DestinationOfRedirectedTrafficAttr := DestinationOfRedirectedTrafficCont.S("imdata").Index(i).S(destinationOfRedirectedTrafficClassName, "attributes")
		DestinationOfRedirectedTrafficDN := G(DestinationOfRedirectedTrafficAttr, "dn")
		ip := G(DestinationOfRedirectedTrafficAttr, "ip")
		if filterChildrenDn(DestinationOfRedirectedTrafficDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				DestinationOfRedirectedTrafficDN,
				DestinationOfRedirectedTrafficDN,
				"aci_destination_of_redirected_traffic",
				"aci",
				map[string]string{
					"service_redirect_policy_dn": GetParentDn(DestinationOfRedirectedTrafficDN, fmt.Sprintf("/RedirectDest_ip-[%s]", ip)),
				},
				[]string{
					"description",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
