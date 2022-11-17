package v1

import (
	"tinyurl/internal/config"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"
)

type handler struct {
	kvStore      kvstore.KvStore
	rdb          rdb.RDB
	serverConfig config.ServerOpts
}

func NewV1Handler(kvStore kvstore.KvStore, rdb rdb.RDB, serverConfig config.ServerOpts) *handler {
	return &handler{
		kvStore:      kvStore,
		rdb:          rdb,
		serverConfig: serverConfig,
	}
}
