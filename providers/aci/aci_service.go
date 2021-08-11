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

	return NewClient(baseUrl, username, Password(password), Insecure(insecure)), nil
}
