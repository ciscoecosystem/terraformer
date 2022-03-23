package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteVrfRegionHubNetworkGenerator struct {
	MSOService
}

func (a *SchemaSiteVrfRegionHubNetworkGenerator) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
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
					if regionsCon.Index(l).Exists("cloudRsCtxProfileToGatewayRouterP") {
						hubNetCon := regionsCon.Index(l).S("cloudRsCtxProfileToGatewayRouterP")
						hubNetMap := hubNetCon.Data().(map[string]interface{})
						re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
						match := re.FindStringSubmatch(vrfRef)
						vrfRefName := match[3]
						tenantName := hubNetMap["tenantName"].(string)
						nName := hubNetMap["name"].(string)
						id := schemaId + "/site/" + siteId + "/template/" + templateName + "/vrf/" + vrfRefName + "/region/" + regionName + "/tenant/" + tenantName + "/" + nName
						name := schemaId + "_" + siteId + "_" + templateName + "_" + vrfRefName + "_" + regionName + "_" + id
						resource := terraformutils.NewResource(
							id,
							name,
							"mso_schema_site_vrf_region_hub_network",
							"mso",
							map[string]string{},
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
