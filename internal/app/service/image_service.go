package service

import (
	"time"

	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
)

type ImageService interface {
	CreateImage(images []request.CreateImageRequest)
	DeleteImage(targetDate time.Time)
	DeleteImageBySku(sku string)
	FindImageSkuByName(name string) (response.ImageResponse, error)
}
