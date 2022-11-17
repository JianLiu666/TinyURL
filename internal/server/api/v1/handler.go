package v1

import (
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"
)

type handler struct {
	kvStore kvstore.KvStore
	rdb     rdb.RDB
}

func NewV1Handler(kvStore kvstore.KvStore, rdb rdb.RDB) *handler {
	return &handler{
		kvStore: kvStore,
		rdb:     rdb,
	}
}
