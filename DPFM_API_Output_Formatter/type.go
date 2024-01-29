package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       	string      `json:"connection_key"`
	Result              	bool        `json:"result"`
	RedisKey            	string      `json:"redis_key"`
	Filepath            	string      `json:"filepath"`
	APIStatusCode       	int         `json:"api_status_code"`
	RuntimeSessionID    	string      `json:"runtime_session_id"`
	BusinessPartnerID   	*int        `json:"business_partner"`
	ServiceLabel        	string      `json:"service_label"`
	APIType             	string      `json:"api_type"`
	Message             	interface{} `json:"message"`
	APISchema           	string      `json:"api_schema"`
	Accepter            	[]string    `json:"accepter"`
	Deleted             	bool        `json:"deleted"`
	APIProcessingResult 	*bool       `json:"api_processing_result"`
	APIProcessingError  	string      `json:"api_processing_error"`
	DeliveryInstructionPdf	string      `json:"delivery_instruction_pdf"`
}

type Message struct {
	Header     []Header   `json:"Header"`
}

type Header struct {
	DeliveryDocument                       int      `json:"DeliveryDocument"`
	DeliveryDocumentDate                   string   `json:"DeliveryDocumentDate"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverToPartyName					   string   `json:"DeliverToPartyName"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverFromPartyName				   string   `json:"DeliverFromPartyName"`
	IsExportImport                         *bool    `json:"IsExportImport"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	PlannedGoodsIssueDate                  string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                string   `json:"PlannedGoodsReceiptTime"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit"`
	Incoterms                              *string  `json:"Incoterms"`
	Items				  				   []Items `json:"Items"`
}

type Items struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  string   `json:"DeliveryDocumentItemCategory"`
	Product                                       string   `json:"Product"`
	ProductSpecification                          *string  `json:"ProductSpecification"`
	SizeOrDimensionText                           *string  `json:"SizeOrDimensionText"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueQuantity                     float32  `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                float32  `json:"PlannedGoodsIssueQtyInBaseUnit"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	PlannedGoodsIssueDate                         string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       string   `json:"PlannedGoodsReceiptTime"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
}
