package mso

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type MSOProvider struct {
	terraformutils.Provider
	baseURL  string
	username string
	password string
	insecure bool
}

func (p MSOProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p MSOProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"aci": map[string]interface{}{
				"username": p.username,
				"password": p.password,
				"url":      p.baseURL,
			},
		},
	}
}

func (p *MSOProvider) Init(args []string) error {
	p.baseURL = args[0]
	p.username = args[1]
	p.password = args[2]
	p.insecure = true
	os.Setenv("MSO_URL", p.baseURL)
	os.Setenv("MSO_USERNAME", p.username)
	os.Setenv("MSO_PASSWORD", p.password)
	return nil
}

func (p *MSOProvider) GetName() string {
	return "mso"
}

func (p *MSOProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username": p.username,
		"password": p.password,
		"base_url": p.baseURL,
		"insecure": p.insecure,
	})
	return nil
}

func (p *MSOProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"schema": &SchemaGenerator{},
	}
}
