package sap_api_output_formatter

type SalesInquiry struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	SalesInquiry  string `json:"sales_inquiry"`
	Deleted       bool   `json:"deleted"`
}    
    
type Header struct {
	SalesQuotation                 string      `json:"SalesQuotation"`
	SalesQuotationType             string      `json:"SalesQuotationType"`
	SalesOrganization              string      `json:"SalesOrganization"`
	DistributionChannel            string      `json:"DistributionChannel"`
	OrganizationDivision           string      `json:"OrganizationDivision"`
	SalesGroup                     string      `json:"SalesGroup"`
	SalesOffice                    string      `json:"SalesOffice"`
	SalesDistrict                  string      `json:"SalesDistrict"`
	SoldToParty                    string      `json:"SoldToParty"`
	CreationDate                   string      `json:"CreationDate"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
	SalesQuotationDate             string      `json:"SalesQuotationDate"`
	TotalNetAmount                 string      `json:"TotalNetAmount"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	SDDocumentReason               string      `json:"SDDocumentReason"`
	PricingDate                    string      `json:"PricingDate"`
	RequestedDeliveryDate          string      `json:"RequestedDeliveryDate"`
	ShippingCondition              string      `json:"ShippingCondition"`
	CompleteDeliveryIsDefined      bool        `json:"CompleteDeliveryIsDefined"`
	ShippingType                   string      `json:"ShippingType"`
	HeaderBillingBlockReason       string      `json:"HeaderBillingBlockReason"`
	DeliveryBlockReason            string      `json:"DeliveryBlockReason"`
	BindingPeriodValidityStartDate string      `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   string      `json:"BindingPeriodValidityEndDate"`
	HdrOrderProbabilityInPercent   string      `json:"HdrOrderProbabilityInPercent"`
	ExpectedOrderNetAmount         string      `json:"ExpectedOrderNetAmount"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	CustomerPriceGroup             string      `json:"CustomerPriceGroup"`
	PriceListType                  string      `json:"PriceListType"`
	PaymentMethod                  string      `json:"PaymentMethod"`
	CustomerTaxClassification1     string      `json:"CustomerTaxClassification1"`
	ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentCategory    string      `json:"ReferenceSDDocumentCategory"`
	SalesQuotationApprovalReason   string      `json:"SalesQuotationApprovalReason"`
	SalesDocApprovalStatus         string      `json:"SalesDocApprovalStatus"`
	OverallSDProcessStatus         string      `json:"OverallSDProcessStatus"`
	TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
	OverallSDDocumentRejectionSts  string      `json:"OverallSDDocumentRejectionSts"`
	ToHeaderPartner                string      `json:"to_Partner"`
	ToItem                         string      `json:"to_Item"`	
}

type Item struct {
	SalesQuotation                string `json:"SalesQuotation"`
	SalesQuotationItem            string `json:"SalesQuotationItem"`
	SalesQuotationItemCategory    string `json:"SalesQuotationItemCategory"`
	SalesQuotationItemText        string `json:"SalesQuotationItemText"`
	PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
	Material                      string `json:"Material"`
	MaterialByCustomer            string `json:"MaterialByCustomer"`
	RequestedQuantity             string `json:"RequestedQuantity"`
	RequestedQuantityUnit         string `json:"RequestedQuantityUnit"`
	ItemOrderProbabilityInPercent string `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               string `json:"ItemGrossWeight"`
	ItemNetWeight                 string `json:"ItemNetWeight"`
	ItemWeightUnit                string `json:"ItemWeightUnit"`
	ItemVolume                    string `json:"ItemVolume"`
	ItemVolumeUnit                string `json:"ItemVolumeUnit"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	NetAmount                     string `json:"NetAmount"`
	MaterialGroup                 string `json:"MaterialGroup"`
	MaterialPricingGroup          string `json:"MaterialPricingGroup"`
	Batch                         string `json:"Batch"`
	Plant                         string `json:"Plant"`
	IncotermsClassification       string `json:"IncotermsClassification"`
	CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
	ProductTaxClassification1     string `json:"ProductTaxClassification1"`
	SalesDocumentRjcnReason       string `json:"SalesDocumentRjcnReason"`
	WBSElement                    string `json:"WBSElement"`
	ProfitCenter                  string `json:"ProfitCenter"`
	ReferenceSDDocument           string `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem       string `json:"ReferenceSDDocumentItem"`
	SDProcessStatus               string `json:"SDProcessStatus"`
	ToItemPricingElement          string `json:"to_PricingElement"`
}

type ToHeaderPartner struct {
	SalesQuotation   string `json:"SalesQuotation"`
	PartnerFunction  string `json:"PartnerFunction"`
	Customer         string `json:"Customer"`
	Supplier         string `json:"Supplier"`
}

type ToItem struct {
	SalesQuotation                string `json:"SalesQuotation"`
	SalesQuotationItem            string `json:"SalesQuotationItem"`
	SalesQuotationItemCategory    string `json:"SalesQuotationItemCategory"`
	SalesQuotationItemText        string `json:"SalesQuotationItemText"`
	PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
	Material                      string `json:"Material"`
	MaterialByCustomer            string `json:"MaterialByCustomer"`
	RequestedQuantity             string `json:"RequestedQuantity"`
	RequestedQuantityUnit         string `json:"RequestedQuantityUnit"`
	ItemOrderProbabilityInPercent string `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               string `json:"ItemGrossWeight"`
	ItemNetWeight                 string `json:"ItemNetWeight"`
	ItemWeightUnit                string `json:"ItemWeightUnit"`
	ItemVolume                    string `json:"ItemVolume"`
	ItemVolumeUnit                string `json:"ItemVolumeUnit"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	NetAmount                     string `json:"NetAmount"`
	MaterialGroup                 string `json:"MaterialGroup"`
	MaterialPricingGroup          string `json:"MaterialPricingGroup"`
	Batch                         string `json:"Batch"`
	Plant                         string `json:"Plant"`
	IncotermsClassification       string `json:"IncotermsClassification"`
	CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
	ProductTaxClassification1     string `json:"ProductTaxClassification1"`
	SalesDocumentRjcnReason       string `json:"SalesDocumentRjcnReason"`
	WBSElement                    string `json:"WBSElement"`
	ProfitCenter                  string `json:"ProfitCenter"`
	ReferenceSDDocument           string `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem       string `json:"ReferenceSDDocumentItem"`
	SDProcessStatus               string `json:"SDProcessStatus"`
	ToItemPricingElement          string `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	SalesQuotation                 string `json:"SalesQuotation"`
	SalesQuotationItem             string `json:"SalesQuotationItem"`
	PricingProcedureStep           string `json:"PricingProcedureStep"`
	PricingProcedureCounter        string `json:"PricingProcedureCounter"`
	ConditionType                  string `json:"ConditionType"`
	PricingDateTime                string `json:"PricingDateTime"`
	PriceConditionDeterminationDte string `json:"PriceConditionDeterminationDte"`
	ConditionCalculationType       string `json:"ConditionCalculationType"`
	ConditionBaseValue             string `json:"ConditionBaseValue"`
	ConditionRateValue             string `json:"ConditionRateValue"`
	ConditionCurrency              string `json:"ConditionCurrency"`
	ConditionQuantity              string `json:"ConditionQuantity"`
	ConditionQuantityUnit          string `json:"ConditionQuantityUnit"`
	ConditionCategory              string `json:"ConditionCategory"`
	PricingScaleType               string `json:"PricingScaleType"`
	ConditionRecord                string `json:"ConditionRecord"`
	ConditionSequentialNumber      string `json:"ConditionSequentialNumber"`
	TaxCode                        string `json:"TaxCode"`
	ConditionAmount                string `json:"ConditionAmount"`
	TransactionCurrency            string `json:"TransactionCurrency"`
	PricingScaleBasis              string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue       string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit        string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency    string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged     bool   `json:"ConditionIsManuallyChanged"`
}
