package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const mgmtconnectivitypreferenceClassName = "mgmtConnectivityPrefs"

type MgmtconnectivitypreferenceGenerator struct {
	ACIService
}

func (a *MgmtconnectivitypreferenceGenerator) InitResources() error {
	client, err := a.createClient()
	if err != nil {
		return err
	}

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, mgmtconnectivitypreferenceClassName)

	MgmtconnectivitypreferenceCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	MgmtconnectivitypreferenceCount, err := strconv.Atoi(stripQuotes(MgmtconnectivitypreferenceCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < MgmtconnectivitypreferenceCount; i++ {
		MgmtconnectivitypreferenceAttr := MgmtconnectivitypreferenceCont.S("imdata").Index(i).S(mgmtconnectivitypreferenceClassName, "attributes")
		MgmtconnectivitypreferenceDN := G(MgmtconnectivitypreferenceAttr, "dn")
		if filterChildrenDn(MgmtconnectivitypreferenceDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				MgmtconnectivitypreferenceDN,
				MgmtconnectivitypreferenceDN,
				"aci_mgmt_preference",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
