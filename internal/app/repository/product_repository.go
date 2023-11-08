package repository

import model "github.com/captainkie/websync-api/internal/app/models"

type ProductRepository interface {
	ConfigurableProduct(products model.ConfigurableProducts)
	FindBySku(sku string) (model.ConfigurableProducts, error)
}
