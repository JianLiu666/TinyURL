package rdb

import (
	"context"
	"tinyurl/internal/storage"
)

const tbUrls = "urls"

type RDB interface {
	CreateUrl(ctx context.Context, data *storage.Url, isCustomAlias bool) (bool, error)
	GetUrl(ctx context.Context, tiny_url string) (res storage.Url, err error)
	Shutdown(ctx context.Context)
}
