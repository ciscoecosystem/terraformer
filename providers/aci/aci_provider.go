package aci

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ACIProvider struct {
	terraformutils.Provider
	baseURL        string
	username       string
	password       string
	parentResource string
	certName       string
	privateKey     string
	insecure       bool
}

func (p ACIProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"application_profile": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"application_epg": {
			"application_profile": []string{"application_profile_dn", "id"},
		},
		"bridge_domain": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"contract": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"contract_subject": {
			"contract": []string{"contract_dn", "id"},
		},
		"filter_entry": {
			"filter": []string{"filter_dn", "id"},
		},
		"filter": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"subnet": {
			"bridge_domain": []string{"parent_dn", "id"},
		},
		"vrf": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"l3_outside": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"ospf_interface_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"logical_node_profile": {
			"l3_outside": []string{"l3_outside_dn", "id"},
		},
		"logical_interface_profile": {
			"logical_node_profile": []string{"logical_node_profile_dn", "id"},
		},
		"logical_node_to_fabric_node": {
			"logical_node_profile": []string{"logical_node_profile_dn", "id"},
		},
		"l3out_static_route": {
			"logical_node_to_fabric_node": []string{"fabric_node_dn", "id"},
		},
		"l3out_static_route_next_hop": {
			"l3out_static_route": []string{"static_route_dn", "id"},
		},
		"dhcp_option_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"external_network_instance_profile": {
			"l3_outside": []string{"l3_outside_dn", "id"},
		},
		"dhcp_relay_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bd_dhcp_label": {
			"bridge_domain": []string{"bridge_domain_dn", "id"},
		},
		"l3_ext_subnet": {
			"external_network_instance_profile": []string{"external_network_instance_profile_dn", "id"},
		},
		"l3out_bgp_external_policy": {
			"l3_outside": []string{"l3_outside_dn", "id"},
		},
		"l3out_ospf_external_policy": {
			"l3_outside": []string{"l3_outside_dn", "id"},
		},
		"l3out_ospf_interface_profile": {
			"logical_interface_profile": []string{"logical_interface_profile_dn", "id"},
		},
		"l3out_path_attachment": {
			"logical_interface_profile": []string{"logical_interface_profile_dn", "id"},
		},
		"l3out_path_attachment_secondary_ip": {
			"l3out_path_attachment": []string{"l3out_path_attachment_dn", "id"},
		},
		"l3out_vpc_member": {
			"l3out_path_attachment": []string{"leaf_port_dn", "id"},
		},
		"bgp_route_summarization": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_prefix": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_connectivity_profile": {
			"l3out_path_attachment": []string{"parent_dn", "id"},
			"logical_node_profile":  []string{"parent_dn", "id"},
		},
		"ospf_route_summarization": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_timers": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_address_family_context": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_best_path_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_route_control_profile": {
			"l3_outside": []string{"parent_dn", "id"},
			"tenant":     []string{"parent_dn", "id"},
		},
		"ospf_timers": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"l3out_route_tag_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"hsrp_interface_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"l3out_bfd_interface_profile": {
			"logical_interface_profile": []string{"logical_interface_profile_dn", "id"},
		},
		"hsrp_group_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"l3out_floating_svi": {
			"logical_interface_profile": []string{"logical_interface_profile_dn", "id"},
		},
		"l3out_bgp_protocol_profile": {
			"logical_node_profile": []string{"logical_node_profile_dn", "id"},
		},
		"endpoint_security_group_selector": {
			"endpoint_security_group": []string{"endpoint_security_group_dn", "id"},
		},
		"endpoint_security_group": {
			"application_profile": []string{"application_profile_dn", "id"},
		},
		"epg_to_contract": {
			"application_epg": []string{"application_epg_dn", "id"},
			"contract":        []string{"contract_dn", "id"},
		},
		"epg_to_domain": {
			"application_epg": []string{"application_epg_dn", "id"},
			"fc_domain":       []string{"tdn", "id"},
		},
		"epg_to_static_path": {
			"application_epg": []string{"application_epg_dn", "id"},
		},
		"imported_contract": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"any": {
			"vrf": []string{"vrf_dn", "id"},
		},
		"cloud_applicationcontainer": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"cloud_aws_provider": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"cloud_context_profile": {
			"tenant": []string{"tenant_dn", "id"},
			"vrf":    []string{"vrf_dn", "id"},
		},
		"cloud_epg": {
			"cloud_applicationcontainer": []string{"cloud_applicationcontainer_dn", "id"},
		},
		"cloud_subnet": {
			"cloud_cidr_pool": []string{"cloud_cidr_pool_dn", "id"},
		},
		"cloud_cidr_pool": {
			"cloud_context_profile": []string{"cloud_context_profile_dn", "id"},
		},
		"taboo_contract": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"l2_outside": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"static_node_mgmt_address": {
			"node_mgmt_epg": []string{"management_epg_dn", "id"},
		},
		"span_sourcedestination_group_match_label": {
			"span_source_group": []string{"span_source_group_dn", "id"},
		},
		"leaf_selector": {
			"leaf_profile": []string{"leaf_profile_dn", "id"},
		},
		"node_block": {
			"leaf_selector": []string{"switch_association_dn", "id"},
		},
		"access_port_selector": {
			"leaf_interface_profile": []string{"leaf_interface_profile_dn", "id"},
		},
		"access_port_block": {
			"access_port_selector": []string{"access_port_selector_dn", "id"},
		},
		"access_sub_port_block": {
			"access_port_selector": []string{"access_port_selector_dn", "id"},
		},
		"access_generic": {
			"attachable_access_entity_profile": []string{"attachable_access_entity_profile_dn", "id"},
		},
		"access_group": {
			"access_port_selector": []string{"access_port_selector_dn", "id"},
		},
		"spine_switch_association": {
			"spine_profile": []string{"spine_profile_dn", "id"},
		},
		"l2out_extepg": {
			"l2_outside": []string{"l2_outside_dn", "id"},
		},
		"x509_certificate": {
			"local_user": []string{"local_user_dn", "id"},
		},
		"monitoring_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"action_rule_profile": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"fex_bundle_group": {
			"fex_profile": []string{"fex_profile_dn", "id"},
		},
		"end_point_retention_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"ranges": {
			"vlan_pool": []string{"vlan_pool_dn", "id"},
		},
		"vlan_encapsulationfor_vxlan_traffic": {
			"attachable_access_entity_profile": []string{"attachable_access_entity_profile_dn", "id"},
		},
		"l4_l7_service_graph_template": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"spine_port_selector": {
			"spine_profile": []string{"spine_profile_dn", "id"},
		},
		"epgs_using_function": {
			"access_generic": []string{"access_generic_dn", "id"},
		},
		"service_redirect_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"destination_of_redirected_traffic": {
			"service_redirect_policy": []string{"service_redirect_policy_dn", "id"},
		},
		"span_destination_group": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"span_source_group": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"logical_device_context": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"function_node": {
			"l4_l7_service_graph_template": []string{"l4_l7_service_graph_template_dn", "id"},
		},
		"connection": {
			"l4_l7_service_graph_template": []string{"l4_l7_service_graph_template_dn", "id"},
		},
		"logical_interface_context": {
			"logical_device_context": []string{"logical_device_context_dn", "id"},
		},
		"l3out_hsrp_interface_profile": {
			"logical_interface_profile": []string{"logical_interface_profile_dn", "id"},
		},
		"l3out_hsrp_interface_group": {
			"l3out_hsrp_interface_profile": []string{"l3out_hsrp_interface_profile_dn", "id"},
		},
		"vmm_credential": {
			"vmm_domain": []string{"vmm_domain_dn", "id"},
		},
		"vmm_controller": {
			"vmm_domain": []string{"vmm_domain_dn", "id"},
		},
		"cloud_vpn_gateway": {
			"cloud_context_profile": []string{"cloud_context_profile_dn", "id"},
		},
		"cloud_external_epg": {
			"cloud_applicationcontainer": []string{"cloud_applicationcontainer_dn", "id"},
		},
		"cloud_endpoint_selector": {
			"cloud_epg": []string{"cloud_epg_dn", "id"},
		},
		"cloud_endpoint_selectorfor_external_epgs": {
			"cloud_external_epg": []string{"cloud_external_epg_dn", "id"},
		},
		"node_block_firmware": {
			"firmware_group": []string{"firmware_group_dn", "id"},
		},
		"maintenance_group_node": {
			"pod_maintenance_group": []string{"pod_maintenance_group_dn", "id"},
			"firmware_group":        []string{"pod_maintenance_group_dn", "id"},
		},
		"vswitch_policy": {
			"vmm_domain": []string{"vmm_domain_dn", "id"},
		},
		"l3out_hsrp_secondary_vip": {
			"l3out_hsrp_interface_group": []string{"l3out_hsrp_interface_group_dn", "id"},
		},
		"vrf_snmp_context": {
			"vrf": []string{"vrf_dn", "id"},
		},
		"vrf_snmp_context_community": {
			"vrf_snmp_context": []string{"vrf_snmp_context_dn", "id"},
		},
		"user_security_domain": {
			"local_user": []string{"local_user_dn", "id"},
		},
		"bfd_interface_policy": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"user_security_domain_role": {
			"user_security_domain": []string{"user_domain_dn", "id"},
		},
		"ldap_group_map_rule_to_group_map": {
			"ldap_group_map": []string{"ldap_group_map_dn", "id"},
		},
		"tacacs_accounting_destination": {
			"tacacs_accounting": []string{"tacacs_accounting_dn", "id"},
		},
		"vrf_to_bgp_address_family_context": {
			"vrf":                        []string{"vrf_dn", "id"},
			"bgp_address_family_context": []string{"bgp_address_family_context_dn", "id"},
		},
		"aci_mgmt_zone": {
			"aci_managed_node_connectivity_group": []string{"aci_managed_node_connectivity_group_dn", "id"},
		},
		"aci_recurring_window": {
			"trigger_schedular": []string{"trigger_schedular_dn", "id"},
		},
		"match_rule": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"route_control_profile": {
			"l3_outside": []string{"parent_dn", "id"},
			"tenant":     []string{"parent_dn", "id"},
		},
		"route_control_context": {
			"route_control_profile": []string{"route_control_profile_dn", "id"},
			"set_rule":              []string{"action_rule_profile_dn", "id"},
		},
		"action_rule_additional_communities": {
			"action_rule_profile": []string{"action_rule_profile_dn", "id"},
		},
	}
}

func (p ACIProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"aci": map[string]interface{}{
				"username":    p.username,
				"password":    p.password,
				"cert_name":   p.certName,
				"private_key": p.privateKey,
				"url":         p.baseURL,
			},
		},
	}
}

func (p *ACIProvider) Init(args []string) error {
	p.baseURL = args[0]
	p.username = args[1]
	p.password = args[2]
	p.certName = args[3]
	p.privateKey = args[4]
	p.parentResource = args[5]
	p.insecure = true
	os.Setenv("ACI_URL", p.baseURL)
	os.Setenv("ACI_USERNAME", p.username)
	os.Setenv("ACI_PASSWORD", p.password)
	os.Setenv("ACI_CERT_NAME", p.certName)
	os.Setenv("ACI_PRIVATE_KEY", p.privateKey)
	return nil
}

func (p *ACIProvider) GetName() string {
	return "aci"
}

func (p *ACIProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username":        p.username,
		"password":        p.password,
		"base_url":        p.baseURL,
		"cert_name":       p.certName,
		"private_key":     p.privateKey,
		"insecure":        p.insecure,
		"parent_resource": p.parentResource,
	})
	return nil
}

func (p *ACIProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"tenant":                                   &TenantGenerator{},
		"application_profile":                      &ApplicationProfileGenerator{},
		"application_epg":                          &ApplicationEPGGenerator{},
		"bridge_domain":                            &BridgeDomainGenerator{},
		"contract":                                 &ContractGenerator{},
		"contract_subject":                         &ContractSubjectGenerator{},
		"subnet":                                   &SubnetGenerator{},
		"filter":                                   &FilterGenerator{},
		"filter_entry":                             &FilterEntryGenerator{},
		"vpc_explicit_protection_group":            &VPCExplicitProtectionGroupGenerator{},
		"vrf":                                      &VRFGenerator{},
		"l3_outside":                               &L3OutsideGenerator{},
		"ospf_interface_policy":                    &ospfInterfacePolicyGenerator{},
		"logical_node_profile":                     &LogicalNodeProfileGenerator{},
		"logical_interface_profile":                &LogicalInterfaceProfileGenerator{},
		"dhcp_option_policy":                       &DhcpOptionPolicyGenerator{},
		"external_network_instance_profile":        &ExtNetInsProGenerator{},
		"dhcp_relay_policy":                        &DHCPRelayPolicyGenerator{},
		"bd_dhcp_label":                            &BDDHCPLabelGenerator{},
		"l3_ext_subnet":                            &L3ExtSubnetGenerator{},
		"l3out_bgp_external_policy":                &L3OutBGPExtPolGenerator{},
		"l3out_loopback_interface_profile":         &L3OutLoopbackInterfaceProGenerator{},
		"l3out_ospf_external_policy":               &L3outOspfExternalPolicyGenerator{},
		"l3out_ospf_interface_profile":             &L3outOspfInterfaceProfileGenerator{},
		"l3out_path_attachment":                    &L3outPathAttachmentGenerator{},
		"l3out_path_attachment_secondary_ip":       &L3outPathAttachmentSecondaryIPGenerator{},
		"bgp_route_summarization":                  &BGPRouteSumGenerator{},
		"bgp_peer_prefix":                          &BGPPeerPrefixGenerator{},
		"bgp_peer_connectivity_profile":            &BGPPeerConnectivityProGenerator{},
		"ospf_route_summarization":                 &OSPFRouteSumGenerator{},
		"bgp_timers":                               &BgpTimersGenerator{},
		"bgp_address_family_context":               &BgpAddressFamilyContextGenerator{},
		"bgp_best_path_policy":                     &BgpBestPathPolicyGenerator{},
		"bgp_route_control_profile":                &BgpRouteControlProfileGenerator{},
		"ospf_timers":                              &OSPFTimersGenerator{},
		"l3out_route_tag_policy":                   &L3OutRouteTagPolicyGenerator{},
		"hsrp_interface_policy":                    &HSRPInterfacePolicyGenerator{},
		"l3out_bfd_interface_profile":              &L3OutBFDInterfaceProfileGenerator{},
		"hsrp_group_policy":                        &HSRPGroupPolicyGenerator{},
		"l3out_floating_svi":                       &L3OutFloatingSviGenerator{},
		"l3out_hsrp_secondary_vip":                 &L3OutHSRPSecondaryVipGenerator{},
		"l3out_bgp_protocol_profile":               &L3OutBGPProtocolProfileGenerator{},
		"any":                                      &AnyGenerator{},
		"endpoint_security_group_selector":         &ApplicationEndpointSecurityGroupSelectorGenerator{},
		"epg_to_static_path":                       &EPGToStaticPathGenerator{},
		"epg_to_contract":                          &EPGToContractGenerator{},
		"epg_to_domain":                            &EPGToDomainGenerator{},
		"endpoint_security_group":                  &ApplicationEndpointSecurityGroupGenerator{},
		"logical_node_to_fabric_node":              &LogicalNodeToFabricNodeGenerator{},
		"cloud_applicationcontainer":               &CloudApplicationContainerGenerator{},
		"cloud_cidr_pool":                          &CloudCidrPoolGenerator{},
		"l3out_static_route":                       &L3OutStaticRouteGenerator{},
		"l3out_static_route_next_hop":              &L3OutStaticRouteNextHopGenerator{},
		"l3out_vpc_member":                         &L3OutVPCMemberGenerator{},
		"cloud_endpoint_selectorfor_external_epgs": &CloudEndpointSelectorForExternalEpgsGenerator{},
		"cloud_endpoint_selector":                  &CloudEndpointSelectorGenerator{},
		"cloud_external_epg":                       &CloudExternalEPGGenerator{},
		"cloud_vpn_gateway":                        &CloudVPNGatewayGenerator{},
		"vmm_controller":                           &VmmControllerGenerator{},
		"vmm_credential":                           &VmmCredentialGenerator{},
		"vswitch_policy":                           &VswitchPolicyGenerator{},
		"cloud_domain_profile":                     &CloudDomainPGenerator{},
		"cloud_context_profile":                    &CloudContextPGenerator{},
		"cloud_epg":                                &CloudEPGGenerator{},
		"cloud_aws_provider":                       &CloudAWSProviderGenerator{},
		"imported_contract":                        &ImportedContractGenerator{},
		"l3out_hsrp_interface_group":               &L3OutHSRPInterfaceGroupGenerator{},
		"l3out_hsrp_interface_profile":             &L3OutHSRPInterfaceProfileGenerator{},
		"attachable_access_entity_profile":         &AttachableAccessEntityProfileGenerator{},
		"epgs_using_function":                      &EPGUsingFunctionGenerator{},
		"leaf_interface_profile":                   &LeafInterfaceProfileGenerator{},
		"cloud_subnet":                             &CloudSubnetGenerator{},
		"lldp_interface_policy":                    &LLDPInterfacePolicyGenerator{},
		"lacp_policy":                              &LacpPolicyGenerator{},
		"cdp_interface_policy":                     &CDPInterfacePolicyGenerator{},
		"vlan_encapsulationfor_vxlan_traffic":      &VlanVxlanTrafficGenerator{},
		"taboo_contract":                           &TabooContractGenerator{},
		"vmm_domain":                               &VmmDomGenerator{},
		"miscabling_protocol_interface_policy":     &MiscablingProtocolInterfacePolicyGenerator{},
		"l2_interface_policy":                      &L2InterfacePolicyGenerator{},
		"port_security_policy":                     &PortSecurityPolicyGenerator{},
		"end_point_retention_policy":               &EndpointRetentionPolicyGenerator{},
		"vlan_pool":                                &VlanPoolGenerator{},
		"ranges":                                   &RangesGenerator{},
		"physical_domain":                          &PhysicalDomGenerator{},
		"l3_domain_profile":                        &L3DomPGenerator{},
		"spine_port_selector":                      &SpinePortSelectorGenerator{},
		"spine_interface_profile":                  &SpineInterfaceProfileGenerator{},
		"spine_port_policy_group":                  &SpinePortPolicyGroupGenerator{},
		"fabric_if_pol":                            &FabricIfPolGenerator{},
		"l2_outside":                               &L2OutsideGenerator{},
		"node_mgmt_epg":                            &NodeMgmtEPGGenerator{},
		"static_node_mgmt_address":                 &StaticNodeMgmtAddressGenerator{},
		"local_user":                               &LocalUserGenerator{},
		"trigger_scheduler":                        &TriggerSchedulerGenerator{},
		"span_destination_group":                   &SpanDestinationGroupGenerator{},
		"span_source_group":                        &SpanSourceGroupGenerator{},
		"span_sourcedestination_group_match_label": &SpanSourceDestGroupMatchGenerator{},
		"maintenance_policy":                       &MaintenancePolicyGenerator{},
		"maintenance_group_node":                   &MaintenanceGroupNodeGenerator{},
		"node_block_firmware":                      &NodeBlockFirmWareGenerator{},
		"configuration_export_policy":              &ConfigExportPolicyGenerator{},
		"leaf_selector":                            &LeafSelectorGenerator{},
		"node_block":                               &NodeBlockGenerator{},
		"leaf_access_bundle_policy_group":          &LeafAccBunPolGGenerator{},
		"leaf_access_port_policy_group":            &LeafAccPorPolGGenerator{},
		"access_port_selector":                     &AccessPortSelectorGenerator{},
		"access_port_block":                        &AccessPortBlkGenerator{},
		"access_sub_port_block":                    &AccessSubPortBlkGenerator{},
		"spanning_tree_interface_policy":           &SpanningTreeInterfacePolicyGenerator{},
		"access_generic":                           &AccessGenericGenerator{},
		"access_group":                             &AccessGroupGenerator{},
		"spine_profile":                            &SpinePGenerator{},
		"spine_switch_association":                 &SpineSwitchAssGenerator{},
		"leaf_breakout_port_group":                 &LeafBreakoutPortGrpGenerator{},
		"vxlan_pool":                               &VxlanPoolGenerator{},
		"l2_domain":                                &L2DomGenerator{},
		"l2out_extepg":                             &L2OutExtEPGGenerator{},
		"aaa_domain":                               &AaaDomGenerator{},
		"x509_certificate":                         &X509CertificateGenerator{},
		"monitoring_policy":                        &MonPolGenerator{},
		"action_rule_profile":                      &ActionRuleProfGenerator{},
		"configuration_import_policy":              &ConfigImportPolicyGenerator{},
		"fabric_node_member":                       &FabricNodeMemberGenerator{},
		"fex_profile":                              &FexProfGenerator{},
		"fex_bundle_group":                         &FexBundleGrpGenerator{},
		"l4_l7_service_graph_template":             &L4L7ServiceGraphTemplateGenerator{},
		"service_redirect_policy":                  &ServiceRedirectPolicyGenerator{},
		"destination_of_redirected_traffic":        &DestinationOfRedirectedTrafficGenerator{},
		"logical_device_context":                   &LogicalDeviceContextGenerator{},
		"interface_fc_policy":                      &InterfaceFCPolicyGenerator{},
		"firmware_group":                           &FirmwareGroupGenerator{},
		"firmware_policy":                          &FirmwarePolicyGenerator{},
		"firmware_download_task":                   &FirmwareDownloadTaskGenerator{},
		"function_node":                            &FunctionNodeGenerator{},
		"connection":                               &ConnectionGenerator{},
		"logical_interface_context":                &LogicalInterfaceContextGenerator{},
		"vsan_pool":                                &VSANPoolGenerator{},
		"fc_domain":                                &FCDomainGenerator{},
		"leaf_profile":                             &LeafProfileGenerator{},
		"pod_maintenance_group":                    &PodMaintenanceGroupGenerator{},
		"l3_interface_policy":                      &L3InterfacePolicyGenerator{},
		"access_switch_policy_group":               &AccessSwitchPolicyGroupGenerator{},
		"mgmt_preference":                          &MgmtconnectivitypreferenceGenerator{},
		"fabric_node_control":                      &FabricNodeControlGenerator{},
		"vrf_snmp_context":                         &SNMPContextProfileGenerator{},
		"vrf_snmp_context_community":               &SNMPCommunityGenerator{},
		"fabric_wide_settings":                     &fabricWideSettingsPolicyGenerator{},
		"encryption_key":                           &AESEncryptionPassphraseandKeysforConfigExportandImportGenerator{},
		"coop_policy":                              &COOPGroupPolicyGenerator{},
		"port_tracking":                            &PortTrackingGenerator{},
		"user_security_domain":                     &UserDomainGenerator{},
		"error_disable_recovery":                   &ErrorDisabledRecoveryPolicyGenerator{},
		"bfd_interface_policy":                     &BFDInterfacePolicyGenerator{},
		"managed_node_connectivity_group":          &ManagedNodeConnectivityGroupGenerator{},
		"spine_switch_policy_group":                &SpineSwitchPolicyGroupGenerator{},
		"duo_provider_group":                       &DuoProviderGroupGenerator{},
		"console_authentication":                   &ConsoleAuthenticationMethodGenerator{},
		"ldap_provider":                            &LDAPProviderGenerator{},
		"tacacs_accounting":                        &TACACSMonitoringDestinationGroupGenerator{},
		"rsa_provider":                             &RSAProviderGenerator{},
		"saml_provider_group":                      &SAMLProviderGroupGenerator{},
		"user_security_domain_role":                &UserRoleGenerator{},
		"mcp_instance_policy":                      &MiscablingProtocolInstancePolicyGenerator{},
		"qos_instance_policy":                      &QOSInstancePolicyGenerator{},
		"ldap_group_map":                           &LDAPGroupMapGenerator{},
		"ldap_group_map_rule_to_group_map":         &LDAPGroupMaprulerefGenerator{},
		"vpc_domain_policy":                        &VPCDomainPolicyGenerator{},
		"mgmt_zone":                                &OOBManagedNodesZoneGenerator{},
		"recurring_window":                         &RecurringWindowGenerator{},
		"file_remote_path":                         &RemotePathofaFileGenerator{},
		"radius_provider_group":                    &RadiusProviderGroupGenerator{},
		"saml_provider":                            &SAMLProviderGenerator{},
		"tacacs_accounting_destination":            &TACACSDestinationGenerator{},
		"vrf_to_bgp_address_family_context":        &BGPAddressFamilyContextPolicyGenerator{},
		"match_rule":                               &MatchRuleGenerator{},
		"annotation":                               &TagGenerator{},
		"route_control_profile":                    &RouteControlProfileGenerator{},
		"route_control_context":                    &RouteControlContextGenerator{},
		"action_rule_additional_communities":       &RtctrlSetAddCommGenerator{},
		"endpoint_loop_protection":                 &EPLoopProtectionPolicyGenerator{},
		"endpoint_controls":                        &EndpointControlPolicyGenerator{},
		"endpoint_ip_aging_profile":                &IPAgingPolicyGenerator{},
	}
}
