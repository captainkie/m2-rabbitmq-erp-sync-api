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
	// DOC_NO is the unique identifier for the order. (หมายเลขคำสั่งซื้อ)
	// It required && must be between 1 and 50 characters.
	DocNo string `validate:"required,min=1,max=50" json:"DOC_NO"`
	// POS_ENTITY is required && fixed value = "XXXXX" (รหัสกลุ่มข้อมูล)
	PosEntity string `validate:"required,min=1" json:"POS_ENTITY"`
	// DOC_CODE is required && fixed value = "DS" (ประเภทเอกสาร)
	DocCode string `validate:"required,min=1" json:"DOC_CODE"`
	// DOC_DATE is required && format = "DD/MM/YYYY" (วันที่เอกสาร)
	DocDate string `validate:"required,min=1" json:"DOC_DATE"`
	// REF_NO is required && fix value space = " " (เลขที่อ้างอิง)
	RefNo string `validate:"required,min=0,max=50" json:"REF_NO"`
	// CUST_NAME is required && must be between 1 and 50 characters. (ชื่อลูกค้า)
	// It must be `${firstname} ${lastname}`
	CustomerName string `validate:"required,min=1,max=50" json:"CUST_NAME"`
	// ADDRESS1 is required && must be between 1 and 50 characters. (ที่อยู่)
	// It must be `${street_adress1}`
	Address1 string `validate:"required,min=1,max=50" json:"ADDRESS1"`
	// ADDRESS2 is required && must be between 1 and 50 characters. (อำเภอ)
	// It must be `${city}`
	Address2 string `validate:"required,min=1,max=50" json:"ADDRESS2"`
	// PROVINCE is required && must be between 1 and 50 characters. (จังหวัด)
	// It must be `${state}` or `${province}`
	Province string `validate:"required,min=1,max=50" json:"PROVINCE"`
	// POSTCODE is required && must be between 1 and 50 characters. (รหัสไปรษณีย์)
	// It must be `${zipcode}` or `${postcode}`
	Postcode string `validate:"required,min=1,max=5" json:"POSTCODE"`
	// TOT_AMT is required && must be decimal 2 digit (จำนวนเงินรวมภาษี)
	// Example: 100.00
	TotAmt float64 `validate:"min=0.00" json:"TOT_AMT"`
	// TOT_VATAMT is required && must be decimal 2 digit (จำนวนเงินภาษี)
	// It fix value = 0.00
	// Example: 100.00
	TotVatamt float64 `validate:"min=0.00" json:"TOT_VATAMT"`
	// TOT_NETAMT is required && must be decimal 2 digit (จำนวนเงินไม่รวมภาษี)
	// Example: 100.00
	TotNetamt float64 `validate:"min=0.00" json:"TOT_NETAMT"`
	// TOT_DISCAMT is required && must be decimal 2 digit (จำนวนเงินส่วนลด)
	// Example: 100.00
	TotDiscamt float64 `validate:"min=0.00" json:"TOT_DISCAMT"`
	// TOT_SUBAMT is required && must be decimal 2 digit (จำนวนเงินก่อนหักส่วนลด)
	// Example: 100.00
	TotSubamt float64 `validate:"min=0.00" json:"TOT_SUBAMT"`
	// VAT_RATE is required && must be decimal 2 digit (อัตราภาษีมูลค่าเพิ่ม)
	// It fix value = 7.00
	// Example: 7.00
	VatRate float64 `validate:"required,min=1" json:"VAT_RATE"`
	// PAY_CODE is required. (รหัสประเภทการจ่ายเงิน)
	// For Credit Card, Fix value = "VISA1"
	// For QR Code,  Fix value = "QR"
	PayCode string `validate:"required,min=1" json:"PAY_CODE"`
	// PAY_DESC is required. (รายละเอียดประเภทการจ่ายเงิน)
	// For Credit Card, Fix value = "VIS1 - VISA ( EDC )"
	// For QR Code,  Fix value = "QR Code"
	PayDesc string `validate:"required,min=1" json:"PAY_DESC"`
	// REMARK1 is required && must be between 1 and 50 characters. (หมายเลขโทรศัพท์)
	Remark1 string `validate:"required,min=1,max=50" json:"REMARK1"`
	// REMARK2 is required && must be between 0 and 50 characters. (อื่นๆ)
	// It fix value space = " "
	Remark2 string `validate:"required,min=0,max=50" json:"REMARK2"`
	// USER_CODE is required && must be between 1 and 50 characters. (ผู้บันทึก)
	// It fix value = "ECOMERCE"
	UserCode string `validate:"required" json:"USER_CODE"`
	// SYS_DATE is required && format = "DD/MM/YYYY" (วันที่บันทึก)
	SysDate string `validate:"required" json:"SYS_DATE"`
	// DOC_STATUS is required && fix value = "NP" (สถานะเอกสาร)
	DocStatus string `validate:"required" json:"DOC_STATUS"`
	// DISC1 is required && must be decimal 2 digit (อัตตาส่วนลด)
	// It must be `(${discount_total} / ${base_total}) * 100`
	// Example: 100.00
	Disc1 float64 `validate:"min=0.00" json:"DISC1"`
	// CUST_CODE is required && fix value = "Z00000" (รหัสลูกค้า)
	CustCode string `validate:"required" json:"CUST_CODE"`
	// NAME_TITLE is required && must be between 0 and 50 characters. (คำนำหน้าชื่อ)
	// It must be `${title}`
	// If no data Fix value space = " "
	NameTitle string `validate:"required,min=0,max=50" json:"NAME_TITLE"`
	// BRANCH is required && must be between 0 and 50 characters. ( สาขาของบริษัท - สำหรับใบกำกับภาษีแบบเต็ม)
	// It must be `${company}`
	// If no data Fix value space = " "
	Branch string `validate:"required,min=0,max=50" json:"BRANCH"`
	// TAX_ID is required && must be between 0 and 50 characters. (เลขที่ผู้เสียภาษี - สำหรับใบกำกับภาษีแบบเต็ม)
	// It must be `${tax_id}`
	// If no data Fix value space = " "
	TaxId string `validate:"required,min=0,max=50" json:"TAX_ID"`
	// Detail is array od products (รายละเอียดสินค้า)
	// If has shipping price, add shipping product to array (หากมาค่าขนส่งให้เพิ่มสินค้าค่าขนส่งเข้าไปใน array)
	// Example data for shipping product (ตัวอย่างข้อมูลสินค้าค่าขนส่ง)
	// `{`
	// `    "Item": 1,`
	// `    "BAR_CODE": "8852012882769",`
	// `    "PROD_CODE": "ZZSV0   00TRANSPORT",`
	// `    "PROD_DESC": "ค่าขนส่ง (Transportation Charge)",`
	// `    "UOM_CODE": "PCS",`
	// `    "UNIT_PRICE": 500.00,`
	// `    "DISC_RATE": 0.00,`
	// `    "DISC_AMT": 0.00,`
	// `    "VAT_AMT": 0.00,`
	// `    "NET_AMT": 500.00,`
	// `    "AMT": 500.00,`
	// `    "QTY": 1.00,`
	// `    "SALE_PD": "Y",`
	// `    "PROD_STATUS": "N"`
	// `}`
	Detail []struct {
		// ITEM is required. (ลำดับที่)
		// It must be `${index}`
		// Example: 1
		Item int `validate:"required" json:"ITEM"`
		// BAR_CODE is required. (บาร์โค้ดสินค้า)
		// It must be `${service_bar_code}`
		BarCode string `validate:"required" json:"BAR_CODE"`
		// PROD_CODE is required. (รหัสสินค้า)
		// It must be `${sku}`
		ProdCode string `validate:"required" json:"PROD_CODE"`
		// PROD_DESC is required && must be between 1 and 50 characters. (ชื่อสินค้า)
		// It must be `${name}`
		ProdDesc string `validate:"required,min=1,max=50" json:"PROD_DESC"`
		// UOM_CODE is required. (หน่วยนับ)
		// It must be `${service_uom_code}`
		UomCode string `validate:"required" json:"UOM_CODE"`
		// UNIT_PRICE is required && must be decimal 2 digit (ราคาต่อหน่วย)
		// It must be `${product->getPrice()}`
		// Example: 100.00
		UnitPrice float64 `validate:"required" json:"UNIT_PRICE"`
		// DISC_RATE is required && must be decimal 2 digit (อัตราส่วนลด)
		// It must be `(${discount} / ${product->getPrice()}) * 100`
		// Example: 100.00
		DiscRate float64 `validate:"required" json:"DISC_RATE"`
		// DISC_AMT is required && must be decimal 2 digit (จำนวนเงินส่วนลด * item qty)
		// It must be `${discount} * ${item->getQtyOrdered()}`
		// Example: 100.00
		DiscAmt float64 `validate:"required" json:"DISC_AMT"`
		// VAT_AMT is required && must be decimal 2 digit (ภาษี)
		// It Fix value = 0.00
		// Example: 0.00
		VatAmt float64 `validate:"required" json:"VAT_AMT"`
		// NET_AMT is required && must be decimal 2 digit (จำนวนเงินไม่รวมภาษี)
		// It must be `${item->getBaseRowTotal()}`
		// Example: 100.00
		NetAmt float64 `validate:"required" json:"NET_AMT"`
		// AMT is required && must be decimal 2 digit (จำนวนเงินรวมภาษี)
		// It must be `${item->getBaseRowTotal()}`
		// Example: 100.00
		Amt float64 `validate:"required" json:"AMT"`
		// QTY is required && must be decimal 2 digit (จำนวน)
		// It must be `${item->getQtyOrdered()}`
		// Example: 100.00
		Qty float64 `validate:"required" json:"QTY"`
		// SALE_PD is required &&  fix value = "Y" (รหัสพนักงานขาย)
		SalePd string `validate:"required,min=0,max=50" json:"SALE_PD"`
		// PROD_STATUS is required &&  fix value = "Y" (สถานะสินค้า)
		ProdStatus string `validate:"required" json:"PROD_STATUS"`
	} `validate:"required" json:"Detail"`
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
