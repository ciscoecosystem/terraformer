package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const TriggerSchedulerClass = "trigSchedP"

type TriggerSchedulerGenerator struct {
	ACIService
}

func (a *TriggerSchedulerGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, TriggerSchedulerClass)

	TriggerSchedulerCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	TriggerSchedulerCount, err := strconv.Atoi(stripQuotes(TriggerSchedulerCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < TriggerSchedulerCount; i++ {
		TriggerSchedulerDN := TriggerSchedulerCont.S("imdata").Index(i).S(TriggerSchedulerClass, "attributes", "dn").String()
		resource := terraformutils.NewSimpleResource(
			stripQuotes(TriggerSchedulerDN),
			stripQuotes(TriggerSchedulerDN),
			"aci_trigger_scheduler",
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
	return nil
}
