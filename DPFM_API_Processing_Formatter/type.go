package dpfm_api_processing_formatter

type Message struct {
	GeneralUpdates               					*GeneralUpdates                 				`json:"GeneralUpdates"`
	TransactionUpdates           					*[]TransactionUpdates           				`json:"TransactionUpdates"`
	BillingRelationUpdates       					*[]BillingRelationUpdates       				`json:"BillingRelationUpdates"`
	PaymentRelationUpdates       					*[]PaymentRelationUpdates       				`json:"PaymentRelationUpdates"`
	DeliveryRelationUpdates      					*[]DeliveryRelationUpdates      				`json:"DeliveryRelationUpdates"`
	DeliveryPlantRelationUpdates 					*[]DeliveryPlantRelationUpdates 				`json:"DeliveryPlantRelationUpdates"`
	DeliveryPlantRelationProductUpdates 			*[]DeliveryPlantRelationProductUpdates			`json:"DeliveryPlantRelationProductUpdates"`
	DeliveryPlantRelationProductMRPAreaUpdates		*[]DeliveryPlantRelationProductMRPAreaUpdates	`json:"DeliveryPlantRelationProductMRPAreaUpdates"`
	ProductionPlantRelationUpdates		 			*[]ProductionPlantRelationUpdates				`json:"ProductionPlantRelationUpdates"`
	ProductionPlantRelationProductMRPAreaUpdates	*[]ProductionPlantRelationProductMRPAreaUpdates	`json:"ProductionPlantRelationProductMRPAreaUpdates"`
	StockConfPlantRelationUpdates 					*[]StockConfPlantRelationUpdates 				`json:"StockConfPlantRelationUpdates"`
	StockConfPlantRelationProductUpdates 			*[]StockConfPlantRelationProductUpdates			`json:"StockConfPlantRelationProductUpdates"`
}

type GeneralUpdates struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
}

type TransactionUpdates struct {
	SupplyChainRelationshipID 	int     `json:"SupplyChainRelationshipID"`
	Buyer                     	int     `json:"Buyer"`
	Seller                    	int     `json:"Seller"`
	TransactionCurrency			*string `json:"TransactionCurrency"`
	PaymentTerms        		*string `json:"PaymentTerms"`
	PaymentMethod       		*string `json:"PaymentMethod"`
	Incoterms           		*string `json:"Incoterms"`
	QuotationIsBlocked  		*bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked      		*bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked   		*bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked    		*bool   `json:"BillingIsBlocked"`
}

type BillingRelationUpdates struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	DefaultRelation              	 *bool   `json:"DefaultRelation"`
	TransactionTaxClassification 	 *string `json:"TransactionTaxClassification"`
}

type PaymentRelationUpdates struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	Payer                            int     `json:"Payer"`
	Payee                            int     `json:"Payee"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	PayerHouseBank                   *string `json:"PayerHouseBank"`
	PayerHouseBankAccount            *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank                   *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount            *string `json:"PayeeHouseBankAccount"`
}

type DeliveryRelationUpdates struct {
	SupplyChainRelationshipID         int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int     `json:"Buyer"`
	Seller                            int     `json:"Seller"`
	DeliverToParty                    int     `json:"DeliverToParty"`
	DeliverFromParty                  int     `json:"DeliverFromParty"`
	DefaultRelation                   *bool   `json:"DefaultRelation"`
}

type DeliveryPlantRelationUpdates struct {
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	DefaultRelation                        *bool   `json:"DefaultRelation"`
	MRPType                                *string `json:"MRPType"`
	MRPController                          *string `json:"MRPController"`
}

type DeliveryPlantRelationProductUpdates struct {
	SupplyChainRelationshipID                 int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int      `json:"Buyer"`
	Seller                                    int      `json:"Seller"`
	DeliverToParty                            int      `json:"DeliverToParty"`
	DeliverFromParty                          int      `json:"DeliverFromParty"`
	DeliverToPlant                            string   `json:"DeliverToPlant"`
	DeliverFromPlant                          string   `json:"DeliverFromPlant"`
	Product                                   string   `json:"Product"`
	DeliverToPlantStorageLocation             string   `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string   `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string   `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32 `json:"QuantityPerPackage"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
    MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
    MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
    StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
    StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
    MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
    MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
    DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration            	  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit			  *string  `json:"StandardDeliveryDurationUnit"`
}

type DeliveryPlantRelationProductMRPAreaUpdates struct {
	SupplyChainRelationshipID                 int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int      `json:"Buyer"`
	Seller                                    int      `json:"Seller"`
	DeliverToParty                            int      `json:"DeliverToParty"`
	DeliverFromParty                          int      `json:"DeliverFromParty"`
	DeliverToPlant                            string   `json:"DeliverToPlant"`
	DeliverFromPlant                          string   `json:"DeliverFromPlant"`
	Product                                   string   `json:"Product"`
	MRPArea                                   string   `json:"MRPArea"`
	DeliverToPlantStorageLocation             string   `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string   `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string   `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32 `json:"QuantityPerPackage"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration				  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit			  *string  `json:"StandardDeliveryDurationUnit"`
}

type StockConfPlantRelationUpdates struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	StockConfirmationBusinessPartner        int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string  `json:"StockConfirmationPlant"`
	DefaultRelation                         *bool   `json:"DefaultRelation"`
	MRPType                                 *string `json:"MRPType"`
	MRPController                           *string `json:"MRPController"`
}

type StockConfPlantRelationProductUpdates struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	StockConfirmationBusinessPartner        int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string  `json:"StockConfirmationPlant"`
	Product                                 string  `json:"Product"`
	MRPType                                 *string `json:"MRPType"`
	MRPController                           *string `json:"MRPController"`
}

type ProductionPlantRelationUpdates struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	DefaultRelation                          *bool   `json:"DefaultRelation"`
	MRPType                                  *string `json:"MRPType"`
	MRPController                            *string `json:"MRPController"`
}

type ProductionPlantRelationProductMRPAreaUpdates struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	Product                                  string  `json:"Product"`
	MRPArea                                  string  `json:"MRPArea"`
	ProductionPlantStorageLocation           *string `json:"ProductionPlantStorageLocation"`
	MRPType                                  *string `json:"MRPType"`
	MRPController                            *string `json:"MRPController"`
}
