package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tinyurl/internal/accessor"
	"tinyurl/internal/server"

	_ "tinyurl/docs/swagger"

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

	// initial infrastructure
	infra := accessor.BuildAccessor()
	defer infra.Close(ctx)

	infra.InitKvStore(ctx)
	infra.InitRDB(ctx)

	// initial opentracing mechanism
	if infra.Config.Jaeger.Enable {
		infra.InitOpenTracing(ctx)
	}

	// initial tiny url server
	app := server.InitTinyUrlServer(
		infra.KvStore,
		infra.RDB,
		infra.Config.Server,
		infra.Config.Jaeger.Enable,
	)
	defer app.Shutdown()

	// enable tiny url server
	app.Run()

	// set graceful shutdown method
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal

	logrus.Infof("main: %s closed.\n", cmd.Name())

	return nil
}
