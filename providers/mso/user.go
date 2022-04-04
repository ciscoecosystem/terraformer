package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type UserGenerator struct {
	MSOService
}

func (a *UserGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/users/")
	if err != nil {
		return err
	}

	usersLen := len(con.S("users").Data().([]interface{}))

	for i := 0; i < usersLen; i++ {
		userCont := con.S("users").Index(i)
		userName := models.G(userCont, "username")
		password := models.G(userCont, "password")
		id := models.G(userCont, "id")

		var firstName, lastName, emailAddress, phoneNumber, accountStatus, domain string

		if userCont.Exists("firstName") {
			firstName = models.G(userCont, "firstName")
		}
		if userCont.Exists("lastName") {
			lastName = models.G(userCont, "lastName")
		}
		if userCont.Exists("emailAddress") {
			emailAddress = models.G(userCont, "emailAddress")
		}
		if userCont.Exists("phoneNumber") {
			phoneNumber = models.G(userCont, "phoneNumber")
		}
		if userCont.Exists("accountStatus") {
			accountStatus = models.G(userCont, "accountStatus")
		}
		if userCont.Exists("domain") {
			domain = models.G(userCont, "domain")
		}

		rolesLen := len(userCont.S("roles").Data().([]interface{}))
		roles := make([]interface{}, 0)
		for j := 0; j < rolesLen; j++ {
			roleCont := userCont.S("roles").Index(j)

			map_role := make(map[string]interface{})

			map_role["roleid"] = models.G(roleCont, "roleId")
			map_role["access_type"] = models.G(roleCont, "access_type")
			roles = append(roles, map_role)
		}
		resource := terraformutils.NewResource(
			id,
			id,
			"mso_user",
			"mso",
			map[string]string{
				"username":       userName,
				"user_password":  password,
				"first_name":     firstName,
				"last_name":      lastName,
				"email":          emailAddress,
				"phone":          phoneNumber,
				"account_status": accountStatus,
				"domain":         domain,
			},
			[]string{},
			map[string]interface{}{
				"roles": roles,
			},
		)
		resource.SlowQueryRequired = SlowQueryRequired
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
