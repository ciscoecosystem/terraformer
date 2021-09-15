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
	parentResource := a.GetArgs()["parent_resource"].(string)

	return NewClient(baseUrl, username, Password(password), Insecure(insecure), ParentResource(parentResource)), nil
}
