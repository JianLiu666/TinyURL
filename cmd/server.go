package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tinyurl/config"
	"tinyurl/pkg/api"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/pkg/storage/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
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
	// enable third-party logger
	logrus.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05-07:00",
	})

	// enable third-party modules
	mysql.Init()
	redis.Init()

	// enable api server
	app := fiber.New()
	api.SetRoutes(app)
	go func() {
		if err := app.Listen(config.Env().Server.Port); err != nil {
			panic(fmt.Errorf("starting fiber HTTP server on %s failed: %s", config.Env().Server.Port, err.Error()))
		}
	}()

	// set graceful shutdown method
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	<-stopSignal

	fmt.Printf("main: graceful shutdown %s...\n", cmd.Name())

	// if err := app.Shutdown(); err != nil {
	// 	fmt.Printf("shuting fiber HTTP server down failed: %v\n", err.Error())
	// 	return err
	// }

	return nil
}
