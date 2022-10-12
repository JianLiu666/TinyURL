package api

import (
	"time"
	v1api "tinyurl/pkg/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetRoutes(app *fiber.App) {
	setMonitor(app)
	setLogger(app)

	app.Get("/:tiny_url", Redirect)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/create", v1api.Create)
}

func setMonitor(app *fiber.App) {
	app.Use(pprof.New())
	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "TinyURL Monitor",
		Refresh: 1 * time.Second,
	}))
}

func setLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] | ${ip} | ${latency} | ${status} | ${method} | ${path} | Req: ${body} | Resp: ${resBody}\n",
	}))
}
