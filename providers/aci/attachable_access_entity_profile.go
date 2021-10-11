package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const AttachableAccessEntityProfileClass = "infraAttEntityP"

type AttachableAccessEntityProfileGenerator struct {
	ACIService
}

func (a *AttachableAccessEntityProfileGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, AttachableAccessEntityProfileClass)

	AttachableAccessEntityProfileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	AttachableAccessEntityProfileCount, err := strconv.Atoi(stripQuotes(AttachableAccessEntityProfileCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < AttachableAccessEntityProfileCount; i++ {
		AttachableAccessEntityProfileDN := stripQuotes(AttachableAccessEntityProfileCont.S("imdata").Index(i).S(AttachableAccessEntityProfileClass, "attributes", "dn").String())
		if filterChildrenDn(AttachableAccessEntityProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				AttachableAccessEntityProfileDN,
				resourceNamefromDn(AttachableAccessEntityProfileClass,AttachableAccessEntityProfileDN,i),
				"aci_attachable_access_entity_profile",
				"aci",
				[]string{
					"name_alias",
					"relation_infra_rs_dom_p",
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
