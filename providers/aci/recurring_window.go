package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const recurringWindowClassName = "trigRecurrWindowP"

type RecurringWindowGenerator struct {
	ACIService
}

func (a *RecurringWindowGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, recurringWindowClassName)

	RecurringWindowCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	RecurringWindowCount, err := strconv.Atoi(stripQuotes(RecurringWindowCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < RecurringWindowCount; i++ {
		RecurringWindowAttr := RecurringWindowCont.S("imdata").Index(i).S(recurringWindowClassName, "attributes")
		RecurringWindowDN := G(RecurringWindowAttr, "dn")
		name := G(RecurringWindowAttr, "name")
		if filterChildrenDn(RecurringWindowDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RecurringWindowDN,
				resourceNamefromDn(recurringWindowClassName, RecurringWindowDN, i),
				"aci_recurring_window",
				"aci",
				map[string]string{
					"scheduler_dn": GetParentDn(RecurringWindowDN, fmt.Sprintf("/recurrwinp-%s", name)),
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
