package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"tinyurl/internal/api"
	"tinyurl/internal/config"
	"tinyurl/internal/storage/mysql"
	"tinyurl/internal/storage/redis"
	"tinyurl/internal/tracer"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	gormopentracing "gorm.io/plugin/opentracing"
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
	// enable opentracing  with jaeger
	tracer.InitGlobalTracer(config.Env().Server.Name, config.Env().Jaeger)

	// enable redis client
	mysql.Init()
	if err := mysql.GetInstance().Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	// enable redis client
	redis.Init()
	redis.GetInstance().AddHook(tracer.NewRedisHook(opentracing.GlobalTracer()))

	// enable api server
	app := fiber.New()
	api.SetRoutes(app)
	api.SetMonitor(app)
	api.SetLogger(app)
	api.SetTracer(app)
	go func() {
		if err := app.Listen(config.Env().Server.Port); err != nil {
			logrus.Panicf("starting fiber HTTP server on %s failed: %s", config.Env().Server.Port, err.Error())
		}
	}()

	// set graceful shutdown method
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal

	logrus.Infof("main: graceful shutdown %s...", cmd.Name())

	if err := app.Shutdown(); err != nil {
		logrus.Errorf("main: shuting fiber HTTP server down failed: %v\n", err.Error())
		return err
	}

	logrus.Infof("main: %s closed.\n", cmd.Name())

	return nil
}
