package cmd

import (
	"os"

	aci_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/aci"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdACIImporter(options ImportOptions) *cobra.Command {
	username := ""
	password := ""
	baseURL := ""
	certName := ""
	privateKey := ""

	parentResource := ""
	cmd := &cobra.Command{
		Use:   "aci",
		Short: "Import current state to Terraform configuration from ACI",
		Long:  "Import current state to Terraform configuration from ACI",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newACIProvider()
			err := Import(provider, options, []string{baseURL, username, password, certName, privateKey, parentResource})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newACIProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "tenant,vrf,subnet", "tenant=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&username, "username", "", os.Getenv("ACI_USERNAME"), "YOUR_ACI_USERNAME or env param ACI_USERNAME")
	cmd.PersistentFlags().StringVarP(&password, "password", "", os.Getenv("ACI_PASSWORD"), "YOUR_ACI_PASSWORD or env param ACI_PASSWORD")
	cmd.PersistentFlags().StringVarP(&privateKey, "private-key", "", os.Getenv("ACI_PRIVATE_KEY"), "YOUR_ACI_PRIVATE_KEY or env param ACI_PRIVATE_KEY")
	cmd.PersistentFlags().StringVarP(&certName, "cert-name", "", os.Getenv("ACI_CERT_NAME"), "YOUR_ACI_CERT_NAME or env param ACI_CERT_NAME")
	cmd.PersistentFlags().StringVarP(&baseURL, "base-url", "", os.Getenv("ACI_URL"), "YOUR_ACI_URL or env param ACI_URL")
	cmd.PersistentFlags().StringVarP(&parentResource, "parent-dn", "", "", "import children resources of a particular DN")

	return cmd
}

func newACIProvider() terraformutils.ProviderGenerator {
	return &aci_terraforming.ACIProvider{}
}