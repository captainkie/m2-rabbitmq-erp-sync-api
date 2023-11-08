package payload

type CustomAttribute struct {
	AttributeCode string      `json:"attribute_code"`
	Value         interface{} `json:"value"`
}

type ConfigurableProductPayload struct {
	Id                  int        `json:"id"`
	Sku                 string     `json:"sku"`
	Name                string     `json:"name"`
	AttributeSetId      int        `json:"attribute_set_id"`
	Price               float64    `json:"price"`
	Status              int        `json:"status"`
	Visibility          int        `json:"visibility"`
	TypeId              string     `json:"type_id"`
	CreatedAt           string     `json:"created_at"`
	UpdatedAt           string     `json:"updated_at"`
	Weight              int        `json:"weight"`
	ExtensionAttributes struct{}   `json:"extension_attributes"`
	ProductLinks        []struct{} `json:"product_links"`
	Options             []struct{} `json:"options"`
	MediaGalleryEntries []struct{} `json:"media_gallery_entries"`
	TierPrices          []struct{} `json:"tier_prices"`
	CustomAttributes    []struct{} `json:"custom_attributes"`
}

type SimpleProductPayload struct {
	Id                  int     `json:"id"`
	Sku                 string  `json:"sku"`
	Name                string  `json:"name"`
	AttributeSetId      int     `json:"attribute_set_id"`
	Price               float64 `json:"price"`
	Status              int     `json:"status"`
	Visibility          int     `json:"visibility"`
	TypeId              string  `json:"type_id"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	Weight              int     `json:"weight"`
	ExtensionAttributes struct {
		WebsiteIds    []int `json:"website_ids"`
		CategoryLinks []struct {
			Position   int    `json:"position"`
			CategoryId string `json:"category_id"`
		} `json:"category_links"`
		StockItem struct {
			ItemId                         int    `json:"item_id"`
			ProductId                      int    `json:"product_id"`
			StockId                        int    `json:"stock_id"`
			Qty                            int    `json:"qty"`
			IsInStock                      bool   `json:"is_in_stock"`
			IsQtyDecimal                   bool   `json:"is_qty_decimal"`
			ShowDefaultNotificationMessage bool   `json:"show_default_notification_message"`
			UseConfigMinQty                bool   `json:"use_config_min_qty"`
			MinQty                         int    `json:"min_qty"`
			UseConfigMinSaleQty            int    `json:"use_config_min_sale_qty"`
			MinSaleQty                     int    `json:"min_sale_qty"`
			UseConfigMaxSaleQty            bool   `json:"use_config_max_sale_qty"`
			MaxSaleQty                     int    `json:"max_sale_qty"`
			UseConfigBackorders            bool   `json:"use_config_backorders"`
			Backorders                     int    `json:"backorders"`
			UseConfigNotifyStockQty        bool   `json:"use_config_notify_stock_qty"`
			NotifyStockQty                 int    `json:"notify_stock_qty"`
			UseConfigQtyIncrements         bool   `json:"use_config_qty_increments"`
			QtyIncrements                  int    `json:"qty_increments"`
			UseConfigEnableQtyInc          bool   `json:"use_config_enable_qty_inc"`
			EnableQtyIncrements            bool   `json:"enable_qty_increments"`
			UseConfigManageStock           bool   `json:"use_config_manage_stock"`
			ManageStock                    bool   `json:"manage_stock"`
			LowStockDate                   string `json:"low_stock_date"`
			IsDecimalDivided               bool   `json:"is_decimal_divided"`
			StockStatusChangedAuto         int    `json:"stock_status_changed_auto"`
		} `json:"stock_item"`
	} `json:"extension_attributes"`
	ProductLinks        []struct{}        `json:"product_links"`
	Options             []struct{}        `json:"options"`
	MediaGalleryEntries []struct{}        `json:"media_gallery_entries"`
	TierPrices          []struct{}        `json:"tier_prices"`
	CustomAttributes    []CustomAttribute `json:"custom_attributes"`
}

type ProductAttributeOptionPayload []struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
