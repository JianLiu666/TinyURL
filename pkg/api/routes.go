package api

import (
	"time"
	"tinyurl/config"
	v1api "tinyurl/pkg/api/v1"

	"github.com/ansrivas/fiberprometheus/v2"
	fibertracing "github.com/aschenmaker/fiber-opentracing"
	"github.com/aschenmaker/fiber-opentracing/fjaeger"
	"github.com/opentracing/opentracing-go"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	jconfig "github.com/uber/jaeger-client-go/config"
)

func SetRoutes(app *fiber.App) {
	setMonitor(app)
	setLogger(app)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/create", v1api.Create)
	v1.Get("/:tiny_url", v1api.Redirect)
}

func setMonitor(app *fiber.App) {
	// enable for jaeger
	fjaeger.New(fjaeger.Config{
		ServiceName: config.Env().Server.Name,
		Reporter: &jconfig.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  config.Env().Jaeger.Address,
		},
	})
	app.Use(fibertracing.New(fibertracing.Config{
		Tracer: opentracing.GlobalTracer(),
		OperationName: func(ctx *fiber.Ctx) string {
			return "HTTP " + ctx.Method() + " URL: " + ctx.Path()
		},
	}))

	// enable fiber monitor middleware
	app.Use(pprof.New())
	app.Get("/fiber/monitor", monitor.New(monitor.Config{
		Title:   "TinyURL Monitor",
		Refresh: 1 * time.Second,
	}))

	// enable for prometheus
	prometheus := fiberprometheus.New("tinyurl")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
}

func setLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))
}
