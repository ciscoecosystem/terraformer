package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const qOSInstancePolicyClassName = "qosInstPol"

type QOSInstancePolicyGenerator struct {
	ACIService
}

func (a *QOSInstancePolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, qOSInstancePolicyClassName)

	QOSInstancePolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	QOSInstancePolicyCount, err := strconv.Atoi(stripQuotes(QOSInstancePolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < QOSInstancePolicyCount; i++ {
		QOSInstancePolicyAttr := QOSInstancePolicyCont.S("imdata").Index(i).S(qOSInstancePolicyClassName, "attributes")
		QOSInstancePolicyDN := G(QOSInstancePolicyAttr, "dn")
		if filterChildrenDn(QOSInstancePolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				QOSInstancePolicyDN,
				resourceNamefromDn(qOSInstancePolicyClassName, QOSInstancePolicyDN, i),
				"aci_qos_instance_policy",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
