package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
