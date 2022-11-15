package api

import (
	"time"
	v1api "tinyurl/internal/api/v1"
	"tinyurl/internal/tracer"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/opentracing/opentracing-go"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/create", v1api.Create)
	v1.Get("/:tiny_url", v1api.Redirect)
}

func SetMonitor(app *fiber.App) {
	// enable fiber monitor plugin
	app.Use(pprof.New())
	app.Get("/fiber/monitor", monitor.New(monitor.Config{
		Title:   "TinyURL Monitor",
		Refresh: 1 * time.Second,
	}))

	// enable prometheus metrics plugin
	prometheus := fiberprometheus.New("tinyurl")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
}

func SetLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))
}

func SetTracer(app *fiber.App) {
	// enable jaeger plugin
	app.Use(tracer.NewFiberMiddleware(tracer.FiberConfig{
		Tracer: opentracing.GlobalTracer(),
	}))
}
