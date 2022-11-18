package cmd

import (
	"context"
	"tinyurl/internal/accessor"
	"tinyurl/internal/integration"

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
	ctx := context.Background()

	infra := accessor.BuildAccessor()
	defer infra.Close(ctx)

	infra.InitKvStore(ctx)
	infra.InitRDB(ctx)

	app := integration.NewIntegrationTester(infra.KvStore, infra.RDB, infra.Config.Server)
	app.Start()

	return nil
}
