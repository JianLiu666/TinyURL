package tracer

import (
	"fmt"
	"sync"
	"time"
	"tinyurl/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var once sync.Once

func InitGlobalTracer() {
	once.Do(func() {
		cfg := jaegercfg.Configuration{
			ServiceName: config.Env().Server.Name,
			RPCMetrics:  config.Env().Jaeger.RPCMetrics,
			Sampler: &jaegercfg.SamplerConfig{
				Type:  config.Env().Jaeger.Sampler.Type,
				Param: float64(config.Env().Jaeger.Sampler.Param),
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:            config.Env().Jaeger.Reporter.LogSpans,
				BufferFlushInterval: time.Duration(config.Env().Jaeger.Reporter.BufferFlushInterval) * time.Second,
				LocalAgentHostPort:  config.Env().Jaeger.Reporter.LocalAgentHostPort,
			},
			Headers: &jaeger.HeadersConfig{
				TraceBaggageHeaderPrefix: config.Env().Jaeger.Headers.TraceBaggageHeaderPrefix,
				TraceContextHeaderName:   config.Env().Jaeger.Headers.TraceContextHeaderName,
			},
		}

		tracer, _, err := cfg.NewTracer()
		if err != nil {
			fmt.Println("init jaeger failed: ", err)
		}

		opentracing.SetGlobalTracer(tracer)
		fmt.Println("init jaeger successful.")
	})
}
