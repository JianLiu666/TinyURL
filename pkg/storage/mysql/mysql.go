package mysql

import (
	"fmt"
	"sync"
	"tinyurl/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var instance *gorm.DB

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
		fmt.Println("connect to mysql successful.")
	})
}
