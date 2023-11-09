package response

type CreateConnectionResponse struct {
	ID                int    `json:"ID"`
	MessageCode       string `json:"MessageCode"`
	MessageDesc       string `json:"MessageDesc"`
	TotalRecordAdd    string `json:"TotalRecordAdd"`
	TotalRecordUpdate string `json:"TotalRecordUpdate"`
	TotalRecordStock  string `json:"TotalRecordStock"`
	TotalRecordStore  string `json:"TotalRecordStore"`
	Created           int64  `json:"Created"`
	Updated           int64  `json:"Updated"`
}

type CreateDailySalesResponse struct {
	DocNo   string `json:"DOC_NO"`
	Status  string `json:"STATUS"`
	Message string `json:"MESSAGE"`
}
