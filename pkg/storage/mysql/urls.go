package mysql

import (
	"database/sql"
	"fmt"
	"time"
	"tinyurl/pkg/storage"
	"tinyurl/util"

	"gorm.io/gorm"
)

const tbUrls = "urls"

func CreateUrl(data *storage.Url, isCustomAlias bool) error {
	txOptions := &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	alias := storage.Url{}

	return instance.Transaction(func(tx *gorm.DB) error {

		// 檢查資料庫是否已經存在相同的短網址
		if result := tx.Table(tbUrls).Where("tiny = ? OR origin = ?", data.Tiny, data.Origin).Limit(1).Find(&alias); result.RowsAffected == 0 {
			// 查無資料, 直接寫入
			return tx.Table(tbUrls).Create(&data).Error
		}

		// 已經存在相同資料
		if alias.Origin == data.Origin {
			// 未預期的錯誤: 相同的原始網址, 但短網址不同
			if alias.Tiny != data.Tiny {
				return fmt.Errorf("unexpected error: Req:%v, DB:%v", data, alias)
			}

			// 以後可以考慮是否要延長 expiration time
			return gorm.ErrInvalidData
		}

		// 短網址發生碰撞, 且短網址是用戶自定義
		if isCustomAlias {
			return gorm.ErrInvalidData
		}

		// 將 timestamp 作為後綴詞加入短網址
		suffix := util.Base10ToBase62(uint64(time.Now().UnixMilli()))
		data.Tiny += suffix
		return tx.Table(tbUrls).Create(&data).Error

	}, txOptions)
}

func GetUrl(tiny_url string) (res storage.Url, err error) {
	err = instance.Table(tbUrls).Where("tiny = ?", tiny_url).First(&res).Error
	return
}
