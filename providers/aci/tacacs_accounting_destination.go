package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const TACACSDestinationClassName = "tacacsTacacsDest"

type TACACSDestinationGenerator struct {
	ACIService
}

func (a *TACACSDestinationGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, TACACSDestinationClassName)
	TACACSDestinationCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	TACACSDestinationCount, err := strconv.Atoi(stripQuotes(TACACSDestinationCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < TACACSDestinationCount; i++ {
		TACACSDestinationAttr := TACACSDestinationCont.S("imdata").Index(i).S(TACACSDestinationClassName, "attributes")
		TACACSDestinationDN := G(TACACSDestinationAttr,"dn")
		host := G(TACACSDestinationAttr,"host")
		port := G(TACACSDestinationAttr,"port")
		if filterChildrenDn(TACACSDestinationDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
					TACACSDestinationDN,
					resourceNamefromDn(TACACSDestinationClassName,TACACSDestinationDN,i),
					"aci_tacacs_accounting_destination",
					"aci",
					map[string]string{
						"tacacs_accounting_dn": GetParentDn(TACACSDestinationDN, fmt.Sprintf("/tacacsdest-%s-port-%s", host,port)),
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