package mso

import (
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/models"
)

type SchemaSiteVrfRegion struct {
	MSOService
}

func (a *SchemaSiteVrfRegion) InitResources() error {
	mso, err := a.getClient()
	if err != nil {
		return err
	}
	con, err := getSchemaContainer(mso)
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
					var vpnGateway bool
					var hubNetworkEnable bool

					if regionCont.Exists("isVpnGatewayRouter") {
						vpnGateway = regionCont.S("isVpnGatewayRouter").Data().(bool)
					}
					if regionCont.Exists("isTGWAttachment") {
						hubNetworkEnable = regionCont.S("isTGWAttachment").Data().(bool)
					}
					cidrList := make([]interface{}, 0, 1)
					cidrs := regionCont.S("cidrs").Data().([]interface{})
					for _, tempCidr := range cidrs {
						cidr := tempCidr.(map[string]interface{})

						cidrMap := make(map[string]interface{})
						cidrMap["cidr_ip"] = cidr["ip"]
						cidrMap["primary"] = cidr["primary"]

						subnets := cidr["subnets"].([]interface{})
						subnetList := make([]interface{}, 0, 1)
						for _, tempSubnet := range subnets {
							subnet := tempSubnet.(map[string]interface{})

							subnetMap := make(map[string]interface{})
							subnetMap["ip"] = subnet["ip"]
							if subnet["zone"] != nil {
								subnetMap["zone"] = subnet["zone"]
							}
							if subnet["usage"] != nil {
								subnetMap["usage"] = subnet["usage"]
							}

							subnetList = append(subnetList, subnetMap)
						}
						cidrMap["subnet"] = subnetList

						cidrList = append(cidrList, cidrMap)
					}
					resourceName := match[1] + "_" + siteId + "_" + match[3] + "_" + regionName
					resource := terraformutils.NewResource(
						regionName,
						resourceName,
						"mso_schema_site_vrf_region",
						"mso",
						map[string]string{
							"schema_id":     match[1],
							"template_name": match[2],
							"site_id":       siteId,
							"vrf_name":      match[3],
							"region_name":   regionName,
						},
						[]string{},
						map[string]interface{}{
							"vpn_gateway":        vpnGateway,
							"hub_network_enable": hubNetworkEnable,
							"cidr":               cidrList,
						},
					)
					resource.SlowQueryRequired = SlowQueryRequired
					a.Resources = append(a.Resources, resource)
				}
			}
		}
	}
	return nil
}
