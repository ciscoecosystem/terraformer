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
		"bgp_route_summarization": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_prefix": {
			"tenant": []string{"tenant_dn", "id"},
		},
		"bgp_peer_connectivity_profile": {
			"logical_node_profile": []string{"parent_dn", "id"},
		},
		"ospf_route_summarization": {
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
		"tenant":                             &TenantGenerator{},
		"application_profile":                &ApplicationProfileGenerator{},
		"application_epg":                    &ApplicationEPGGenerator{},
		"bridge_domain":                      &BridgeDomainGenerator{},
		"contract":                           &ContractGenerator{},
		"contract_subject":                   &ContractSubjectGenerator{},
		"subnet":                             &SubnetGenerator{},
		"filter":                             &FilterGenerator{},
		"filter_entry":                       &FilterEntryGenerator{},
		"vpc_explicit_protection_group":      &VPCExplicitProtectionGroupGenerator{},
		"vrf":                                &VRFGenerator{},
		"l3_outside":                         &L3OutsideGenerator{},
		"ospf_interface_policy":              &ospfInterfacePolicyGenerator{},
		"logical_node_profile":               &LogicalNodeProfileGenerator{},
		"logical_interface_profile":          &LogicalInterfaceProfileGenerator{},
		"dhcp_option_policy":                 &DhcpOptionPolicyGenerator{},
		"external_network_instance_profile":  &ExtNetInsProGenerator{},
		"dhcp_relay_policy":                  &DHCPRelayPolicyGenerator{},
		"bd_dhcp_label":                      &BDDHCPLabelGenerator{},
		"l3_ext_subnet":                      &L3ExtSubnetGenerator{},
		"l3out_bgp_external_policy":          &L3OutBGPExtPolGenerator{},
		"l3out_loopback_interface_profile":   &L3OutLoopbackInterfaceProGenerator{},
		"l3out_ospf_external_policy":         &L3outOspfExternalPolicyGenerator{},
		"l3out_ospf_interface_profile":       &L3outOspfInterfaceProfileGenerator{},
		"l3out_path_attachment":              &L3outPathAttachmentGenerator{},
		"l3out_path_attachment_secondary_ip": &L3outPathAttachmentSecondaryIPGenerator{},
		"bgp_route_summarization":            &BGPRouteSumGenerator{},
		"bgp_peer_prefix":                    &BGPPeerPrefixGenerator{},
		"bgp_peer_connectivity_profile":      &BGPPeerConnectivityProGenerator{},
		"ospf_route_summarization":           &OSPFRouteSumGenerator{},
	}
}
