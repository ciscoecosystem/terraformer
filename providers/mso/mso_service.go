package mso

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type MSOService struct {
	terraformutils.Service
}

func (s MSOService) getClient() (*client.Client, error) {
	username := s.GetArgs()["username"].(string)
	password := s.GetArgs()["password"].(string)
	baseUrl := s.GetArgs()["base_url"].(string)
	insecure := s.GetArgs()["insecure"].(bool)
	domain := s.GetArgs()["domain"].(string)
	platform := s.GetArgs()["platform"].(string)
	if platform == "" {
		platform = "mso"
	}
	if password != "" && username != "" && baseUrl != "" {
		return client.GetClient(baseUrl, username, client.Password(password), client.Insecure(insecure), client.Domain(domain), client.Platform(platform)), nil
	} else {
		return nil, fmt.Errorf("one of password, username and base_url is missing")
	}
}
