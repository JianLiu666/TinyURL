package tracer

import (
	"sync"
	"time"
	"tinyurl/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var once sync.Once

func InitGlobalTracer(serverName string, opts config.JaegerOpts) {
	once.Do(func() {
		cfg := jaegercfg.Configuration{
			ServiceName: serverName,
			RPCMetrics:  opts.RPCMetrics,
			Sampler: &jaegercfg.SamplerConfig{
				Type:  opts.Sampler.Type,
				Param: float64(opts.Sampler.Param),
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:            opts.Reporter.LogSpans,
				BufferFlushInterval: time.Duration(opts.Reporter.BufferFlushInterval) * time.Second,
				LocalAgentHostPort:  opts.Reporter.LocalAgentHostPort,
			},
			Headers: &jaeger.HeadersConfig{
				TraceBaggageHeaderPrefix: opts.Headers.TraceBaggageHeaderPrefix,
				TraceContextHeaderName:   opts.Headers.TraceContextHeaderName,
			},
		}

		tracer, _, err := cfg.NewTracer()
		if err != nil {
			logrus.Panicf("failed to init jaeger: %v", err)
		}

		opentracing.SetGlobalTracer(tracer)
	})
}
