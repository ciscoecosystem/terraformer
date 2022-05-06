package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const tagOriginalClassName = "tagTag"

type tagOriginalGenerator struct {
	ACIService
}

func (a *tagOriginalGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, tagOriginalClassName)

	TagCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	TagCount, err := strconv.Atoi(stripQuotes(TagCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < TagCount; i++ {
		TagAttr := TagCont.S("imdata").Index(i).S(tagOriginalClassName, "attributes")
		TagDN := G(TagAttr, "dn")
		key := G(TagAttr, "key")
		if filterChildrenDn(TagDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				TagDN,
				resourceNamefromDn(tagOriginalClassName, TagDN, i),
				"aci_tag",
				"aci",
				map[string]string{
					"parent_dn": GetParentDn(TagDN, fmt.Sprintf("/tagKey-%s", key)),
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
