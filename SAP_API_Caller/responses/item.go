package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
		} `json:"results"`
	} `json:"d"`
}
