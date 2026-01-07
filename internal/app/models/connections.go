package model

import "time"

type Connections struct {
	ID                int    `gorm:"primaryKey; autoIncrement"`
	MessageCode       string `gorm:"type:varchar(191);"`
	MessageDesc       string `gorm:"type:varchar(191);"`
	TotalRecordAdd    string `gorm:"type:bigint; default:0;"`
	TotalRecordUpdate string `gorm:"type:bigint; default:0;"`
	TotalRecordStock  string `gorm:"type:bigint; default:0;"`
	TotalRecordStore  string `gorm:"type:bigint; default:0;"`
	SyncDate          time.Time
	Created           int64 `gorm:"autoCreateTime"`
	Updated           int64 `gorm:"autoUpdateTime"`
}
