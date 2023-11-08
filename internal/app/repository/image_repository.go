package repository

import (
	"time"

	model "github.com/captainkie/websync-api/internal/app/models"
)

type ImageRepository interface {
	CreateImage(images []model.Images)
	DeleteImage(targetDate time.Time)
	DeleteImageBySku(sku string)
	FindImageSkuByName(name string) (model.Images, error)
}
