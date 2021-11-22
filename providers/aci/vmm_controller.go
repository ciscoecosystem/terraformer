package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const VmmControllerClass = "vmmCtrlrP"

type VmmControllerGenerator struct {
	ACIService
}

func (a *VmmControllerGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, VmmControllerClass)

	VmmControllerCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	VmmControllerCount, err := strconv.Atoi(stripQuotes(VmmControllerCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < VmmControllerCount; i++ {
		VmmControllerDN := stripQuotes(VmmControllerCont.S("imdata").Index(i).S(VmmControllerClass, "attributes", "dn").String())
		if filterChildrenDn(VmmControllerDN, client.parentResource) != "" {
			resource := terraformutils.NewSimpleResource(
				VmmControllerDN,
				fmt.Sprintf("%s_%s_%d", VmmControllerClass, GetMOName(VmmControllerDN), i),
				"aci_vmm_controller",
				"aci",
				[]string{
					"dvs_version",
					"inventory_trig_st",
					"mode",
					"msft_config_err_msg",
					"n1kv_stats_mode",
					"port",
					"scope",
					"seq_num",
					"stats_mode",
					"vxlan_depl_pref",
					"relation_vmm_rs_acc",
					"relation_vmm_rs_ctrlr_p_mon_pol",
					"relation_vmm_rs_mcast_addr_ns",
					"relation_vmm_rs_mgmt_e_pg",
					"relation_vmm_rs_to_ext_dev_mgr",
					"relation_vmm_rs_vmm_ctrlr_p",
					"relation_vmm_rs_vxlan_ns",
					"relation_vmm_rs_vxlan_ns_def",
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
