package dpfm_api_processing_formatter

func ConvertToGeneralUpdates(generalUpdates GeneralUpdates) *GeneralUpdates {
	data := generalUpdates

	return &GeneralUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
	}
}

func ConvertToTransactionUpdates(transactionUpdates TransactionUpdates) *TransactionUpdates {
	data := transactionUpdates

	return &TransactionUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
		TransactionCurrency: 			data.TransactionCurrency,
		PaymentTerms:        			data.PaymentTerms,
		PaymentMethod:       			data.PaymentMethod,
		Incoterms:           			data.Incoterms,
		QuotationIsBlocked:  			data.QuotationIsBlocked,
		OrderIsBlocked:      			data.OrderIsBlocked,
		DeliveryIsBlocked:   			data.DeliveryIsBlocked,
		BillingIsBlocked:    			data.BillingIsBlocked,
	}
}

func ConvertToBillingRelationUpdates(billingRelationUpdates BillingRelationUpdates) *BillingRelationUpdates {
	data := billingRelationUpdates

	return &BillingRelationUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
		BillToParty:                 	data.BillToParty,
		BillFromParty:                	data.BillFromParty,
		DefaultRelation:              	data.DefaultRelation,
		TransactionTaxClassification: 	data.TransactionTaxClassification,
	}
}

func ConvertToPaymentRelationUpdates(paymentRelationUpdates PaymentRelationUpdates) *PaymentRelationUpdates {
	data := paymentRelationUpdates

	return &PaymentRelationUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
		BillToParty:                 	data.BillToParty,
		BillFromParty:                	data.BillFromParty,
		Payer:                 			data.Payer,
		Payee:                 			data.Payee,
		DefaultRelation:       			data.DefaultRelation,
		PayerHouseBank:        			data.PayerHouseBank,
		PayerHouseBankAccount: 			data.PayerHouseBankAccount,
		PayeeHouseBank:        			data.PayeeHouseBank,
		PayeeHouseBankAccount: 			data.PayeeHouseBankAccount,
	}
}

func ConvertToDeliveryRelationUpdates(deliveryRelationUpdates DeliveryRelationUpdates) *DeliveryRelationUpdates {
	data := deliveryRelationUpdates

	return &DeliveryRelationUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
		DeliverToParty:   				data.DeliverToParty,
		DeliverFromParty: 				data.DeliverFromParty,
		DefaultRelation:  				data.DefaultRelation,
	}
}

func ConvertToDeliveryPlantRelationUpdates(deliveryPlantRelationUpdates DeliveryPlantRelationUpdates) *DeliveryPlantRelationUpdates {
	data := deliveryPlantRelationUpdates

	return &DeliveryPlantRelationUpdates{
		SupplyChainRelationshipID:		data.SupplyChainRelationshipID,
		Buyer:               			data.Buyer,
		Seller:              			data.Seller,
		DeliverToParty:   				data.DeliverToParty,
		DeliverFromParty: 				data.DeliverFromParty,
		DeliverToPlant:      			data.DeliverToPlant,
		DeliverFromPlant:    			data.DeliverFromPlant,
		DefaultRelation:     			data.DefaultRelation,
		MRPType:            			data.MRPType,
		MRPController:       			data.MRPController,
	}
}

func ConvertToDeliveryPlantRelationProductUpdates(deliveryPlantRelationProductUpdates DeliveryPlantRelationProductUpdates) *DeliveryPlantRelationProductUpdates {
	data := deliveryPlantRelationProductUpdates

	return &DeliveryPlantRelationProductUpdates{
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
	}
}

func ConvertToDeliveryPlantRelationProductMRPAreaUpdates(deliveryPlantRelationProductMRPAreaUpdates DeliveryPlantRelationProductMRPAreaUpdates) *DeliveryPlantRelationProductMRPAreaUpdates {
	data := deliveryPlantRelationProductMRPAreaUpdates

	return &DeliveryPlantRelationProductMRPAreaUpdates{
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
	}
}

func ConvertToStockConfPlantRelationUpdates(stockConfPlantRelationUpdates StockConfPlantRelationUpdates) *StockConfPlantRelationUpdates {
	data := stockConfPlantRelationUpdates

	return &StockConfPlantRelationUpdates{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			DefaultRelation:                         data.DefaultRelation,
			MRPType:                                 data.MRPType,
			MRPController:                           data.MRPController,
	}
}

func ConvertToStockConfPlantRelationProductUpdates(stockConfPlantRelationProductUpdates StockConfPlantRelationProductUpdates) *StockConfPlantRelationProductUpdates {
	data := stockConfPlantRelationProductUpdates

	return &StockConfPlantRelationProductUpdates{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			Product:                                 data.Product,
			MRPType:                                 data.MRPType,
			MRPController:                           data.MRPController,
	}
}

func ConvertToProductionPlantRelationUpdates(productionPlantRelationUpdates ProductionPlantRelationUpdates) *ProductionPlantRelationUpdates {
	data := productionPlantRelationUpdates

	return &ProductionPlantRelationUpdates{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			DefaultRelation:                          data.DefaultRelation,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
	}
}

func ConvertToProductionPlantRelationProductUpdates(productionPlantRelationProductUpdates ProductionPlantRelationProductUpdates) *ProductionPlantRelationProductUpdates {
	data := productionPlantRelationProductUpdates

	return &ProductionPlantRelationProductUpdates{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			Product:                                  data.Product,
			ProductionPlantStorageLocation:           data.ProductionPlantStorageLocation,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
	}
}
