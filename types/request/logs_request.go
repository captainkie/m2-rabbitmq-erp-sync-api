package request

import "time"

type CreateLogRequest struct {
	TransactionID string    `validate:"required" json:"transaction_id"`
	Status        string    `validate:"required" json:"status"`
	StatusCode    int       `validate:"required" json:"status_code"`
	Message       string    `validate:"required" json:"message"`
	SyncJson      string    `validate:"required" json:"sync_json"`
	SyncData      string    `validate:"required" json:"sync_data"`
	SyncDate      time.Time `validate:"required" json:"sync_date"`
}

type CreateLogImageRequest struct {
	TransactionID string    `validate:"required" json:"transaction_id"`
	Status        string    `validate:"required" json:"status"`
	StatusCode    int       `validate:"required" json:"status_code"`
	Message       string    `validate:"required" json:"message"`
	SyncImage     string    `validate:"required" json:"sync_image"`
	SyncJson      string    `validate:"required" json:"sync_json"`
	SyncData      string    `validate:"required" json:"sync_data"`
	SyncDate      time.Time `validate:"required" json:"sync_date"`
}

type CreateLogDailyRequest struct {
	TransactionID string    `validate:"required" json:"transaction_id"`
	OrderID       string    `validate:"required" json:"order_id"`
	Status        string    `validate:"required" json:"status"`
	StatusCode    int       `validate:"required" json:"status_code"`
	Message       string    `validate:"required" json:"message"`
	SyncJson      string    `validate:"required" json:"sync_json"`
	SyncData      string    `validate:"required" json:"sync_data"`
	SyncDate      time.Time `validate:"required" json:"sync_date"`
}
