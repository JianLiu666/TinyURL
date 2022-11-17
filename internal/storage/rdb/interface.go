package rdb

import (
	"context"
	"tinyurl/internal/storage"

	"github.com/opentracing/opentracing-go"
)

const tbUrls = "urls"

type RDB interface {
	SetOpenTracing(tracer opentracing.Tracer)
	CreateUrl(ctx context.Context, data *storage.Url, isCustomAlias bool) (bool, error)
	GetUrl(ctx context.Context, tiny_url string) (res storage.Url, err error)
	Shutdown(ctx context.Context)
}
