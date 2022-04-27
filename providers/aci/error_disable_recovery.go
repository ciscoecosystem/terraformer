package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const errorDisabledRecoveryPolicyClassName = "edrErrDisRecoverPol"

type ErrorDisabledRecoveryPolicyGenerator struct {
	ACIService
}

func (a *ErrorDisabledRecoveryPolicyGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, errorDisabledRecoveryPolicyClassName)

	ErrorDisabledRecoveryPolicyCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	ErrorDisabledRecoveryPolicyCount, err := strconv.Atoi(stripQuotes(ErrorDisabledRecoveryPolicyCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	eventMap, eventIds, err := getEventMap(client)
	if err != nil {
		return err
	}

	for i := 0; i < ErrorDisabledRecoveryPolicyCount; i++ {
		ErrorDisabledRecoveryPolicyAttr := ErrorDisabledRecoveryPolicyCont.S("imdata").Index(i).S(errorDisabledRecoveryPolicyClassName, "attributes")
		ErrorDisabledRecoveryPolicyDN := G(ErrorDisabledRecoveryPolicyAttr, "dn")
		if filterChildrenDn(ErrorDisabledRecoveryPolicyDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				ErrorDisabledRecoveryPolicyDN,
				resourceNamefromDn(errorDisabledRecoveryPolicyClassName, ErrorDisabledRecoveryPolicyDN, i),
				"aci_error_disable_recovery",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{
					"edr_event":     eventMap,
					"edr_event_ids": eventIds,
				},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}

func getEventMap(c *ACIClient) ([]map[string]string, []string, error) {
	cont, err := c.GetViaURL("api/node/class/edrEventP.json")
	if err != nil {
		return nil, nil, err
	}
	eventValues := make([]map[string]string, 0, 1)
	eventIds := make([]string, 0, 1)
	eventData := cont.S("imdata")
	for i := 0; i < len(eventData.Data().([]interface{})); i++ {
		eventCont := eventData.Index(i)
		eventContOut := eventCont.S("edrEventP")
		eventAttrCont := eventContOut.S("attributes")
		event := G(eventAttrCont, "event")
		name := G(eventAttrCont, "name")
		recover := G(eventAttrCont, "recover")
		id := G(eventAttrCont, "dn")
		eventValue := make(map[string]string, 0)
		eventValue["event"] = event
		eventValue["name"] = name
		eventValue["recover"] = recover
		eventValues = append(eventValues, eventValue)
		eventIds = append(eventIds, id)
	}
	return eventValues, eventIds, nil
}
