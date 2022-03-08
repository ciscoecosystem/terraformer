package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteVrfRegionCidr struct {
	MSOService
}

func (a *SchemaSiteVrfRegionCidr) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas/")
	if err != nil {
		return err
	}
	schemaLength := len(con.S("schemas").Data().([]interface{}))
	schemaCon := con.S("schemas")
	for i := 0; i < schemaLength; i++ {
		schemaId := stripQuotes(schemaCon.Index(i).S("id").String())
		sitesCon := schemaCon.Index(i).S("sites")
		sitesCount := 0

		if schemaCon.Index(i).Exists("sites") {
			sitesCount = len(schemaCon.Index(i).S("sites").Data().([]interface{}))
		}

		for j := 0; j < sitesCount; j++ {
			siteId := stripQuotes(sitesCon.Index(j).S("siteId").String())
			templateName := stripQuotes(sitesCon.Index(j).S("templateName").String())
			vrfsCount := len(sitesCon.Index(j).S("vrfs").Data().([]interface{}))
			vrfsCon := sitesCon.Index(j).S("vrfs")

			for k := 0; k < vrfsCount; k++ {
				vrfRef := stripQuotes(vrfsCon.Index(k).S("vrfRef").String())
				regionsCount := len(vrfsCon.Index(k).S("regions").Data().([]interface{}))
				regionsCon := vrfsCon.Index(k).S("regions")

				for l := 0; l < regionsCount; l++ {
					regionName := stripQuotes(regionsCon.Index(l).S("name").String())
					cidrsCount := len(regionsCon.Index(l).S("cidrs").Data().([]interface{}))
					cidrsCon := regionsCon.Index(l).S("cidrs")

					for m := 0; m < cidrsCount; m++ {
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
						match := re.FindStringSubmatch(vrfRef)
						vrfRefName := match[3]

						id := stripQuotes(cidrsCon.Index(m).S("ip").String())
						primary := stripQuotes(cidrsCon.Index(m).S("primary").String())

						name := schemaId + "_" + siteId + "_" + templateName + "_" + vrfRefName + "_" + regionName + "_" + id
						resource := terraformutils.NewResource(
							id,
							name,
							"mso_schema_site_vrf_region_cidr",
							"mso",
							map[string]string{
								"site_id":       siteId,
								"template_name": templateName,
								"schema_id":     schemaId,
								"vrf_name":      vrfRefName,
								"region_name":   regionName,
								"ip":            id,
								"primary":       primary,
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
	}
	return nil
}
