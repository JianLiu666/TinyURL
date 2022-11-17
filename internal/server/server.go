package server

import (
	"tinyurl/internal/config"
	v1 "tinyurl/internal/server/api/v1"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"
	"tinyurl/internal/tracer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type server struct {
	app          *fiber.App
	serverConfig config.ServerOpts
}

func InitTinyUrlServer(kvStore kvstore.KvStore, rdb rdb.RDB, serverConfig config.ServerOpts) *server {
	app := fiber.New()

	// set middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))

	app.Use(tracer.NewFiberMiddleware(tracer.FiberConfig{
		Tracer: opentracing.GlobalTracer(),
	}))

	// set routes
	api := app.Group("/api")

	v1Api := api.Group("/v1")
	handler := v1.NewV1Handler(kvStore, rdb, serverConfig)
	v1Api.Post("/create", handler.Create)
	v1Api.Get("/:tiny_url", handler.Redirect)

	return &server{
		app:          app,
		serverConfig: serverConfig,
	}
}

func (s *server) Run() {
	go func() {
		if err := s.app.Listen(s.serverConfig.Port); err != nil {
			logrus.Panicf("starting fiber HTTP server on %s failed: %s", s.serverConfig.Port, err.Error())
		}
	}()
}

func (s *server) Shutdown() {
	if err := s.app.Shutdown(); err != nil {
		logrus.Errorf("main: shuting fiber HTTP server down failed: %v\n", err.Error())
	}
}
