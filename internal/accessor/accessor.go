package accessor

import (
	"context"
	"fmt"
	"sync"
	"time"
	"tinyurl/internal/config"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/internal/storage/rdb"
	"tinyurl/internal/tracer"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type shutdownHandler func(context.Context)

type accessor struct {
	shutdownOnce     sync.Once
	shutdownHandlers []shutdownHandler

	Config  *config.Config  // configuration managment
	KvStore kvstore.KvStore // key-value store instance
	RDB     rdb.RDB         // relational database instance
}

func BuildAccessor() *accessor {
	return &accessor{
		Config: config.NewFromViper(),
	}
}

func (a *accessor) Close(ctx context.Context) {
	a.shutdownOnce.Do(func() {
		logrus.Infoln("start to close accessors.")
		for _, handler := range a.shutdownHandlers {
			handler(ctx)
		}
	})

	logrus.Infoln("all accessors closed.")
}

func (a *accessor) InitOpenTracing(ctx context.Context) {
	tracer.InitGlobalTracer(a.Config.Server.Name, a.Config.Jaeger)

	if a.KvStore != nil {
		a.KvStore.SetOpenTracing(opentracing.GlobalTracer())
	}

	if a.RDB != nil {
		a.RDB.SetOpenTracing(opentracing.GlobalTracer())
	}

	logrus.Infoln("initial open tracing accessor successful.")
}

func (a *accessor) InitKvStore(ctx context.Context) {
	a.KvStore = kvstore.NewRedisClient(ctx,
		a.Config.Redis.Address,
		a.Config.Redis.Password,
		a.Config.Redis.DB,
	)

	a.shutdownHandlers = append(a.shutdownHandlers, func(c context.Context) {
		a.KvStore.Shutdown(c)
		logrus.Infoln("key-value store accessor closed.")
	})

	logrus.Infoln("initial key-value store accessor successful.")
}

func (a *accessor) InitRDB(ctx context.Context) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		a.Config.MySQL.UserName,
		a.Config.MySQL.Password,
		a.Config.MySQL.Address,
		a.Config.MySQL.DBName,
	)

	a.RDB = rdb.NewMySqlClient(ctx, dsn,
		time.Duration(a.Config.MySQL.ConnMaxLifetime)*time.Minute,
		a.Config.MySQL.MaxOpenConns,
		a.Config.MySQL.MaxIdleConns,
	)

	a.shutdownHandlers = append(a.shutdownHandlers, func(c context.Context) {
		a.RDB.Shutdown(c)
		logrus.Infoln("relational database accessor closed.")
	})

	logrus.Infoln("initial relational database accessor successful.")
}
