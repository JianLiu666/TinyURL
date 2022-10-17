package api

import (
	"time"
	v1api "tinyurl/pkg/api/v1"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
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
	app.Use(pprof.New())
	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "TinyURL Monitor",
		Refresh: 1 * time.Second,
	}))

	prometheus := fiberprometheus.New("tinyurl")
	prometheus.RegisterAt(app, "/prometheus/metrics")
	app.Use(prometheus.Middleware)

}

func setLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))
}
