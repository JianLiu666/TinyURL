package mysql

import (
	"tinyurl/pkg/storage"

	"gorm.io/gorm/clause"
)

const tbUrls = "urls"

func CreateUrl(data *storage.Url, isCustomAlias bool) error {
	return instance.Table(tbUrls).Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func GetUrl(tiny_url string) (res storage.Url, err error) {
	err = instance.Table(tbUrls).Where("tiny = ?", tiny_url).First(&res).Error
	return
}
