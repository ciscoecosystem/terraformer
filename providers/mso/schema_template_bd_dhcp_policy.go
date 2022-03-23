package mso

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaTemplateBdDhcpPolicyGenerator struct {
	MSOService
}

func (a *SchemaTemplateBdDhcpPolicyGenerator) InitResources() error {
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
		siteLen := 0
		if schemaCont.Exists("templates") {
			siteLen = len(schemaCont.S("templates").Data().([]interface{}))
		}

		for j := 0; j < siteLen; j++ {
			templateCont := schemaCont.S("templates").Index(j)
			templateName := stripQuotes(templateCont.S("name").String())
			bdsLen := 0
			if templateCont.Exists("bds") {
				bdsLen = len(templateCont.S("bds").Data().([]interface{}))
			}

			for k := 0; k < bdsLen; k++ {
				bdCont := templateCont.S("bds").Index(k)
				bdName := stripQuotes(bdCont.S("name").String())
				dhcpLen := 0
				if bdCont.Exists("dhcpLabels") {
					dhcpLen = len(bdCont.S("dhcpLabels").Data().([]interface{}))
				}

				for m := 0; m < dhcpLen; m++ {
					dhcpCont := bdCont.S("dhcpLabels").Index(m)
					name := stripQuotes(dhcpCont.S("name").String())
					resourceName := schemaId + "_" + templateName + "_" + bdName + "_" + name
					dhcpId := fmt.Sprintf("/schemas/%s/templates/%s/bds/%s/dhcpLabels/%s", schemaId, templateName, bdName, name)
					resource := terraformutils.NewResource(
						dhcpId,
						resourceName,
						"mso_schema_template_bd_dhcp_policy",
						"mso",
						map[string]string{
							"schema_id":     schemaId,
							"template_name": templateName,
							"bd_name":       bdName,
							"name":          name,
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
