package model

import "time"

type ConnectionLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type AddLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type UpdateLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type StockLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type StoreLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type PostflagLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type ImageLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncImage     string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}

type DailysaleLogs struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	OrderID       string `gorm:"type:varchar(191);not null"`
	Status        string `gorm:"type:varchar(191);"`
	StatusCode    int    `gorm:"type:int;"`
	Message       string `gorm:"type:longtext;"`
	SyncJson      string `gorm:"type:longtext;"`
	SyncData      string `gorm:"type:longtext;"`
	SyncDate      time.Time
	Created       int64 `gorm:"autoCreateTime"`
	Updated       int64 `gorm:"autoUpdateTime"`
}
