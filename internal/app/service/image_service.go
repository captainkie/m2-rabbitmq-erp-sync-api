package service

import (
	"time"

	"github.com/captainkie/websync-api/types/request"
)

type ImageService interface {
	CreateImage(images []request.CreateImageRequest)
	DeleteImage(targetDate time.Time)
	DeleteImageBySku(sku string)
}
