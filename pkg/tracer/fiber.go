package tracer

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func NewFiberMiddleware(config FiberConfig) fiber.Handler {
	// Set default config
	cfg := setDefaultFiberConfig(config)
	return func(c *fiber.Ctx) error {
		// Filter the request no need for tracing
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		var span opentracing.Span
		operationName := cfg.OperationName(c)
		tracer := cfg.Tracer
		header := make(http.Header)

		// traverse the header from fasthttp
		// and then set to http header for extract trace information
		c.Request().Header.VisitAll(func(key, value []byte) {
			header.Set(string(key), string(value))
		})

		// extract trace-id from header
		sc, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
		if err != nil {
			span = tracer.StartSpan(operationName, opentracing.ChildOf(sc))
		} else if !cfg.SkipSpanWithoutParent {
			span = tracer.StartSpan(operationName)
		} else {
			return c.Next()
		}

		cfg.Modify(c, span)

		defer func() {
			status := c.Response().StatusCode()
			ext.HTTPStatusCode.Set(span, uint16(status))
			if status >= fiber.StatusInternalServerError {
				ext.Error.Set(span, true)
			}
			span.Finish()
		}()
		return c.Next()
	}
}

type FiberConfig struct {
	Tracer                opentracing.Tracer
	OperationName         func(*fiber.Ctx) string
	Filter                func(*fiber.Ctx) bool
	Modify                func(*fiber.Ctx, opentracing.Span)
	SkipSpanWithoutParent bool
}

var fiberDefaultConfig = FiberConfig{
	Tracer: opentracing.NoopTracer{},
	OperationName: func(ctx *fiber.Ctx) string {
		return "HTTP " + ctx.Method() + " URL: " + ctx.Path()
	},
	Filter: nil,
	Modify: func(ctx *fiber.Ctx, span opentracing.Span) {
		span.SetTag("http.method", ctx.Method())
		span.SetTag("http.remote_address", ctx.IP())
		span.SetTag("http.path", ctx.Path())
		span.SetTag("http.host", ctx.Hostname())
		span.SetTag("http.url", ctx.OriginalURL())
	},
	SkipSpanWithoutParent: false,
}

// setDefaultFiberConfig function to return defalut values
func setDefaultFiberConfig(config ...FiberConfig) FiberConfig {
	// return default config if no config provided
	if len(config) < 1 {
		return fiberDefaultConfig
	}

	cfg := config[0]

	if cfg.Tracer == nil {
		cfg.Tracer = fiberDefaultConfig.Tracer
	}

	if cfg.OperationName == nil {
		cfg.OperationName = fiberDefaultConfig.OperationName
	}

	if cfg.Modify == nil {
		cfg.Modify = fiberDefaultConfig.Modify
	}

	return cfg
}
