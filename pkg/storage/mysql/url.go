package mysql

import "time"

const tbUrls = "urls"

type Url struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Hash      string    `json:"hash" gorm:"column:hash"`
	Origin    string    `json:"origin" gorm:"column:origin"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	ExpiresAt time.Time `json:"expires_at" gorm:"column:expires_at"`
}

// Create url if same origin url not found
func CreateUrl(data *Url) error {
	return instance.Table(tbUrls).Where(Url{Origin: data.Origin}).FirstOrCreate(data).Error
}
