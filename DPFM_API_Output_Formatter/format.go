package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-instruction-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
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
				DeliveryDocument:						dataItems.DeliveryDocument,
				DeliveryDocumentItem:					dataItems.DeliveryDocumentItem,
				Product:								dataItems.Product,
				ProductSpecification:					dataItems.ProductSpecification,
				SizeOrDimensionText:					dataItems.SizeOrDimensionText,
				DeliveryDocumentItemText:        		dataItems.DeliveryDocumentItemText,
				DeliveryDocumentItemTextByBuyer:        dataItems.DeliveryDocumentItemTextByBuyer,
				DeliveryDocumentItemTextBySeller:       dataItems.DeliveryDocumentItemTextBySeller,
				PlannedGoodsIssueQuantity:				dataItems.PlannedGoodsIssueQuantity,
				PlannedGoodsIssueQtyInBaseUnit:			dataItems.PlannedGoodsIssueQtyInBaseUnit,
				BaseUnit:								dataItems.BaseUnit,
				DeliveryUnit:        					dataItems.DeliveryUnit,
				PlannedGoodsIssueDate:        			dataItems.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:        			dataItems.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:        		dataItems.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:        		dataItems.PlannedGoodsReceiptTime,
				ItemWeightUnit:							dataItems.ItemWeightUnit,
				ItemNetWeight:							dataItems.ItemNetWeight,
				ItemGrossWeight:						dataItems.ItemGrossWeight,
				Project: 								dataItems.Project,
				WBSElement: 							dataItems.WBSElement,
			})
		}
	}

	pm = &Header{
				DeliveryDocument:   					headerData.DeliveryDocument,
				DeliveryDocumentDate:   				headerData.DeliveryDocumentDate,
				DeliverToParty: 						headerData.DeliverToParty,
				DeliverToPartyName: 					headerData.DeliverToPartyName,
				DeliverFromParty: 						headerData.DeliverFromParty,
				DeliverFromPartyName:              		headerData.DeliverFromPartyName,
				IsExportImport:							headerData.IsExportImport,
				OrderID:								headerData.OrderID,
				OrderItem:								headerData.OrderItem,
				PlannedGoodsIssueDate: 					headerData.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime: 					headerData.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate: 				headerData.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: 				headerData.PlannedGoodsReceiptTime,
				HeaderGrossWeight: 						headerData.HeaderGrossWeight,
				HeaderNetWeight: 						headerData.HeaderNetWeight,
				HeaderWeightUnit: 						headerData.HeaderWeightUnit,
				Incoterms: 								headerData.Incoterms,
				Items:     								items,
	}

	return pm
}
