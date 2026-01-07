package model

import "time"

type Images struct {
	ID          int    `gorm:"primaryKey; autoIncrement"`
	Sku         string `gorm:"type:varchar(191);not null"`
	ProductType string `gorm:"type:varchar(191);"`
	SyncDate    time.Time
	Image       string `gorm:"type:longtext;"`
	Created     int64  `gorm:"autoCreateTime"`
	Updated     int64  `gorm:"autoUpdateTime"`
}
