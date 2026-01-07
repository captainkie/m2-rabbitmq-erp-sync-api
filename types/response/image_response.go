package response

import "time"

type ImageResponse struct {
	ID          int       `json:"id"`
	Sku         string    `json:"sku"`
	ProductType string    `json:"product_type"`
	SyncDate    time.Time `json:"sync_date"`
	Image       string    `json:"image"`
}
