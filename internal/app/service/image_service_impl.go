package service

import (
	"time"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/go-playground/validator/v10"
)

type ImageServiceImpl struct {
	ImageRepository repository.ImageRepository
	Validate        *validator.Validate
}

func NewImageServiceImpl(imageRepository repository.ImageRepository, validate *validator.Validate) ImageService {
	return &ImageServiceImpl{
		ImageRepository: imageRepository,
		Validate:        validate,
	}
}

// Create implements ImageService
func (i *ImageServiceImpl) CreateImage(images []request.CreateImageRequest) {
	var newImages []model.Images
	for _, value := range images {
		newImages = append(newImages, model.Images{
			Sku:         value.Sku,
			ProductType: value.ProductType,
			SyncDate:    value.SyncDate,
			Image:       value.Image,
		})
	}

	i.ImageRepository.CreateImage(newImages)
}

func (i *ImageServiceImpl) DeleteImage(targetDate time.Time) {
	i.ImageRepository.DeleteImage(targetDate)
}

func (i *ImageServiceImpl) DeleteImageBySku(sku string) {
	i.ImageRepository.DeleteImageBySku(sku)
}

func (i *ImageServiceImpl) FindImageSkuByName(name string) (response.ImageResponse, error) {
	productData, err := i.ImageRepository.FindImageSkuByName(name)
	if err != nil {
		return response.ImageResponse{}, err
	}

	productResponse := response.ImageResponse{
		ID:          productData.ID,
		Sku:         productData.Sku,
		ProductType: productData.ProductType,
		SyncDate:    productData.SyncDate,
		Image:       productData.Image,
	}

	return productResponse, nil
}
