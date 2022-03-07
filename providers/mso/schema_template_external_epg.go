package mso

import (
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type SchemaTemplateExternalEPG struct {
	MSOService
}

func (a *SchemaTemplateExternalEPG) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/schemas")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("schemas").Data().([]interface{})); i++ {
		schemaCont := con.S("schemas").Index(i)
		schemaId := stripQuotes(schemaCont.S("id").String())

		for j := 0; j < len(schemaCont.S("templates").Data().([]interface{})); j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())

			for k := 0; k < len(templateCont.S("bds").Data().([]interface{})); k++ {
				bdCont := templateCont.S("bds").Index(k)
				bdName := stripQuotes(bdCont.S("name").String())

				for m := 0; m < len(bdCont.S("subnets").Data().([]interface{})); m++ {
					subnetCont := bdCont.S("subnets").Index(m)
					subnetIp := stripQuotes(subnetCont.S("ip").String())
					subnetIpArray := strings.Split(subnetIp, "/")
					subnetID := subnetIpArray[0]
					subnetScope := stripQuotes(subnetCont.S("scope").String())
					resourceName := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k) + "_" + strconv.Itoa(m)
					resource := terraformutils.NewResource(
						subnetID,
						resourceName,
						"mso_schema_template_bd_subnet",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"bd_name":       bdName,
							"ip":            subnetIp,
							"scope":         subnetScope,
						},
						[]string{},
						map[string]interface{}{},
					)
					resource.SlowQueryRequired = true
					a.Resources = append(a.Resources, resource)
				}
			}
		}
	}
	return nil
}
