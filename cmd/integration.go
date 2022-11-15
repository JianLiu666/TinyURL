package cmd

import (
	"tinyurl/test/integration"

	"github.com/spf13/cobra"
)

var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "",
	Long:  ``,
	RunE:  RunIntegrationCmd,
}

func init() {
	rootCmd.AddCommand(integrationCmd)
}

func RunIntegrationCmd(cmd *cobra.Command, args []string) error {
	integration.Start()
	return nil
}
