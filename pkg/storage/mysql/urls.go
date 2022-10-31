package mysql

import (
	"errors"
	"tinyurl/pkg/storage"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const tbUrls = "urls"

func CreateUrl(data *storage.Url, isCustomAlias bool) error {
	return instance.Table(tbUrls).Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func GetUrl(tiny_url string) (res storage.Url, err error) {
	err = instance.Table(tbUrls).Where("tiny = ?", tiny_url).First(&res).Error

	// 查無資料時的初始化流程
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.Tiny = tiny_url
	}

	return
}
