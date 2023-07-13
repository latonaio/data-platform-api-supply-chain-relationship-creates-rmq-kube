package existence_conf

type Returns struct {
	ConnectionKey                        string                                `json:"connection_key"`
	Result                               bool                                  `json:"result"`
	RedisKey                             string                                `json:"redis_key"`
	RuntimeSessionID                     string                                `json:"runtime_session_id"`
	BusinessPartnerID                    *int                                  `json:"business_partner"`
	Filepath                             string                                `json:"filepath"`
	ServiceLabel                         string                                `json:"service_label"`
	BPGeneralReturn                      BPGeneralReturn                       `json:"BusinessPartnerGeneral"`
	SupplyChainRelationshipGeneralReturn *SupplyChainRelationshipGeneralReturn `json:"SupplyChainRelationshipGeneralReturn"`
	PaymentTermsReturn                   *PaymentTermsReturn                   `json:"PaymentTermsReturn"`
	PaymentMethodReturn                  *PaymentMethodReturn                  `json:"PaymentMethodReturn"`
	IncotermsReturn                      *IncotermsReturn                      `json:"IncotermsReturn"`
	CurrencyReturn                       *CurrencyReturn                       `json:"CurrencyReturn"`
	APISchema                            string                                `json:"api_schema"`
	Accepter                             []string                              `json:"accepter"`
	Deleted                              bool                                  `json:"deleted"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type SupplyChainRelationshipGeneralReturn struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type IncotermsReturn struct {
	Incoterms string `json:"Incoterms"`
}

type PaymentTermsReturn struct {
	PaymentTerms string `json:"PaymentTerms"`
}

type PaymentMethodReturn struct {
	PaymentMethod string `json:"PaymentMethod"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}
