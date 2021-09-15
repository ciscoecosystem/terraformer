#### Building the terraformer provider

## Install Go Lang
* https://golang.org/doc/install

## Building the terraform provider
1.  Run `git clone https://github.com/CiscoDevNet/terraform-provider-aci.git`
2.  Run `git checkout parent-dn-update`
3.  Run `make fmt && make` 
4.  Copy provider executable from this path `$GOPATH/bin/terraform-provider-aci`
5.  Put provider executable at `$HOME/.terraform.d/registry.terraform.io/ciscodevnet/aci/0.7.1/darwin_amd64`

## Building terraformer
1.  Run `git clone https://github.com/ciscoecosystem/terraformer.git`
2.  Run `git checkout aci-provider`
3.  Run `go run build/aci-multi-build/main.go`
4.  Rename prefered OS build to `terraformer`
5.  Move executable to prefered location and make sure PATH is set in Environment Variables

## Running terraformer
1.  Terraformer uses terraform provider to refresh state of imported resources hence, we need to make sure that provider authentication is performed via following Environment Variables
```
   ACI_USERNAME = Cisco_APIC_username
   ACI_PASSWORD = Cisco_APIC_password
   ACI_URL = Cisco_APIC_url
```
2.  Create terraform file `main.tf` with following configuration
```
terraform {
  required_providers {
    aci = {
      source = "registry.terraform.io/ciscodevnet/aci"
    }
  }
}
```
3.  Run `terraform init`
4.  One should be able to run all kinds of terraformer import statements following this configuration

## NOTE
1.  The state file generated through terraformer is in terraform version `0.12.31`. Version `0.12.X` is not supported after terraform `0.13.X`. Hence, need to upgrade state file to enable support to version >= `0.13`.
2.  Run following commands using terraform executable of versions `0.13.X`.
```
terraform 0.13upgrade
terraform init
terraform apply
```