package mysql

import (
	"fmt"
	"sync"
	"time"
	"tinyurl/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var instance *gorm.DB

func GetInstance() *gorm.DB {
	return instance
}

func Init() {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Env().MySQL.UserName,
			config.Env().MySQL.Password,
			config.Env().MySQL.Address,
			config.Env().MySQL.DBName,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = db

		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(config.Env().MySQL.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.Env().MySQL.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(config.Env().MySQL.ConnMaxLifetime * int(time.Minute)))

		fmt.Println("connect to mysql successful.")
	})
}
