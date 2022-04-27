
### Use with ACI

* use `--username`, `--password`, `--private-key`, `--cert-name` and `--base-url` options to specify server credentials for the first run. It should set respective environment variables of server credentials. Once environment variables are set there is no need to specify these options in the import command.

```
terraformer import aci --resources=tenant,application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
OR
terraformer import aci --resources=tenant --base-url=Cisco_APIC_url --username=Cisco_APIC_username --private-key="private.key" --cert-name="user_cert"

terraformer import aci --resources=* --excludes=application_profile
terraformer import aci --resources=tenant --filter=tenant=tenant_dn1:tenant_dn2
```
for more information regarding all supported flags use `--help` or `-h` flag.
```
terraformer import aci --help
```
#### Supported services

* `tenant`
    * `action_rule_profile`
    * `monitoring_policy`
    * `cloud_applicationcontainer`
        * `cloud_epg`
            * `cloud_endpoint_selector`
        * `cloud_external_epg`
            * `cloud_endpoint_selectorfor_external_epgs`
    * `application_profile`
        * `endpoint_security_group`
            * `endpoint_security_group_selector`
        * `application_epg`
            * `epg_to_contract`
            * `epg_to_domain`
            * `epg_to_static_path`
    * `vrf`
        * `any`
        * `vrf_snmp_context`
            * `vrf_snmp_context_community`
    * `bgp_route_control_profile`
    * `cloud_context_profile`
         * `cloud_cidr_pool`
            * `cloud_subnet`
         * `cloud_vpn_gateway`
    * `l3_outside`
        * `l3out_ospf_external_policy`
        * `external_network_instance_profile`
            * `l3_ext_subnet`
        * `logical_node_profile`
            * `bgp_peer_connectivity_profile`
            * `logical_node_to_fabric_node`
                * `l3out_static_route`
                    * `l3out_static_route_next_hop`
            * `l3out_bgp_protocol_profile`
            * `logical_interface_profile`
                * `l3out_ospf_interface_profile`
                * `l3out_floating_svi`
                * `l3out_path_attachment`
                    * `l3out_path_attachment_secondary_ip`
                    * `bgp_peer_connectivity_profile`
                    * `l3out_vpc_member`
                * `l3out_bfd_interface_profile`
                * `l3out_hsrp_interface_profile`
                    * `l3out_hsrp_interface_group`
                      * `l3out_hsrp_secondary_vip`
            * `logical_node_to_fabric_node`
              * `l3out_loopback_interface_profile`
        * `l3out_bgp_external_policy`
        * `bgp_route_control_profile` 
    * `bridge_domain` 
        * `subnet`
        * `bd_dhcp_label`
    * `contract`
        * `contract_subject`
        * `epg_to_contract`
    * `filter`
        * `filter_entry`
    * `ospf_interface_policy`
    * `dhcp_option_policy`
    * `dhcp_relay_policy`
    * `bgp_route_summarization`
    * `bgp_peer_prefix`
    * `ospf_route_summarization`
    * `bgp_address_family_context`
    * `bgp_best_path_policy`
    * `bgp_timers`
    * `ospf_timers`
    * `l3out_route_tag_policy`
    * `hsrp_interface_policy`
    * `hsrp_group_policy`
    * `imported_contract`
    * `cloud_aws_provider`
    * `taboo_contract`
    * `l2_outside`
        * `l2out_extepg`
    * `service_redirect_policy`
        * `destination_of_redirected_traffic`
    * `logical_device_context`
        * `logical_interface_context`
    * `l4_l7_service_graph_template`
        * `function_node`
        * `connection`
    * `span_destination_group`
    * `span_source_group`
        * `span_sourcedestination_group_match_label`
    * `end_point_retention_policy`
* `vpc_explicit_protection_group`
* `l3_domain_profile`
* `vmm_domain`
   * `vmm_controller`
   * `vmm_credential`
   * `vswitch_policy`
* `cloud_domain_profile`
* `attachable_access_entity_profile`
    * `access_generic`
        * `epgs_using_function`
    * `vlan_encapsulationfor_vxlan_traffic`
* `leaf_interface_profile`
    * `access_port_selector`
        * `access_group`
        * `access_port_block`
        * `access_sub_port_block`
* `leaf_profile`
    * `leaf_selector`
        * `node_block`
* `leaf_access_bundle_policy_group`
* `leaf_access_port_policy_group`
* `leaf_breakout_port_group`       
* `lldp_interface_policy`
* `lacp_policy`
* `cdp_interface_policy`
* `vxlan_pool`
* `vlan_pool`
    * `ranges`
* `physical_domain`    
* `miscabling_protocol_interface_policy`
* `l2_interface_policy`
* `l2_domain`
* `port_security_policy`
* `spine_profile`
    * `spine_switch_association`
    * `spine_port_selector`
* `spine_interface_profile`
* `spine_port_policy_group`
* `fabric_if_pol`
* `node_mgmt_epg`
    * `static_node_mgmt_address`
* `local_user`
    * `x509_certificate`
* `trigger_scheduler`
* `spanning_tree_interface_policy`
* `maintenance_policy`
* `configuration_export_policy`
* `aaa_domain`
* `configuration_import_policy`
* `fabric_node_member`
* `fex_profile`
    * `fex_bundle_group`
* `interface_fc_policy`
* `firmware_policy`
* `firmware_group`
    * `node_block_firmware`
* `firmware_download_task`
* `vsan_pool`
* `fc_domain`
* `pod_maintenance_group`
    * `maintenance_group_node`
* `access_switch_policy_group`
* `l3_interface_policy`
* `mgmt_preference`
* `fabric_node_control`
* `fabric_wide_settings`
* `encryption_key`
#### Attribute filters

Attribute filters allow filtering across different resource types by their attributes.

```
terraformer import aci --resources=tenant,application_profile --filter="Name=name;Value=val1:val2" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
Will import tenants and application_profiles having attribute `name` which has value either `val1`, `val2` or `val3`(supports any number of values saperated by ':'). Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applied by providing `Type=<type>` parameter. For example:
```
terraformer import aci --resources=tenant,application_profile --filter="Type=tenant;Name=name;Value=val1:val2:val3" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
It Will work as same as the example above with a change the filter will apply only to `tenant` resources.
```
terraformer import aci --resources=application_profile --filter="Name=annotation;Value='orchestrator:terraform':'tag'" --password=Cisco_APIC_password --username=Cisco_APIC_username --base-url=Cisco_APIC_url
```
`--parent-dn`: specify DN of parent resource. following command will import whole tree of child resources under `test` tenant.
```
terraformer import aci --resources=* --parent-dn="uni/tn-test" --password=Cisco_APIC_password --username=Cisco_APIC_username --base-url=Cisco_APIC_url
terraformer import aci --resources=* --parent-dn="uni/tn-infra:uni/tn-common"
```
If the value of the attribute has `: (colon)` in its value, then to pass this value inside --filter tag user needs to pass the entire value in `'' (single quotes)` like the above mentioned example.  


#### Building the terraformer provider

From source:
1.  Run `git clone <terraformer repo>`
2.  Run `go mod download`
3.  Run `go run build/main.go aci` for creating build ACI terraformer for the Local OS.
4.  Run `go run build/multi-build/main.go aci` to create build for multiple OS. 
5.  Run ```terraform init``` against an ```versions.tf``` file to install the plugins required for your platform. For example, if you need plugins for the ACI provider, ```versions.tf``` should contain:

```
terraform {
  required_providers {
    aci = {
      source = "ciscodevnet/aci"
    }
  }
  required_version = ">= 0.13"
}
```
Or alternatively

*  Copy your Terraform provider's plugin(s) to folder
    `~/.terraform.d/plugins/{darwin,linux}_amd64/`, as appropriate.

6. Set following environment variables
```
   ACI_USERNAME = Cisco_APIC_username
   ACI_PASSWORD = Cisco_APIC_password
   ACI_URL = Cisco_APIC_url
   ACI_CERT_NAME = Cisco_APIC_user_certificate_name
   ACI_PRIVATE_KEY = Cisco_APIC_user_private_key
```
Note: If above environment variables have not been set, use `--username`, `--password` and `--base-url` options.

From Releases:

* Linux

```
export PROVIDER={all,google,aws,kubernetes}
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-${PROVIDER}-linux-amd64
chmod +x terraformer-${PROVIDER}-linux-amd64
sudo mv terraformer-${PROVIDER}-linux-amd64 /usr/local/bin/terraformer
```
* MacOS

```
export PROVIDER={all,google,aws,kubernetes}
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-${PROVIDER}-darwin-amd64
chmod +x terraformer-${PROVIDER}-darwin-amd64
sudo mv terraformer-${PROVIDER}-darwin-amd64 /usr/local/bin/terraformer
```
