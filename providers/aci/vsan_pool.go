package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const vSANPoolClassName = "fvnsVsanInstP"

type VSANPoolGenerator struct {
	ACIService
}

func (a *VSANPoolGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, vSANPoolClassName)

	VSANPoolCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	VSANPoolCount, err := strconv.Atoi(stripQuotes(VSANPoolCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < VSANPoolCount; i++ {
		VSANPoolAttr := VSANPoolCont.S("imdata").Index(i).S(vSANPoolClassName, "attributes")
		VSANPoolDN := G(VSANPoolAttr, "dn")
		if filterChildrenDn(VSANPoolDN, client.parentResource) != "" {

			resource := terraformutils.NewResource(
				VSANPoolDN,
				fmt.Sprintf("%s_%d", vSANPoolClassName, i),
				"aci_vsan_pool",
				"aci",
				map[string]string{},
				[]string{
					"description",
					"annotation",
					"name_alias",
				},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
