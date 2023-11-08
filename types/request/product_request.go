package request

type ConfigurableProductRequest struct {
	Sku           string `validate:"required" json:"sku"`
	FirstChildSku string `validate:"required" json:"first_child_sku"`
}
