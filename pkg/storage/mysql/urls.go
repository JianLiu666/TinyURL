package mysql

import (
	"errors"
	"time"
	"tinyurl/util"

	"gorm.io/gorm"
)

const tbUrls = "urls"

type Url struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Tiny      string    `json:"tiny" gorm:"column:tiny"`
	Origin    string    `json:"origin" gorm:"column:origin"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	ExpiresAt time.Time `json:"expires_at" gorm:"column:expires_at"`
}

func CreateUrl(data *Url, isCustomAlias bool) error {
	alias := Url{}

	return instance.Transaction(func(tx *gorm.DB) error {
		// 檢查資料庫是否已經存在相同的短網址
		if err := tx.Table(tbUrls).Select("origin").Where("tiny = ?", data.Tiny).First(&alias).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			// 短網址未發生碰撞, 直接將短網址的 meatadata 寫入資料庫
			// 若本來就存在相同的原始網址, 直接覆蓋新的 meatadata 即可
			return tx.Table(tbUrls).Where(Url{Origin: data.Origin}).FirstOrCreate(&data).Error
		}

		// 已經存在相同資料, 不處理也不增加 expired time
		if alias.Origin == data.Origin {
			return nil
		}

		// 短網址發生碰撞, 且短網址是用戶自定義
		if isCustomAlias {
			return gorm.ErrInvalidData
		}

		// 將 timestamp 作為後綴詞加入短網址
		suffix := util.Base10ToBase62(uint64(time.Now().UnixMilli()))
		data.Tiny += suffix
		return tx.Table(tbUrls).Where(Url{Origin: data.Origin}).FirstOrCreate(&data).Error
	})
}

func GetUrl(tiny_url string) (res Url, err error) {
	err = instance.Table(tbUrls).Where("tiny = ?", tiny_url).First(&res).Error
	return
}
