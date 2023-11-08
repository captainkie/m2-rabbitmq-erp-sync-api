package request

import "time"

type CreateImageRequest struct {
	Sku         string    `validate:"required" json:"Sku"`
	ProductType string    `validate:"required" json:"ProductType"`
	SyncDate    time.Time `validate:"required" json:"SyncDate"`
	Image       string    `validate:"required" json:"Image"`
}
