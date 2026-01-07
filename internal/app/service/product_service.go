package service

import (
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
)

type ProductService interface {
	CreateConfigurableProduct(product request.ConfigurableProductRequest)
	FindBySku(sku string) response.ConfigurableProductResponse
}
