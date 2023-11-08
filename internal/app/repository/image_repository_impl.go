package repository

import (
	"time"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"

	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
	Db *gorm.DB
}

func NewImageRepositoryImpl(Db *gorm.DB) ImageRepository {
	return &ImageRepositoryImpl{Db: Db}
}

func (i *ImageRepositoryImpl) CreateImage(images []model.Images) {
	result := i.Db.Create(images)
	helpers.ErrorPanic(result.Error)
}

// Update queue type postflag implements QueueRepository
func (i *ImageRepositoryImpl) DeleteImage(targetDate time.Time) {
	var images model.Images
	result := i.Db.Where("sync_date < ?", targetDate).Delete(&images)
	helpers.ErrorPanic(result.Error)
}

func (i *ImageRepositoryImpl) DeleteImageBySku(sku string) {
	var images model.Images
	result := i.Db.Where("sku = ?", sku).Delete(&images)
	helpers.ErrorPanic(result.Error)
}
