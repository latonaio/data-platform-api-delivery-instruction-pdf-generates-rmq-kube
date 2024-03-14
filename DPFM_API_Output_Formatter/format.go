package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-instruction-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	"fmt"
	"strconv"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputItems []dpfm_api_input_reader.Items,
) *Header {
	pm := &Header{}

	var items []Items
	for _, dataItems := range inputItems {
		if headerData.DeliveryDocument == dataItems.DeliveryDocument {
			items = append(items, Items{
				DeliveryDocument:                 dataItems.DeliveryDocument,
				DeliveryDocumentItem:             dataItems.DeliveryDocumentItem,
				Product:                          dataItems.Product,
				ProductSpecification:             dataItems.ProductSpecification,
				SizeOrDimensionText:              dataItems.SizeOrDimensionText,
				DeliveryDocumentItemText:         dataItems.DeliveryDocumentItemText,
				DeliveryDocumentItemTextByBuyer:  dataItems.DeliveryDocumentItemTextByBuyer,
				DeliveryDocumentItemTextBySeller: dataItems.DeliveryDocumentItemTextBySeller,
				PlannedGoodsIssueQuantity:        dataItems.PlannedGoodsIssueQuantity,
				PlannedGoodsIssueQtyInBaseUnit:   dataItems.PlannedGoodsIssueQtyInBaseUnit,
				BaseUnit:                         dataItems.BaseUnit,
				DeliveryUnit:                     dataItems.DeliveryUnit,
				PlannedGoodsIssueDate:            dataItems.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:            dataItems.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:          dataItems.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:          dataItems.PlannedGoodsReceiptTime,
				ItemWeightUnit:                   dataItems.ItemWeightUnit,
				ItemNetWeight:                    dataItems.ItemNetWeight,
				ItemGrossWeight:                  dataItems.ItemGrossWeight,
				Project:                          dataItems.Project,
				WBSElement:                       dataItems.WBSElement,
			})
		}
	}

	orderID := strconv.Itoa(*headerData.OrderID)

	pm = &Header{
		DeliveryDocument:                    fmt.Sprintf("%d", headerData.DeliveryDocument),
		DeliveryDocumentDate:                headerData.DeliveryDocumentDate,
		DeliverToParty:                      headerData.DeliverToParty,
		DeliverToPartyName:                  headerData.DeliverToPartyName,
		DeliverToPlant:                      headerData.DeliverToPlant,
		DeliverToPlantName:                  headerData.DeliverToPlantName,
		DeliverFromParty:                    headerData.DeliverFromParty,
		DeliverFromPartyName:                headerData.DeliverFromPartyName,
		DeliverFromPlant:                    headerData.DeliverFromPlant,
		DeliverFromPlantName:                headerData.DeliverFromPlantName,
		IsExportImport:                      headerData.IsExportImport,
		OrderID:                             &orderID,
		OrderItem:                           headerData.OrderItem,
		Project:                             headerData.Project,
		WBSElement:                          headerData.WBSElement,
		WBSElementDescription:               headerData.WBSElementDescription,
		ProductionOrder:                     headerData.ProductionOrder,
		ProductionOrderItem:                 headerData.ProductionOrderItem,
		PlannedGoodsIssueDate:               headerData.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:               headerData.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:             headerData.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:             headerData.PlannedGoodsReceiptTime,
		HeaderGrossWeight:                   headerData.HeaderGrossWeight,
		HeaderNetWeight:                     headerData.HeaderNetWeight,
		HeaderWeightUnit:                    headerData.HeaderWeightUnit,
		Incoterms:                           headerData.Incoterms,
		Items:                               items,
		TotalPlannedGoodsIssueQuantity:      calculateTotalPlannedGoodsIssueQuantity(inputItems),
		TotalPlannedGoodsIssueQtyInBaseUnit: calculatePlannedGoodsIssueQtyInBaseUnit(inputItems),
	}

	return pm
}

func calculateTotalPlannedGoodsIssueQuantity(
	inputItems []dpfm_api_input_reader.Items,
) string {
	var totalPlannedGoodsIssueQuantity float32

	for _, dataItems := range inputItems {
		totalPlannedGoodsIssueQuantity += dataItems.PlannedGoodsIssueQuantity
	}

	return fmt.Sprint(totalPlannedGoodsIssueQuantity)
}

func calculatePlannedGoodsIssueQtyInBaseUnit(
	inputItems []dpfm_api_input_reader.Items,
) string {
	var totalPlannedGoodsIssueQtyInBaseUnit float32

	for _, dataItems := range inputItems {
		totalPlannedGoodsIssueQtyInBaseUnit += dataItems.PlannedGoodsIssueQtyInBaseUnit
	}

	return fmt.Sprint(totalPlannedGoodsIssueQtyInBaseUnit)
}
