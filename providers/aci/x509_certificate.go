package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const x509CertificateClass = "aaaUserCert"

type X509CertificateGenerator struct {
	ACIService
}

func (a *X509CertificateGenerator) InitResources() error {

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, x509CertificateClass)

	x509CertificateCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	totalCount := stripQuotes(x509CertificateCont.S("totalCount").String())

	if totalCount == "{}" {
		totalCount = "0"
	}

	x509CertificateCount, err := strconv.Atoi(stripQuotes(x509CertificateCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < x509CertificateCount; i++ {
		x509CertificateProfileDN := stripQuotes(x509CertificateCont.S("imdata").Index(i).S(x509CertificateClass, "attributes", "dn").String())
		if filterChildrenDn(x509CertificateProfileDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				x509CertificateProfileDN,
				x509CertificateProfileDN,
				"aci_x509_certificate",
				"aci",
				[]string{
					"name_alias",
					"annotation",
					"description",
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}

	return nil
}
