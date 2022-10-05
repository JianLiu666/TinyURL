package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tinyurl/config"
	"tinyurl/pkg/api"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// ApiServerCmd api server init point
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE:  RunServerCmd,
}

func RunServerCmd(cmd *cobra.Command, args []string) error {
	app := fiber.New()

	api.SetRoutes(app)

	go func() {
		if err := app.Listen(config.Env().Server.Port); err != nil {
			panic(fmt.Errorf("starting fiber HTTP server on %s failed: %s", config.Env().Server.Port, err.Error()))
		}
	}()

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
