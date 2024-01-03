package request

type CreateAttributeOptionRequest struct {
	Option struct {
		Label       string `json:"label"`
		Value       string `json:"value"`
		StoreLabels []struct {
			StoreID int    `json:"store_id"`
			Label   string `json:"label"`
		} `json:"store_labels"`
	} `json:"option"`
}

type StockItem struct {
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
}

type StockItemsRequest struct {
	StockItem `json:"stock_item"`
}

type CreateMediaContentRequest struct {
	Base64EncodedData string `json:"base64_encoded_data"`
	Type              string `json:"type"`
	Name              string `json:"name"`
}

type CreateMediaEntryRequest struct {
	MediaType string                    `json:"media_type"`
	Label     string                    `json:"label"`
	Position  int                       `json:"position"`
	Disabled  bool                      `json:"disabled"`
	Content   CreateMediaContentRequest `json:"content"`
}

type UpdateMediaEntryRequest struct {
	ID        int                       `json:"id"`
	MediaType string                    `json:"media_type"`
	Label     string                    `json:"label"`
	Position  int                       `json:"position"`
	Disabled  bool                      `json:"disabled"`
	Types     []string                  `json:"types"`
	File      string                    `json:"file"`
	Content   CreateMediaContentRequest `json:"content"`
}

type CreateMediaRequest struct {
	Entry CreateMediaEntryRequest `json:"entry"`
}

type UpdateMediaRequest struct {
	Entry UpdateMediaEntryRequest `json:"entry"`
}
