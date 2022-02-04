package cmd

import (
	"os"

	mso_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/mso"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdMsoImporter(options ImportOptions) *cobra.Command {
	username := ""
	password := ""
	baseURL := ""

	cmd := &cobra.Command{
		Use:   "mso",
		Short: "Import current state to Terraform configuration from MSO",
		Long:  "Import current state to Terraform configuration from MSO",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMsoProvider()
			err := Import(provider, options, []string{baseURL, username, password})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newMsoProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "schema,schema_site,site", "schema=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&username, "username", "", os.Getenv("MSO_USERNAME"), "YOUR_MSO_USERNAME or env param MSO_USERNAME")
	cmd.PersistentFlags().StringVarP(&password, "password", "", os.Getenv("MSO_PASSWORD"), "YOUR_MSO_PASSWORD or env param MSO_PASSWORD")
	cmd.PersistentFlags().StringVarP(&baseURL, "base-url", "", os.Getenv("MSO_URL"), "YOUR_MSO_URL or env param MSO_URL")
	return cmd
}

func newMsoProvider() terraformutils.ProviderGenerator {
	return &mso_terraforming.MSOProvider{}
}