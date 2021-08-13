
### Use with ACI

Example:

```
terraformer import aci --resources=tenant,application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=* --excludes=application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=tenant --filter=tenant=tenant_dn1:tenant_dn2 --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```

#### Supported services

* `tenant`
    * `application_profile`
        * `application_epg`
    * `bridge_domain` 
        * `subnet`
    * `contract`
        * `contract_subject`
    * `filter`
        * `filter_entry`   
* `vpc_explicit_protection_group`

#### Attribute filters

Attribute filters allow filtering across different resource types by its attributes.

```
terraformer import aci --resources=tenant,application_profile --filter="Name=name;Value=val1:val2" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
Will import tenants and application_profiles having attribute `name` which has value either `val1` or `val2`. Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applicable to by providing `Type=<type>` parameter. For example:
```
terraformer import aci --resources=tenant,application_profile --filter="Type=tenant;Name=name;Value=val1:val2" --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```
Will work as same as example above with a change the filter will be applicable only to `tenant` resources.

#### NOTE 

Steps to build the terraformer build:

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
