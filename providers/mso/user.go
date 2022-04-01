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
	// con, err := getSchemaContainer(mso)
	if err != nil {
		return err
	}

	userName := models.G(con, "username")
	password := models.G(con, "password")
	id := models.G(con, "id")

	var firstName, lastName, emailAddress, phoneNumber, accountStatus, domain string

	if con.Exists("firstName") {
		firstName = models.G(con, "firstName")
	}
	if con.Exists("lastName") {
		lastName = models.G(con, "lastName")
	}
	if con.Exists("emailAddress") {
		emailAddress = models.G(con, "emailAddress")
	}
	if con.Exists("phoneNumber") {
		phoneNumber = models.G(con, "phoneNumber")
	}
	if con.Exists("accountStatus") {
		accountStatus = models.G(con, "accountStatus")
	}
	if con.Exists("domain") {
		domain = models.G(con, "domain")
	}

	rolesLen := len(con.S("roles").Data().([]interface{}))
	roles := make([]interface{}, 0)
	roleId := ""
	for i := 0; i < rolesLen; i++ {
		roleCont := con.S("roles").Index(i)

		map_role := make(map[string]interface{})

		map_role["roleid"] = models.G(roleCont, "roleId")
		map_role["access_type"] = models.G(roleCont, "access_type")
		roles = append(roles, map_role)
		if i == 0 {
			roleId = models.G(roleCont, "roleId")
		}
	}
	resourceName := roleId
	resource := terraformutils.NewResource(
		id,
		resourceName,
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
	return nil
}
