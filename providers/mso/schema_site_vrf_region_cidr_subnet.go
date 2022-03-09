package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteVrfRegionCidrSubnet struct {
	MSOService
}

func (a *SchemaSiteVrfRegionCidrSubnet) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := mso.GetViaURL("api/v1/schemas")
	if err != nil {
		return err
	}
	schemaLen := len(con.S("schemas").Data().([]interface{}))
	for i := 0; i < schemaLen; i++ {
		schemaCont := con.S("schemas").Index(i)
		siteLen := 0
		if schemaCont.Exists("sites") {
			siteLen = len(schemaCont.S("sites").Data().([]interface{}))
		}

		for j := 0; j < siteLen; j++ {
			siteCont := schemaCont.S("sites").Index(j)
			siteId := models.G(siteCont, "siteId")

			vrfsLen := 0
			if siteCont.Exists("vrfs") {
				vrfsLen = len(siteCont.S("vrfs").Data().([]interface{}))
			}

			for k := 0; k < vrfsLen; k++ {
				vrfCont := siteCont.S("vrfs").Index(k)
				vrfRef := models.G(vrfCont, "vrfRef")
				re := regexp.MustCompile("/schemas/(.*)/templates/(.*)/vrfs/(.*)")
				match := re.FindStringSubmatch(vrfRef)
				regionsLen := 0
				if vrfCont.Exists("regions") {
					regionsLen = len(vrfCont.S("regions").Data().([]interface{}))
				}

				for m := 0; m < regionsLen; m++ {
					regionCont := vrfCont.S("regions").Index(m)
					regionName := models.G(regionCont, "name")

					cidrsLen := 0
					if regionCont.Exists("cidrs") {
						cidrsLen = len(regionCont.S("cidrs").Data().([]interface{}))
					}

					for n := 0; n < cidrsLen; n++ {
						cidrCont := regionCont.S("cidrs").Index(n)

						cidrIP := models.G(cidrCont, "ip")

						subnetsLen := 0
						if cidrCont.Exists("subnets") {
							subnetsLen = len(cidrCont.S("subnets").Data().([]interface{}))
						}

						for l := 0; l < subnetsLen; l++ {
							subnetCont := cidrCont.S("subnets").Index(l)
							ip := models.G(subnetCont, "ip")
							zone := models.G(subnetCont, "zone")
							usage := models.G(subnetCont, "usage")

							resourceName := match[1] + "_" + siteId + "_" + match[3] + "_" + regionName + "_" + cidrIP + "_" + ip
							resource := terraformutils.NewResource(
								ip,
								resourceName,
								"mso_schema_site_vrf_region_cidr_subnet",
								"mso",
								map[string]string{
									"schema_id":     match[1],
									"template_name": match[2],
									"site_id":       siteId,
									"vrf_name":      match[3],
									"region_name":   regionName,
									"cidr_ip":       cidrIP,
									"ip":            ip,
									"zone":          zone,
									"usage":         usage,
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
	}
	return nil
}
