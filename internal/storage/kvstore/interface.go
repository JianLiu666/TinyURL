package kvstore

import (
	"context"
	"time"
	"tinyurl/internal/storage"

	"github.com/opentracing/opentracing-go"
)

const (
	ErrNotFound = iota
	ErrInvalidData
	ErrUnexpected
	ErrKeyNotFound
)

type KvStore interface {
	SetOpenTracing(tracer opentracing.Tracer)
	SetTinyUrl(ctx context.Context, data *storage.Url, expiration time.Duration) int
	GetOriginUrl(ctx context.Context, tiny string) (string, int)
	CheckTinyUrl(ctx context.Context, data *storage.Url, isCustomAlias bool, retryCount int) int
	Shutdown(ctx context.Context)
}

func getTinyKey(tiny string) string {
	return "tiny:" + tiny
}
