package rdb

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"tinyurl/internal/storage"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormopentracing "gorm.io/plugin/opentracing"
)

type mysqlClient struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func NewMySqlClient(ctx context.Context, dsn string, connMaxLifetime time.Duration, maxOpenConns, maxIdleConns int) RDB {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("failed to open database by gorm: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		logrus.Panicf("failed to get sql.DB : %v", err)
	}

	sqlDB.SetConnMaxLifetime(connMaxLifetime)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	return &mysqlClient{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

func (c *mysqlClient) SetOpenTracing(tracer opentracing.Tracer) {
	if err := c.gormDB.Use(gormopentracing.New()); err != nil {
		logrus.Panicf("failed to use open tracing: %v", err)
	}
}

func (c *mysqlClient) Shutdown(ctx context.Context) {
	if err := c.sqlDB.Close(); err != nil {
		logrus.Panicf("failed to close sql.DB : %v", err)
	}
}

func (c *mysqlClient) Exec(sql string) {
	c.gormDB.Exec(sql)
}

func (c *mysqlClient) CreateUrl(ctx context.Context, data *storage.Url, isCustomAlias bool) (bool, error) {
	tx := c.gormDB.WithContext(ctx).Table(tbUrls).Clauses(clause.OnConflict{UpdateAll: true}).Create(&data)
	return tx.RowsAffected == 0, tx.Error
}

func (c *mysqlClient) GetUrl(ctx context.Context, tiny_url string) (res storage.Url, err error) {
	err = c.gormDB.WithContext(ctx).Table(tbUrls).Where("tiny = ?", tiny_url).First(&res).Error

	// 查無資料時的初始化流程
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.Tiny = tiny_url
	}

	return
}
