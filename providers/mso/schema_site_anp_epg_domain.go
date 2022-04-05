package mso

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchemaSiteAnpEpgDomain struct {
	MSOService
}

func (a *SchemaSiteAnpEpgDomain) InitResources() error {
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
			anpsCount := len(sitesCon.Index(j).S("anps").Data().([]interface{}))
			anpsCon := sitesCon.Index(j).S("anps")

			for k := 0; k < anpsCount; k++ {
				anpRef := stripQuotes(anpsCon.Index(k).S("anpRef").String())
				epgsCount := len(anpsCon.Index(k).S("epgs").Data().([]interface{}))
				epgsCon := anpsCon.Index(k).S("epgs")

				for l := 0; l < epgsCount; l++ {
					epgRef := stripQuotes(epgsCon.Index(l).S("epgRef").String())
					domainAssociationsCount := len(epgsCon.Index(l).S("domainAssociations").Data().([]interface{}))
					domainAssociationsCon := epgsCon.Index(l).S("domainAssociations")

					for m := 0; m < domainAssociationsCount; m++ {
						anpRefSplit := strings.Split(anpRef, "/")
						anpRefName := anpRefSplit[6]
						epgRefSplit := strings.Split(epgRef, "/")
						epgRefName := epgRefSplit[8]

						domainAssociationsID := stripQuotes(domainAssociationsCon.Index(m).S("dn").String())

						var domainAssociationsName string
						var vmmDomainProfile string = "VMware"

						domainAssociationsType := stripQuotes(domainAssociationsCon.Index(m).S("domainType").String())

						if domainAssociationsType == "vmmDomain" {
							re := regexp.MustCompile("uni/vmmp-(.*)/dom-(.*)")
							match := re.FindStringSubmatch(domainAssociationsID)
							domainAssociationsName = match[2]
							vmmDomainProfile = match[1]
						} else if domainAssociationsType == "l3ExtDomain" {
							re := regexp.MustCompile("uni/l3dom-(.*)")
							match := re.FindStringSubmatch(domainAssociationsID)
							domainAssociationsName = match[1]
						} else if domainAssociationsType == "l2ExtDomain" {
							re := regexp.MustCompile("uni/l2dom-(.*)")
							match := re.FindStringSubmatch(domainAssociationsID)
							domainAssociationsName = match[1]
						} else if domainAssociationsType == "physicalDomain" {
							re := regexp.MustCompile("uni/phys-(.*)")
							match := re.FindStringSubmatch(domainAssociationsID)
							domainAssociationsName = match[1]
						} else {
							re := regexp.MustCompile("uni/fc-(.*)")
							match := re.FindStringSubmatch(domainAssociationsID)
							domainAssociationsName = match[1]
						}

						deployImmediacy := stripQuotes(domainAssociationsCon.Index(m).S("deployImmediacy").String())
						resolutionImmediacy := stripQuotes(domainAssociationsCon.Index(m).S("resolutionImmediacy").String())

						name := schemaId + "_" + templateName + "_" + siteId + "_" + anpRefName + "_" + epgRefName + "_" + domainAssociationsType + "_" + strconv.Itoa(rand.Intn(1000))
						resource := terraformutils.NewResource(
							domainAssociationsID,
							name,
							"mso_schema_site_anp_epg_domain",
							"mso",
							map[string]string{
								"site_id":              siteId,
								"template_name":        templateName,
								"schema_id":            schemaId,
								"anp_name":             anpRefName,
								"epg_name":             epgRefName,
								"domain_type":          domainAssociationsType,
								"dn":                   domainAssociationsName,
								"deploy_immediacy":     deployImmediacy,
								"resolution_immediacy": resolutionImmediacy,
								"vmm_domain_profile":   vmmDomainProfile,
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
	}
	return nil
}
