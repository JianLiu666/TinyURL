package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tinyurl/internal/accessor"
	"tinyurl/internal/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE:  RunServerCmd,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func RunServerCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	infra := accessor.BuildAccessor(ctx, "server")
	defer infra.Close(ctx)

	infra.InitKvStore(ctx)
	infra.InitRDB(ctx)
	infra.InitOpenTracing(ctx)

	app := server.InitTinyUrlServer(infra.KvStore, infra.RDB, infra.Config.Server)
	defer app.Shutdown()
	app.Run()

	// set graceful shutdown method
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal

	logrus.Infof("main: %s closed.\n", cmd.Name())

	return nil
}
