package aci

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ACIProvider struct {
	terraformutils.Provider
	baseURL  string
	username string
	password string
	insecure bool
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
		"logical_node_to_fabric_node_profile": {
			"logical_node_profile": []string{"logical_node_profile_dn", "id"},
		},
		"l3out_static_route": {
			"logical_node_to_fabric_node_profile": []string{"logical_node_to_fabric_node_profile_dn", "id"},
		},
		"l3out_static_route_next_hop": {
			"l3out_static_route": []string{"l3out_static_route_dn", "id"},
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
			"l3out_path_attachment": []string{"l3out_path_attachment_dn", "id"},
		},
		"bgp_route_summarization": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_prefix": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_connectivity_profile": {
			"l3out_path_attachment": []string{"parent_dn", "id"},
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
			"endpoint_security_group": []string{"endpoint_security_group", "id"},
		},
		"endpoint_security_group": {
			"application_profile": []string{"application_profile", "id"},
		},
		"epg_to_contract": {
			"application_epg": []string{"application_epg_dn", "id"},
		},
		"epg_to_domain": {
			"application_epg": []string{"application_epg_dn", "id"},
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
	}
}

func (p ACIProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"aci": map[string]interface{}{
				"username": p.username,
				"password": p.password,
				"url":      p.baseURL,
			},
		},
	}
}

func (p *ACIProvider) Init(args []string) error {
	p.baseURL = args[0]
	p.username = args[1]
	p.password = args[2]
	p.insecure = true
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
		"username": p.username,
		"password": p.password,
		"base_url": p.baseURL,
		"insecure": p.insecure,
	})
	return nil
}

// GetSupportedService return map of support service for Github
func (p *ACIProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		// "members":               &MembersGenerator{},

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
		"taboo_contract":			    			&TabooContractGenerator{},
		"vmm_domain":				    			&VmmDomGenerator{},
		"miscabling_protocol_interface_policy":     &MiscablingProtocolInterfacePolicyGenerator{},
		"l2_interface_policy":                      &L2InterfacePolicyGenerator{},
		"port_security_policy":                     &PortSecurityPolicyGenerator{},
		"end_point_retention_policy":               &EndpointRetentionPolicyGenerator{},
		"vlan_pool":								&VlanPoolGenerator{},
		"ranges":									&RangesGenerator{},
		"physical_domain":							&PhysicalDomGenerator{},
		"l3_domain_profile":						&L3DomPGenerator{},

	}
}
