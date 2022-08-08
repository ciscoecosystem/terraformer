package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const remotePathofaFileClassName = "fileRemotePath"

type RemotePathofaFileGenerator struct {
	ACIService
}

func (a *RemotePathofaFileGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}
	client := clientImpl
	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, remotePathofaFileClassName)
	RemotePathofaFileCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}
	RemotePathofaFileCount, err := strconv.Atoi(stripQuotes(RemotePathofaFileCont.S("totalCount").String()))
	if err != nil {
		return err
	}
	for i := 0; i < RemotePathofaFileCount; i++ {
		RemotePathofaFileAttr := RemotePathofaFileCont.S("imdata").Index(i).S(remotePathofaFileClassName, "attributes")
		RemotePathofaFileDN := G(RemotePathofaFileAttr, "dn")
		if filterChildrenDn(RemotePathofaFileDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				RemotePathofaFileDN,
				resourceNamefromDn(remotePathofaFileClassName, RemotePathofaFileDN, i),
				"aci_file_remote_path",
				"aci",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
