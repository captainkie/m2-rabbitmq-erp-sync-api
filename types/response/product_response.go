package response

type ConfigurableProductResponse struct {
	ID            int    `json:"ID"`
	Sku           string `json:"Sku"`
	FirstChildSku string `json:"FirstChildSku"`
	Created       int64  `json:"Created"`
	Updated       int64  `json:"Updated"`
}

type ImageDataResponse struct {
	PIC_FILE  string
	PIC_FILE2 string
	PIC_FILE3 string
	PIC_FILE4 string
	PIC_FILE5 string
}
