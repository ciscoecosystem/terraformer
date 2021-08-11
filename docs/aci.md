
### Use with ACI

Example:

```
terraformer import aci --resources=tenant,application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=* --excludes=application_profile --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
terraformer import aci --resources=tenant --filter=tenant_dn1:tenant_dn2 --username=Cisco_APIC_username --password=Cisco_APIC_password --base-url=Cisco_APIC_url
```

#### Supported services

* `tenant`
    * `application_profile`
        * `application_epg`
    * `bridge_domain` 
    * `contract`
        * `contract_subject`
    * `filter`
        * `filter_entry`
    * `subnet`   
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



