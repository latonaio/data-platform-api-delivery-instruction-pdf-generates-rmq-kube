package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey          string      `json:"connection_key"`
	Result                 bool        `json:"result"`
	RedisKey               string      `json:"redis_key"`
	Filepath               string      `json:"filepath"`
	APIStatusCode          int         `json:"api_status_code"`
	RuntimeSessionID       string      `json:"runtime_session_id"`
	BusinessPartnerID      *int        `json:"business_partner"`
	ServiceLabel           string      `json:"service_label"`
	APIType                string      `json:"api_type"`
	Message                interface{} `json:"message"`
	APISchema              string      `json:"api_schema"`
	Accepter               []string    `json:"accepter"`
	Deleted                bool        `json:"deleted"`
	APIProcessingResult    *bool       `json:"api_processing_result"`
	APIProcessingError     string      `json:"api_processing_error"`
	DeliveryInstructionPdf string      `json:"delivery_instruction_pdf"`
	MountPath              *string     `json:"mount_path"`
}

type Message struct {
	Header []Header `json:"Header"`
}

type Header struct {
	DeliveryDocument                    string   `json:"DeliveryDocument"`
	DeliveryDocumentDate                string   `json:"DeliveryDocumentDate"`
	DeliverToParty                      int      `json:"DeliverToParty"`
	DeliverToPartyName                  string   `json:"DeliverToPartyName"`
	DeliverToPlant                      string   `json:"DeliverToPlant"`
	DeliverToPlantName                  string   `json:"DeliverToPlantName"`
	DeliverFromParty                    int      `json:"DeliverFromParty"`
	DeliverFromPartyName                string   `json:"DeliverFromPartyName"`
	DeliverFromPlant                    string   `json:"DeliverFromPlant"`
	DeliverFromPlantName                string   `json:"DeliverFromPlantName"`
	IsExportImport                      *bool    `json:"IsExportImport"`
	OrderID                             *string  `json:"OrderID"`
	OrderItem                           *int     `json:"OrderItem"`
	Contract                            *int     `json:"Contract"`
	ContractItem                        *int     `json:"ContractItem"`
	Project                             *int     `json:"Project"`
	WBSElement                          *int     `json:"WBSElement"`
	WBSElementDescription               *string  `json:"WBSElementDescription"`
	ProductionOrder                     *int     `json:"ProductionOrder"`
	ProductionOrderItem                 *int     `json:"ProductionOrderItem"`
	PlannedGoodsIssueDate               string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime               string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate             string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime             string   `json:"PlannedGoodsReceiptTime"`
	HeaderGrossWeight                   *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                     *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                    *string  `json:"HeaderWeightUnit"`
	Incoterms                           *string  `json:"Incoterms"`
	IncotermsText                       *string  `json:"IncotermsText"`
	TotalPlannedGoodsIssueQuantity      string   `json:"TotalPlannedGoodsIssueQuantity"`
	TotalPlannedGoodsIssueQtyInBaseUnit string   `json:"TotalPlannedGoodsIssueQtyInBaseUnit"`
	Items                               []Items  `json:"Items"`
}

type Items struct {
	DeliveryDocument                 int      `json:"DeliveryDocument"`
	DeliveryDocumentItem             int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory     string   `json:"DeliveryDocumentItemCategory"`
	Product                          string   `json:"Product"`
	ProductSpecification             *string  `json:"ProductSpecification"`
	SizeOrDimensionText              *string  `json:"SizeOrDimensionText"`
	DeliveryDocumentItemText         *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer  *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller *string  `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueQuantity        float32  `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit   float32  `json:"PlannedGoodsIssueQtyInBaseUnit"`
	BaseUnit                         string   `json:"BaseUnit"`
	DeliveryUnit                     string   `json:"DeliveryUnit"`
	PlannedGoodsIssueDate            string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime            string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate          string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime          string   `json:"PlannedGoodsReceiptTime"`
	ItemWeightUnit                   *string  `json:"ItemWeightUnit"`
	ItemNetWeight                    *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                  *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                 *float32 `json:"ProductNetWeight"`
	Project                          *int     `json:"Project"`
	WBSElement                       *int     `json:"WBSElement"`
	OrderID                          *int     `json:"OrderID"`
	OrderItem                        *int     `json:"OrderItem"`
}
