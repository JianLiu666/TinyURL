package server

import (
	"tinyurl/internal/config"
	v1Api "tinyurl/internal/server/api/v1"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/swagger"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type server struct {
	app          *fiber.App
	serverConfig config.ServerOpts
}

// @title        TinyURL Swagger
// @version      1.0
// @description  Tiny URL swagger documentation
//
// @contact.name   API Support
// @contact.url    https://github.com/JianLiu666/TinyURL/issues
// @contact.email  jianliu0616@gmail.com
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host      localhost:6600
func InitTinyUrlServer(kvStore kvstore.KvStore, rdb rdb.RDB, serverConfig config.ServerOpts, enableOpentracing bool) *server {
	app := fiber.New()

	// set web server logger format
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))

	// enable jaeger plugin
	if enableOpentracing {
		app.Use(newFiberMiddleware(fiberConfig{
			Tracer: opentracing.GlobalTracer(),
		}))
	}

	// enable prometheus metrics plugin
	prometheus := fiberprometheus.New("tinyurl")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// enable pprof plugin
	app.Use(pprof.New())

	// set routes
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:6600/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")
	handler := v1Api.NewV1Handler(kvStore, rdb, serverConfig)
	v1.Post("/create", handler.Create)
	v1.Get("/:tiny_url", handler.Redirect)

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
