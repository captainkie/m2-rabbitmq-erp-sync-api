package repository

import (
	"errors"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p *ProductRepositoryImpl) ConfigurableProduct(product model.ConfigurableProducts) {
	result := p.Db.Create(&product)
	helpers.ErrorPanic(result.Error)
}

func (p *ProductRepositoryImpl) FindBySku(sku string) (model.ConfigurableProducts, error) {
	var products model.ConfigurableProducts
	result := p.Db.First(&products, "first_child_sku = ?", sku)

	if result != nil {
		return products, nil
	} else {
		return products, errors.New("products is not found")
	}
}
