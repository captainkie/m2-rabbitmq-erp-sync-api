package request

import "time"

type CreateImageRequest struct {
	Sku         string    `validate:"required" json:"Sku"`
	ProductType string    `validate:"required" json:"ProductType"`
	SyncDate    time.Time `validate:"required" json:"SyncDate"`
	Image       string    `validate:"required" json:"Image"`
}

type CreateImageQueueRequest struct {
	TransactionID string `json:"transaction_id"`
	Image         string `json:"image"`
	DirectoryPath string `json:"directory_path"`
	SyncDate      string `json:"sync_date"`
}
