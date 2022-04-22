package aci

import (
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ACIService struct {
	terraformutils.Service
}

func (a *ACIService) createClient() (*ACIClient, error) {
	log.Println("[DEBUG] initialising the ACI client")

	username := a.GetArgs()["username"].(string)
	password := a.GetArgs()["password"].(string)
	baseUrl := a.GetArgs()["base_url"].(string)
	insecure := a.GetArgs()["insecure"].(bool)
	privateKey := a.GetArgs()["private_key"].(string)
	certName := a.GetArgs()["cert_name"].(string)
	parentResource := a.GetArgs()["parent_resource"].(string)

	if password == "" {
		return GetClient(baseUrl, username, PrivateKey(privateKey), AdminCert(certName), Insecure(insecure), ParentResource(parentResource)), nil
	}
	
	return GetClient(baseUrl, username, Password(password), Insecure(insecure), ParentResource(parentResource)), nil
}
