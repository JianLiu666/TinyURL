package storage

import "time"

type Url struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Tiny      string    `json:"tiny" gorm:"column:tiny"`
	Origin    string    `json:"origin" gorm:"column:origin"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	ExpiresAt time.Time `json:"expires_at" gorm:"column:expires_at"`
}
