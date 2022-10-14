package cmd

import (
	"tinyurl/integration"

	"github.com/spf13/cobra"
)

var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "",
	Long:  ``,
	RunE:  RunIntegrationCmd,
}

func RunIntegrationCmd(cmd *cobra.Command, args []string) error {
	integration.Start()
	return nil
}
