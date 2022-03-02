package mso

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type LabelGenerator struct {
	MSOService
}

func (a *LabelGenerator) InitResources() error {
	mso := a.getClient().(*client.Client)
	con, err := mso.GetViaURL("api/v1/labels/")
	if err != nil {
		return err
	}
	for i := 0; i < len(con.S("labels").Data().([]interface{})); i++ {
		labelId := stripQuotes(con.S("labels").Index(i).S("id").String())
		resource := terraformutils.NewResource(
			labelId,
			strconv.Itoa(i),
			"mso_label",
			"mso",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
