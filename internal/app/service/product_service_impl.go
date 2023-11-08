package service

import (
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// Create implements ProductService
func (p *ProductServiceImpl) CreateConfigurableProduct(product request.ConfigurableProductRequest) {
	newProduct := model.ConfigurableProducts{
		Sku:           product.Sku,
		FirstChildSku: product.FirstChildSku,
	}

	p.ProductRepository.ConfigurableProduct(newProduct)
}

// FindById implements UsersService
func (p *ProductServiceImpl) FindBySku(sku string) response.ConfigurableProductResponse {
	productData, err := p.ProductRepository.FindBySku(sku)
	helpers.ErrorPanic(err)

	productResponse := response.ConfigurableProductResponse{
		ID:  productData.ID,
		Sku: productData.Sku,
	}

	return productResponse
}
