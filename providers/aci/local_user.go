package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const LocalUserClassName = "aaaUser"

type LocalUserGenerator struct {
	ACIService
}

func (a *LocalUserGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, LocalUserClassName)

	LocalUserCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	LocalUserCount, err := strconv.Atoi(stripQuotes(LocalUserCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < LocalUserCount; i++ {
		LocalUserDN := stripQuotes(LocalUserCont.S("imdata").Index(i).S(LocalUserClassName, "attributes", "dn").String())
		if filterChildrenDn(LocalUserDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				stripQuotes(LocalUserDN),
				stripQuotes(LocalUserDN),
				"aci_local_user",
				"aci",
				[]string{
					"account_status",
					"cert_attribute",
					"clear_pwd_history",
					"email",
					"expiration",
					"expires",
					"first_name",
					"last_name",
					"name_alias",
					"otpenable",
					"otpkey",
					"phone",
					"pwd",
					"pwd_life_time",
					"pwd_update_required",
					"rbac_string",
					"unix_user_id",
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
