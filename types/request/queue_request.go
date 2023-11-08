package request

// ========= ERP PRODUCTION =========
type CreateConnectionRequest struct {
	MessageCode       string `json:"MessageCode"`
	MessageDesc       string `json:"MessageDesc"`
	TotalRecordAdd    string `json:"TotalRecordAdd"`
	TotalRecordUpdate string `json:"TotalRecordUpdate"`
	TotalRecordStock  string `json:"TotalRecordStock"`
	TotalRecordStore  string `json:"TotalRecordStore"`
}

type AddUpdateProductRequest struct {
	TRANS_ID          string `json:"TRANS_ID"`
	TYPE              string `json:"TYPE"`
	PROD_CODE         string `json:"PROD_CODE"`
	PROD_TNAME        string `json:"PROD_TNAME"`
	PROD_ENAME        string `json:"PROD_ENAME"`
	UOM_CODE          string `json:"UOM_CODE"`
	BAR_CODE          string `json:"BAR_CODE"`
	PDGRP_CODE        string `json:"PDGRP_CODE"`
	PDBRND_CODE       string `json:"PDBRND_CODE"`
	PDTYPE_CODE       string `json:"PDTYPE_CODE"`
	PDDSGN_CODE       string `json:"PDDSGN_CODE"`
	PDSIZE_CODE       string `json:"PDSIZE_CODE"`
	PDCOLOR_CODE      string `json:"PDCOLOR_CODE"`
	PDMISC_CODE       string `json:"PDMISC_CODE"`
	PDGRP_DESC        string `json:"PDGRP_DESC"`
	PDBRND_DESC       string `json:"PDBRND_DESC"`
	PDTYPE_DESC       string `json:"PDTYPE_DESC"`
	PDDSGN_DESC       string `json:"PDDSGN_DESC"`
	PDCOLOR_DESC      string `json:"PDCOLOR_DESC"`
	PDSIZE_DESC       string `json:"PDSIZE_DESC"`
	PDMISC_DESC       string `json:"PDMISC_DESC"`
	REF_NO            string `json:"REF_NO"`
	PROD_ST           string `json:"PROD_ST"`
	PROD_CLASS        string `json:"PROD_CLASS"`
	PROD_LINE         string `json:"PROD_LINE"`
	PROD_TYPE         string `json:"PROD_TYPE"`
	UNIT_PRICE        string `json:"UNIT_PRICE"`
	PDMODEL_CODE      string `json:"PDMODEL_CODE"`
	PDMODEL_DESC      string `json:"PDMODEL_DESC"`
	COLOR1            string `json:"COLOR1"`
	COLOR2            string `json:"COLOR2"`
	MATERIAL1         string `json:"MATERIAL1"`
	MATERIAL2         string `json:"MATERIAL2"`
	DIMENSION         string `json:"DIMENSION"`
	PRODUCT_TYPE      string `json:"PRODUCT_TYPE"`
	SHORT_DESC_TH     string `json:"SHORT_DESC_TH"`
	SHORT_DESC_EN     string `json:"SHORT_DESC_EN"`
	DESC_TH           string `json:"DESC_TH"`
	DESC_EN           string `json:"DESC_EN"`
	MATERIAL1_TH      string `json:"MATERIAL1_TH"`
	MATERIAL2_TH      string `json:"MATERIAL2_TH"`
	DIMENSION_DESC_TH string `json:"DIMENSION_DESC_TH"`
	DIMENSION_DESC_EN string `json:"DIMENSION_DESC_EN"`
	UOM_TH            string `json:"UOM_TH"`
	UOM_EN            string `json:"UOM_EN"`
	WEIGHT            string `json:"WEIGHT"`
	PDGRP_TH          string `json:"PDGRP_TH"`
	PDTYPE_TH         string `json:"PDTYPE_TH"`
	PDDSGN_TH         string `json:"PDDSGN_TH"`
	PDCOLOR_TH        string `json:"PDCOLOR_TH"`
	PDSIZE_TH         string `json:"PDSIZE_TH"`
	PIC_FILE          string `json:"PIC_FILE"`
	PIC_FILE2         string `json:"PIC_FILE2"`
	PIC_FILE3         string `json:"PIC_FILE3"`
	PIC_FILE4         string `json:"PIC_FILE4"`
	PIC_FILE5         string `json:"PIC_FILE5"`
	PDNAME_TH         string `json:"PDNAME_TH"`
	PDNAME_EN         string `json:"PDNAME_EN"`
	STOCK_QTY         string `json:"STOCK_QTY"`
	SEND_FLAG         string `json:"SEND_FLAG"`
	ErrorMessage      string `json:"ErrorMessage"`
}

type UpdateStockRequest struct {
	TRANS_ID     int    `json:"TRANS_ID"`
	RecordType   string `json:"RecordType"`
	PROD_CODE    string `json:"PROD_CODE"`
	StockQty     string `json:"StockQty"`
	SEND_FLAG    string `json:"SEND_FLAG"`
	ErrorMessage string `json:"ErrorMessage"`
}

type UpdateStoreRequest struct {
	TRANS_ID     int    `json:"TRANS_ID"`
	RecordType   string `json:"RecordType"`
	PROD_CODE    string `json:"PROD_CODE"`
	StockQty     string `json:"StockQty"`
	SEND_FLAG    string `json:"SEND_FLAG"`
	ErrorMessage string `json:"ErrorMessage"`
}

type CreatePostflagRequest struct {
	TransactionId int    `validate:"required" json:"TransactionId"`
	FlagStatus    string `validate:"required" json:"FlagStatus"`
	ErrMsg        string `validate:"required" json:"ErrMsg"`
}

type CreateDailySalesRequest struct {
	OrderID string `validate:"required,min=1,max=100" json:"OrderID"`
}

// ========= STRAPI DEVELOPMENT =========
type StrapiConnectionRequest struct {
	Data struct {
		Attributes struct {
			Body struct {
				MessageCode       string `json:"MessageCode"`
				MessageDesc       string `json:"MessageDesc"`
				TotalRecordAdd    string `json:"TotalRecordAdd"`
				TotalRecordUpdate string `json:"TotalRecordUpdate"`
				TotalRecordStock  string `json:"TotalRecordStock"`
				TotalRecordStore  string `json:"TotalRecordStore"`
			} `json:"body"`
		} `json:"attributes"`
	} `json:"data"`
}

type StrapiAddUpdateProductRequest struct {
	Data struct {
		Attributes struct {
			Body []struct {
				TRANS_ID          string `json:"TRANS_ID"`
				TYPE              string `json:"TYPE"`
				PROD_CODE         string `json:"PROD_CODE"`
				PROD_TNAME        string `json:"PROD_TNAME"`
				PROD_ENAME        string `json:"PROD_ENAME"`
				UOM_CODE          string `json:"UOM_CODE"`
				BAR_CODE          string `json:"BAR_CODE"`
				PDGRP_CODE        string `json:"PDGRP_CODE"`
				PDBRND_CODE       string `json:"PDBRND_CODE"`
				PDTYPE_CODE       string `json:"PDTYPE_CODE"`
				PDDSGN_CODE       string `json:"PDDSGN_CODE"`
				PDSIZE_CODE       string `json:"PDSIZE_CODE"`
				PDCOLOR_CODE      string `json:"PDCOLOR_CODE"`
				PDMISC_CODE       string `json:"PDMISC_CODE"`
				PDGRP_DESC        string `json:"PDGRP_DESC"`
				PDBRND_DESC       string `json:"PDBRND_DESC"`
				PDTYPE_DESC       string `json:"PDTYPE_DESC"`
				PDDSGN_DESC       string `json:"PDDSGN_DESC"`
				PDCOLOR_DESC      string `json:"PDCOLOR_DESC"`
				PDSIZE_DESC       string `json:"PDSIZE_DESC"`
				PDMISC_DESC       string `json:"PDMISC_DESC"`
				REF_NO            string `json:"REF_NO"`
				PROD_ST           string `json:"PROD_ST"`
				PROD_CLASS        string `json:"PROD_CLASS"`
				PROD_LINE         string `json:"PROD_LINE"`
				PROD_TYPE         string `json:"PROD_TYPE"`
				UNIT_PRICE        string `json:"UNIT_PRICE"`
				PDMODEL_CODE      string `json:"PDMODEL_CODE"`
				PDMODEL_DESC      string `json:"PDMODEL_DESC"`
				COLOR1            string `json:"COLOR1"`
				COLOR2            string `json:"COLOR2"`
				MATERIAL1         string `json:"MATERIAL1"`
				MATERIAL2         string `json:"MATERIAL2"`
				DIMENSION         string `json:"DIMENSION"`
				PRODUCT_TYPE      string `json:"PRODUCT_TYPE"`
				SHORT_DESC_TH     string `json:"SHORT_DESC_TH"`
				SHORT_DESC_EN     string `json:"SHORT_DESC_EN"`
				DESC_TH           string `json:"DESC_TH"`
				DESC_EN           string `json:"DESC_EN"`
				MATERIAL1_TH      string `json:"MATERIAL1_TH"`
				MATERIAL2_TH      string `json:"MATERIAL2_TH"`
				DIMENSION_DESC_TH string `json:"DIMENSION_DESC_TH"`
				DIMENSION_DESC_EN string `json:"DIMENSION_DESC_EN"`
				UOM_TH            string `json:"UOM_TH"`
				UOM_EN            string `json:"UOM_EN"`
				WEIGHT            string `json:"WEIGHT"`
				PDGRP_TH          string `json:"PDGRP_TH"`
				PDTYPE_TH         string `json:"PDTYPE_TH"`
				PDDSGN_TH         string `json:"PDDSGN_TH"`
				PDCOLOR_TH        string `json:"PDCOLOR_TH"`
				PDSIZE_TH         string `json:"PDSIZE_TH"`
				PIC_FILE          string `json:"PIC_FILE"`
				PIC_FILE2         string `json:"PIC_FILE2"`
				PIC_FILE3         string `json:"PIC_FILE3"`
				PIC_FILE4         string `json:"PIC_FILE4"`
				PIC_FILE5         string `json:"PIC_FILE5"`
				PDNAME_TH         string `json:"PDNAME_TH"`
				PDNAME_EN         string `json:"PDNAME_EN"`
				STOCK_QTY         string `json:"STOCK_QTY"`
				SEND_FLAG         string `json:"SEND_FLAG"`
				ErrorMessage      string `json:"ErrorMessage"`
			} `json:"body"`
		} `json:"attributes"`
	} `json:"data"`
}

type StrapiUpdateStockRequest struct {
	Data struct {
		Attributes struct {
			Body []struct {
				TRANS_ID     int    `json:"TRANS_ID"`
				RecordType   string `json:"RecordType"`
				PROD_CODE    string `json:"PROD_CODE"`
				StockQty     string `json:"StockQty"`
				SEND_FLAG    string `json:"SEND_FLAG"`
				ErrorMessage string `json:"ErrorMessage"`
			} `json:"body"`
		} `json:"attributes"`
	} `json:"data"`
}

type StrapiUpdateStoreRequest struct {
	Data struct {
		Attributes struct {
			Body []struct {
				TRANS_ID     int    `json:"TRANS_ID"`
				RecordType   string `json:"RecordType"`
				PROD_CODE    string `json:"PROD_CODE"`
				StockQty     string `json:"StockQty"`
				SEND_FLAG    string `json:"SEND_FLAG"`
				ErrorMessage string `json:"ErrorMessage"`
			} `json:"body"`
		} `json:"attributes"`
	} `json:"data"`
}

type StrapiDailySalesRequest struct {
	OrderID string `validate:"required,min=1,max=100" json:"OrderID"`
}
