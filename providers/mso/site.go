package mso

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type SiteGenerator struct {
	MSOService
}

func (a *SiteGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	platform := mso.GetPlatform()
	var version string
	if platform == "nd" {
		version = "v2"
	} else {
		version = "v1"
	}
	path := fmt.Sprintf("api/%s/sites", version)
	con, err := mso.GetViaURL(path)
	if err != nil {
		return err
	}
	if version == "v2" {
		for i := 0; i < len(con.S("sites").Data().([]interface{})); i++ {
			labels := []string{}
			siteCont := con.S("sites").Index(i).S("common")
			name := stripQuotes(siteCont.S("name").String())
			apicSiteID := stripQuotes(siteCont.S("siteId").String())
			siteId := stripQuotes(con.S("sites").Index(i).S("id").String())
			if siteCont.Exists("labels") {
				labels = siteCont.S("labels").Data().([]string)
			}
			siteName := siteId + "_" + apicSiteID + "_" + name
			resource := terraformutils.NewResource(
				siteId,
				siteName,
				"mso_site",
				"mso",
				map[string]string{
					"name":         name,
					"apic_site_id": apicSiteID,
				},
				[]string{},
				map[string]interface{}{
					"labels": labels,
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
		return nil
	} else {
		for i := 0; i < len(con.S("sites").Data().([]interface{})); i++ {
			urls := []interface{}{}
			labels := []string{}
			location := map[string]interface{}{}
			siteCont := con.S("sites").Index(i)
			name := stripQuotes(siteCont.S("name").String())
			apicSiteID := stripQuotes(siteCont.S("apicSiteId").String())
			siteId := stripQuotes(siteCont.S("id").String())
			userName := stripQuotes(siteCont.S("username").String())
			platformName := stripQuotes(siteCont.S("platform").String())
			maintenanceMode := siteCont.S("maintenanceMode").Data().(bool)
			if siteCont.Exists("urls") {
				urls = siteCont.S("urls").Data().([]interface{})
			}
			if siteCont.Exists("labels") {
				labels = siteCont.S("labels").Data().([]string)
			}
			if siteCont.Exists("location") {
				loc1 := con.S("location").Data()
				if loc1 != nil {
					loc := loc1.(map[string]interface{})
					location["lat"] = fmt.Sprintf("%v", loc["lat"])
					location["long"] = fmt.Sprintf("%v", loc["long"])
				} else {
					location = nil
				}
			}
			siteName := siteId + "_" + apicSiteID + "_" + name
			resource := terraformutils.NewResource(
				siteId,
				siteName,
				"mso_site",
				"mso",
				map[string]string{
					"name":         name,
					"apic_site_id": apicSiteID,
					"username":     userName,
					"platform":     platformName,
				},
				[]string{},
				map[string]interface{}{
					"maintenance_mode": maintenanceMode,
					"urls":             urls,
					"location":         location,
					"labels":           labels,
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
		return nil
	}
}
