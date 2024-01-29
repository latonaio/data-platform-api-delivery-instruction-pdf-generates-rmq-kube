package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey       	string         `json:"connection_key"`
	Result              	bool           `json:"result"`
	RedisKey            	string         `json:"redis_key"`
	Filepath            	string         `json:"filepath"`
	APIStatusCode       	int            `json:"api_status_code"`
	RuntimeSessionID    	string         `json:"runtime_session_id"`
	BusinessPartnerID   	*int           `json:"business_partner"`
	ServiceLabel        	string         `json:"service_label"`
	Header		        	[]Header 	   `json:"DeliveryDocumentHeaderWithItem"`
	Items		    		[]Items		   `json:"DeliveryDocumentItem"`
	APISchema           	string         `json:"api_schema"`
	Accepter            	[]string       `json:"accepter"`
	Deleted             	bool           `json:"deleted"`
	DocData             	string         `json:"doc_data"`
	APIProcessingResult 	*bool          `json:"api_processing_result"`
	APIProcessingError  	string         `json:"api_processing_error"`
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
