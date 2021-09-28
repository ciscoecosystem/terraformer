package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const l3OutVPCMemberClass = "l3extMember"

type L3OutVPCMemberGenerator struct {
	ACIService
}

func (a *L3OutVPCMemberGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, l3OutVPCMemberClass)

	l3OutVPCMemberCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(l3OutVPCMemberCont.S("totalCount").String())

	if totalCount == "{}"{
		totalCount = "0"
	}

	l3OutVPCMemberCount, err := strconv.Atoi(stripQuotes(l3OutVPCMemberCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < l3OutVPCMemberCount; i++ {
		l3OutVPCMemberProfileDN := stripQuotes(l3OutVPCMemberCont.S("imdata").Index(i).S(l3OutVPCMemberClass, "attributes", "dn").String())
		resource := terraformutils.NewSimpleResource(
			l3OutVPCMemberProfileDN,
			l3OutVPCMemberProfileDN,
			"aci_l3out_vpc_member",
			"aci",
			[]string{
				"addr",
				"ipv6_dad",
				"ll_addr",
				"name_alias",
				"annotation",
				"description",
			},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}

	return nil
}