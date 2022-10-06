package mysql

import "time"

const tbUrls = "urls"

type Url struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Hash      string    `json:"hash" gorm:"column:hash"`
	Origin    string    `json:"origin" gorm:"column:origin"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func CreateUrl(data *Url) error {
	return instance.Table(tbUrls).Create(data).Error
}
