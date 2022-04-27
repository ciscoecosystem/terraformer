package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const aESEncryptionPassphraseandKeysforConfigExportandImportClassName = "pkiExportEncryptionKey"

type AESEncryptionPassphraseandKeysforConfigExportandImportGenerator struct {
	ACIService
}

func (a *AESEncryptionPassphraseandKeysforConfigExportandImportGenerator) InitResources() error {
	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, aESEncryptionPassphraseandKeysforConfigExportandImportClassName)

	AESEncryptionPassphraseandKeysforConfigExportandImportCont, err := client.GetViaURL(dnURL)
	if err != nil {
		return err
	}

	AESEncryptionPassphraseandKeysforConfigExportandImportCount, err := strconv.Atoi(stripQuotes(AESEncryptionPassphraseandKeysforConfigExportandImportCont.S("totalCount").String()))
	if err != nil {
		return err
	}

	for i := 0; i < AESEncryptionPassphraseandKeysforConfigExportandImportCount; i++ {
		AESEncryptionPassphraseandKeysforConfigExportandImportAttr := AESEncryptionPassphraseandKeysforConfigExportandImportCont.S("imdata").Index(i).S(aESEncryptionPassphraseandKeysforConfigExportandImportClassName, "attributes")
		AESEncryptionPassphraseandKeysforConfigExportandImportDN := G(AESEncryptionPassphraseandKeysforConfigExportandImportAttr, "dn")
		if filterChildrenDn(AESEncryptionPassphraseandKeysforConfigExportandImportDN, client.parentResource) != "" {
			resource := terraformutils.NewResource(
				AESEncryptionPassphraseandKeysforConfigExportandImportDN,
				resourceNamefromDn(aESEncryptionPassphraseandKeysforConfigExportandImportClassName, AESEncryptionPassphraseandKeysforConfigExportandImportDN, i),
				"aci_encryption_key",
				"aci",
				map[string]string{
					"clear_encryption_key": "no",
				},
				[]string{},
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			a.Resources = append(a.Resources, resource)
		}
	}
	return nil
}
