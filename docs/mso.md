### Use with MSO

* use `--username`, `--password`, `--base-url`, `--domain` and `--platform` options to specify server credentials for the first run. It should set respective environment variables of server credentials. Once environment variables are set there is no need to specify these options in the import command.

```
terraformer import mso --resources=schema,schema_site,site --username=Cisco_MSO_username --password=Cisco_MSO_password --base-url=Cisco_MSO_url

terraformer import mso --resources=* --excludes=schema_site

terraformer import mso --resources=schema --filter=schema=schema_id1:schema_id2
```
for more information regarding all supported flags use `--help` or `-h` flag.
```
terraformer import mso --help
```
#### Supported services

* `schema`
* `schema_site`
* `label`                                
* `schema_template_anp_epg`
* `site`                                  
* `tenant`                               
* `schema_template_bd`
* `schema_template`
* `schema_template_bd_subnet`
* `schema_template_anp`
* `schema_template_anp_epg_subnet`
* `schema_template_vrf`
* `schema_template_external_epg_contract`
* `schema_template_anp_epg_contract`
* `schema_site_anp_epg_domain`
* `schema_template_l3out`
* `schema_site_vrf_region_cidr`
* `schema_template_filter_entry`
* `schema_site_anp_epg_static_port`
* `schema_template_contract_filter`
* `schema_site_anp_epg_static_leaf`
* `schema_site_vrf_region`
* `schema_template_external_epg`


#### Attribute filters

Attribute filters allow filtering across different resource types by their attributes.

```
terraformer import mso --resources=schema_template_bd,schema_template_vrf --filter="Name=name;Value=val1:val2:val3" --username=Cisco_MSO_username --password=Cisco_MSO_password --base-url=Cisco_MSO_url
```
Will import schema template bridge domains and schema template vrfs having attribute `name` which has value either `val1`, `val2` or `val3`(supports any number of values saperated by ':'). Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applied by providing `Type=<type>` parameter. For example:
```
terraformer import mso --resources=schema_template_bd,schema_template_vrf --filter="Type=schema_template_bd;Name=name;Value=val1:val2:val3" --username=Cisco_MSO_username --password=Cisco_MSO_password --base-url=Cisco_MSO_url
```
It Will work as same as the example above with a change the filter will apply only to `schema_template_bd` service.
