
### Use with ACI

Example:

```
terraformer import aci --resources=tenant,application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=* --excludes=application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=tenant --filter=tenant=tenant_dn1:tenant_dn2 --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```

#### Supported services

* `tenant`
    * `cloud_applicationcontainer`
        * `cloud_epg`
    * `application_profile`
        * `endpoint_security_group_selector`
        * `endpoint_security_group`
        * `application_epg`
            * `epg_to_contract`
            * `epg_to_domain`
            * `epg_to_static_path`
    * `vrf`
        * `cloud_context_profile`
        * `any`
    * `l3_outside`
        * `l3out_ospf_external_policy`
        * `external_network_instance_profile`
            * `l3_ext_subnet`
        * `logical_node_profile`
            *`logical_node_to_fabric_profile`
                * `l3out_static_route`
                    * `l3out_static_route_next_hop`
            * `l3out_bgp_protocol_profile`
            * `logical_interface_profile`
                * `l3out_hsrp_secondary_vip`
                * `l3out_ospf_interface_profile`
                * `l3out_floating_svi`
                * `l3out_path_attachment`
                    * `l3out_path_attachment_secondary_ip`
                    * `bgp_peer_connectivity_profile`
                    * `l3out_vpc_member`
                * `l3out_bfd_interface_profile`
        * `l3out_bgp_external_policy`
        * `bgp_route_control_profile`   
    * `bridge_domain` 
        * `subnet`
        * `bd_dhcp_label`
    * `contract`
        * `contract_subject`
    * `filter`
        * `filter_entry`
    * `ospf_inteface_policy`
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
    * `cloud_applicationcontainer`
    * `cloud_aws_provider`
* `vpc_explicit_protection_group`
* `l3out_loopback_interface_profile`
* `logical_node_to_fabric_node`
* `cloud_cidr_pool`
* `cloud_endpoint_selectorfor_external_epgs`
* `cloud_endpoint_selector`
* `cloud_external_epg`
* `cloud_domain_profile`
* `cloud_vpn_gateway`

#### Attribute filters

Attribute filters allow filtering across different resource types by their attributes.

```
terraformer import aci --resources=tenant,application_profile --filter="Name=name;Value=val1:val2" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
Will import tenants and application_profiles having attribute `name` which has value either `val1` or `val2`. Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applied by providing `Type=<type>` parameter. For example:
```
terraformer import aci --resources=tenant,application_profile --filter="Type=tenant;Name=name;Value=val1:val2" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
It Will work as same as the example above with a change the filter will apply only to `tenant` resources.
```
terraformer import aci --resources=application_profile --filter="Name=annotation;Value='orchestrator:terraform':'tag'" --password=Cisco_APIC_password --username=Cisco_APIC_username --base-url=Cisco_APIC_url
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
```

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
