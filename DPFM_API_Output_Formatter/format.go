package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneral(SDC *dpfm_api_input_reader.SDC) *General {
	data := SDC.General

	generalCreates := &General{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		CreationDate:              data.CreationDate,
		LastChangeDate:            data.LastChangeDate,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}

	return generalCreates

}

func ConvertToTransaction(SDC *dpfm_api_input_reader.SDC) *[]Transaction {
	var transaction []Transaction

	for _, data := range SDC.General.Transaction {
		transaction = append(transaction, Transaction{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			TransactionCurrency:       data.TransactionCurrency,
			PaymentTerms:              data.PaymentTerms,
			PaymentMethod:             data.PaymentMethod,
			Incoterms:                 data.Incoterms,
			AccountAssignmentGroup:    data.AccountAssignmentGroup,
			CreationDate:              data.CreationDate,
			LastChangeDate:            data.LastChangeDate,
			QuotationIsBlocked:        data.QuotationIsBlocked,
			OrderIsBlocked:            data.OrderIsBlocked,
			DeliveryIsBlocked:         data.DeliveryIsBlocked,
			BillingIsBlocked:          data.BillingIsBlocked,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}

	return &transaction
}

func ConvertToBillingRelation(SDC *dpfm_api_input_reader.SDC) *[]BillingRelation {
	var billingRelation []BillingRelation

	for _, data := range SDC.General.BillingRelation {
		billingRelation = append(billingRelation, BillingRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			DefaultRelation:                  data.DefaultRelation,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			IsExportImport:                   data.IsExportImport,
			TransactionTaxCategory:           data.TransactionTaxCategory,
			TransactionTaxClassification:     data.TransactionTaxClassification,
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}

	return &billingRelation
}

func ConvertToPaymentRelation(SDC *dpfm_api_input_reader.SDC) *[]PaymentRelation {
	var paymentRelation []PaymentRelation

	for _, billingRelation := range SDC.General.BillingRelation {
		for _, data := range billingRelation.PaymentRelation {
			paymentRelation = append(paymentRelation, PaymentRelation{
				SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
				SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
				SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
				Buyer:                            data.Buyer,
				Seller:                           data.Seller,
				BillToParty:                      data.BillToParty,
				BillFromParty:                    data.BillFromParty,
				Payer:                            data.Payer,
				Payee:                            data.Payee,
				DefaultRelation:                  data.DefaultRelation,
				PayerHouseBank:                   data.PayerHouseBank,
				PayerHouseBankAccount:            data.PayerHouseBankAccount,
				PayeeHouseBank:                   data.PayeeHouseBank,
				PayeeHouseBankAccount:            data.PayeeHouseBankAccount,
				CreationDate:                     data.CreationDate,
				LastChangeDate:                   data.LastChangeDate,
				IsMarkedForDeletion:              data.IsMarkedForDeletion,
			})
		}
	}

	return &paymentRelation
}

func ConvertToDeliveryRelation(SDC *dpfm_api_input_reader.SDC) *[]DeliveryRelation {
	var deliveryRelation []DeliveryRelation

	for _, data := range SDC.General.DeliveryRelation {
		deliveryRelation = append(deliveryRelation, DeliveryRelation{
			SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
			Buyer:                             data.Buyer,
			Seller:                            data.Seller,
			DeliverToParty:                    data.DeliverToParty,
			DeliverFromParty:                  data.DeliverFromParty,
			DefaultRelation:                   data.DefaultRelation,
			CreationDate:                      data.CreationDate,
			LastChangeDate:                    data.LastChangeDate,
			IsMarkedForDeletion:               data.IsMarkedForDeletion,
		})
	}

	return &deliveryRelation
}

func ConvertToDeliveryPlantRelation(SDC *dpfm_api_input_reader.SDC) *[]DeliveryPlantRelation {
	var deliveryPlantRelation []DeliveryPlantRelation

	for _, deliveryRelation := range SDC.General.DeliveryRelation {
		for _, data := range deliveryRelation.DeliveryPlantRelation {
			deliveryPlantRelation = append(deliveryPlantRelation, DeliveryPlantRelation{
				SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
				SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
				SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
				Buyer:                                  data.Buyer,
				Seller:                                 data.Seller,
				DeliverToParty:                         data.DeliverToParty,
				DeliverFromParty:                       data.DeliverFromParty,
				DeliverToPlant:                         data.DeliverToPlant,
				DeliverFromPlant:                       data.DeliverFromPlant,
				DefaultRelation:                        data.DefaultRelation,
				MRPType:                                data.MRPType,
				MRPController:                          data.MRPController,
				CreationDate:                           data.CreationDate,
				LastChangeDate:                         data.LastChangeDate,
				IsMarkedForDeletion:                    data.IsMarkedForDeletion,
			})
		}
	}

	return &deliveryPlantRelation
}

func ConvertToDeliveryPlantRelationProduct(SDC *dpfm_api_input_reader.SDC) *[]DeliveryPlantRelationProduct {
	var deliveryPlantRelationProduct []DeliveryPlantRelationProduct

	for _, deliveryRelation := range SDC.General.DeliveryRelation {
		for _, deliveryPlantRelation := range deliveryRelation.DeliveryPlantRelation {
			for _, data := range deliveryPlantRelation.DeliveryPlantRelationProduct {
				deliveryPlantRelationProduct = append(deliveryPlantRelationProduct, DeliveryPlantRelationProduct{
			SupplyChainRelationshipID:                 data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:         data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:    data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                     data.Buyer,
			Seller:                                    data.Seller,
			DeliverToParty:                            data.DeliverToParty,
			DeliverFromParty:                          data.DeliverFromParty,
			DeliverToPlant:                            data.DeliverToPlant,
			DeliverFromPlant:                          data.DeliverFromPlant,
			Product:                                   data.Product,
			DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
			DeliveryUnit:                              data.DeliveryUnit,
			QuantityPerPackage:                        data.QuantityPerPackage,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
			MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
				})
			}
		}
	}

	return &deliveryPlantRelationProduct
}

func ConvertToDeliveryPlantRelationProductMRPArea(SDC *dpfm_api_input_reader.SDC) *[]DeliveryPlantRelationProductMRPArea {
	var deliveryPlantRelationProductMRPArea []DeliveryPlantRelationProductMRPArea

	for _, deliveryRelation := range SDC.General.DeliveryRelation {
		for _, deliveryPlantRelation := range deliveryRelation.DeliveryPlantRelation {
			for _, deliveryPlantRelationProduct := range deliveryPlantRelation.DeliveryPlantRelationProduct {
				for _, data := range deliveryPlantRelationProduct.DeliveryPlantRelationProductMRPArea {
					deliveryPlantRelationProductMRPArea = append(deliveryPlantRelationProductMRPArea, DeliveryPlantRelationProductMRPArea{
			SupplyChainRelationshipID:                 data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:         data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:    data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                     data.Buyer,
			Seller:                                    data.Seller,
			DeliverToParty:                            data.DeliverToParty,
			DeliverFromParty:                          data.DeliverFromParty,
			DeliverToPlant:                            data.DeliverToPlant,
			DeliverFromPlant:                          data.DeliverFromPlant,
			Product:                                   data.Product,
			MRPArea:                                   data.MRPArea,
			DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
			DeliveryUnit:                              data.DeliveryUnit,
			QuantityPerPackage:                        data.QuantityPerPackage,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
			MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
					})
				}
			}
		}
	}

	return &deliveryPlantRelationProductMRPArea
}

func ConvertToStockConfPlantRelation(SDC *dpfm_api_input_reader.SDC) *[]StockConfPlantRelation {
	var stockConfPlantRelation []StockConfPlantRelation

	for _, data := range SDC.General.StockConfPlantRelation {
		stockConfPlantRelation = append(stockConfPlantRelation, StockConfPlantRelation{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			DefaultRelation:                         data.DefaultRelation,
			MRPType:                                 data.MRPType,
			MRPController:                           data.MRPController,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}

	return &stockConfPlantRelation
}

func ConvertToStockConfPlantRelationProduct(SDC *dpfm_api_input_reader.SDC) *[]StockConfPlantRelationProduct {
	var stockConfPlantRelationProduct []StockConfPlantRelationProduct

	for _, stockConfPlantRelation := range SDC.General.StockConfPlantRelation {
		for _, data := range stockConfPlantRelation.StockConfPlantRelationProduct {
			stockConfPlantRelationProduct = append(stockConfPlantRelationProduct, StockConfPlantRelationProduct{
				SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
				SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
				Buyer:                                   data.Buyer,
				Seller:                                  data.Seller,
				StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
				StockConfirmationPlant:                  data.StockConfirmationPlant,
				Product:                                 data.Product,
				MRPType:                                 data.MRPType,
				MRPController:                           data.MRPController,
				CreationDate:                            data.CreationDate,
				LastChangeDate:                          data.LastChangeDate,
				IsMarkedForDeletion:                     data.IsMarkedForDeletion,
			})
		}
	}

	return &stockConfPlantRelationProduct
}

func ConvertToProductionPlantRelation(SDC *dpfm_api_input_reader.SDC) *[]ProductionPlantRelation {
	var productionPlantRelation []ProductionPlantRelation

	for _, data := range SDC.General.ProductionPlantRelation {
		productionPlantRelation = append(productionPlantRelation, ProductionPlantRelation{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			DefaultRelation:                          data.DefaultRelation,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}

	return &productionPlantRelation
}

func ConvertToProductionPlantRelationProductMRPArea(SDC *dpfm_api_input_reader.SDC) *[]ProductionPlantRelationProductMRPArea {
	var productionPlantRelationProductMRPArea []ProductionPlantRelationProductMRPArea

	for _, productionPlantRelation := range SDC.General.ProductionPlantRelation {
		for _, data := range productionPlantRelation.ProductionPlantRelationProductMRPArea {
			productionPlantRelationProductMRPArea = append(productionPlantRelationProductMRPArea, ProductionPlantRelationProductMRPArea{
				SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
				SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
				Buyer:                                    data.Buyer,
				Seller:                                   data.Seller,
				ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
				ProductionPlant:                          data.ProductionPlant,
				Product:                                  data.Product,
				MRPArea:                                  data.MRPArea,
				ProductionPlantStorageLocation:           data.ProductionPlantStorageLocation,
				MRPType:                                  data.MRPType,
				MRPController:                            data.MRPController,
				CreationDate:                             data.CreationDate,
				LastChangeDate:                           data.LastChangeDate,
				IsMarkedForDeletion:                      data.IsMarkedForDeletion,
			})
		}
	}

	return &productionPlantRelationProductMRPArea
}
