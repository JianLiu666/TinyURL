package kvstore

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

var _ redis.Hook = RedisHook{}

// NewRedisHook creates a new go-redis hook instance and that will collect spans using the provided tracer.
func newRedisHook(tracer opentracing.Tracer) redis.Hook {
	return &RedisHook{
		tracer: tracer,
	}
}

type RedisHook struct {
	tracer opentracing.Tracer
}

func (hook RedisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	span, ctx := hook.createSpan(ctx, cmd.FullName())
	span.SetTag("db.type", "redis")
	return ctx, nil
}

func (hook RedisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()

	span.LogKV("cmd.string", cmd.String())

	if err := cmd.Err(); err != nil {
		hook.recordError(ctx, "db.error", span, err)
	}

	return nil
}
func (hook RedisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	span, ctx := hook.createSpan(ctx, "pipeline")
	span.SetTag("db.type", "redis")
	span.SetTag("db.redis.num_cmd", len(cmds))
	return ctx, nil
}

func (hook RedisHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()

	for i, cmd := range cmds {
		if err := cmd.Err(); err != nil {
			hook.recordError(ctx, "db.error"+strconv.Itoa(i), span, err)
		}
	}
	return nil
}

func (hook RedisHook) createSpan(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		childSpan := hook.tracer.StartSpan(operationName, opentracing.ChildOf(span.Context()))
		return childSpan, opentracing.ContextWithSpan(ctx, childSpan)
	}
	return opentracing.StartSpanFromContextWithTracer(ctx, hook.tracer, operationName)
}

func (hook RedisHook) recordError(ctx context.Context, errorTag string, span opentracing.Span, err error) {
	if err != redis.Nil {
		span.SetTag(string(ext.Error), true)
		span.SetTag(errorTag, err.Error())
	}
}
