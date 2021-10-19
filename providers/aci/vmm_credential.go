package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const VmmCredentialClass = "vmmUsrAccP"

type VmmCredentialGenerator struct {
	ACIService
}

func (a *VmmCredentialGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, VmmCredentialClass)

	VmmCredentialCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	VmmCredentialCount, err := strconv.Atoi(stripQuotes(VmmCredentialCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < VmmCredentialCount; i++ {
		VmmCredentialDN := stripQuotes(VmmCredentialCont.S("imdata").Index(i).S(VmmCredentialClass, "attributes", "dn").String())
		if filterChildrenDn(VmmCredentialDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				VmmCredentialDN,
				VmmCredentialDN,
				"aci_vmm_credential",
				"aci",
				[]string{
					"pwd",
					"usr",
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
