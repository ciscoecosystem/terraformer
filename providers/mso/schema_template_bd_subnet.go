package mso

import (
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateBDSubnet struct {
	MSOService
}

func (a *SchemaTemplateBDSubnet) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
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
					resourceName := schemaId + "_" + templateName + "_" + bdName + "_" + strconv.Itoa(int(hash(subnetID)))
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
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}
			}
		}
	}
	return nil
}
