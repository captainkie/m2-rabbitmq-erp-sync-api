package model

type ConfigurableProducts struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	Sku           string `gorm:"type:varchar(191);not null"`
	FirstChildSku string `gorm:"type:varchar(191);not null"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}
