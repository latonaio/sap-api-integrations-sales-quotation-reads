# sap-api-integrations-sales-quotation-reads 
sap-api-integrations-sales_quotation-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で販売見積データ を取得するマイクロサービスです。    
sap-api-integrations-sales_quotation-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-sales_quotation-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SALES_QUOTATION_SRV_0001/overview

## 動作環境  
sap-api-integrations-sales-quotation-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-sales-quotation-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-quotation-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_QUOTATION_SRV_0001/overview
* APIサービス名(=baseURL): API_SALES_QUOTATION_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-quotation-reads には、次の API をコールするためのリソースが含まれています。  

* A_SalesQuotation（販売見積 - ヘッダ）※販売見積の詳細データを取得するために、ToHeaderPartner、ToItem、と合わせて利用されます。
* ToItem（販売見積 - 明細）※販売見積明細の詳細データを取得するために、ToItemPricingElementと合わせて利用されます。
* ToHeaderPartner（販売見積 - ヘッダ取引先）
* ToItemPricingElement（販売見積 - 明細価格条件）
* A_SalesQuotationItem（販売見積明細）※販売見積明細の詳細データを取得するために、ToItemPricingElementと合わせて利用されます。
* ToItemPricingElement（販売見積 - 明細価格条件）

## API への 値入力条件 の 初期値
sap-api-integrations-sales-quotation-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.SalesQuotation.SalesQuotation（販売見積番号）
* inputSDC.SalesQuotation.SalesQuotationItem.SalesQuotationItem（販売見積明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPSalesQuotationReads",
	"accepter": ["Header"],
	"sales_quotaion": "20000001",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPSalesQuotationReads",
	"accepter": ["All"],
	"sales_quotaion": "20000001",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetSalesQuotation(salesQuotation, salesQuotationItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(salesQuotation)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(salesQuotation, salesQuotationItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 販売見積 の ヘッダデータ が取得された結果の JSON の例です。
以下の項目のうち、"SalesInquiry" ～ "ToHeaderPartner" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-quotation-reads/SAP_API_Caller/caller.go#L60",
	"function": "sap-api-integrations-sales-quotation-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"SalesQuotation": "20000001",
			"SalesQuotationType": "QT",
			"SalesOrganization": "0001",
			"DistributionChannel": "01",
			"OrganizationDivision": "01",
			"SalesGroup": "",
			"SalesOffice": "",
			"SalesDistrict": "000001",
			"SoldToParty": "1",
			"CreationDate": "2022-09-20",
			"LastChangeDate": "",
			"PurchaseOrderByCustomer": "",
			"CustomerPurchaseOrderDate": "",
			"SalesQuotationDate": "2022-09-20",
			"TotalNetAmount": "20000.00",
			"TransactionCurrency": "EUR",
			"SDDocumentReason": "",
			"PricingDate": "2022-09-20",
			"RequestedDeliveryDate": "2022-09-30",
			"ShippingCondition": "01",
			"CompleteDeliveryIsDefined": false,
			"ShippingType": "",
			"HeaderBillingBlockReason": "",
			"DeliveryBlockReason": "",
			"BindingPeriodValidityStartDate": "2022-09-20",
			"BindingPeriodValidityEndDate": "2022-09-22",
			"HdrOrderProbabilityInPercent": "70",
			"ExpectedOrderNetAmount": "14000.00",
			"IncotermsClassification": "FH",
			"CustomerPaymentTerms": "0001",
			"CustomerPriceGroup": "01",
			"PriceListType": "",
			"PaymentMethod": "",
			"CustomerTaxClassification1": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentCategory": "",
			"SalesQuotationApprovalReason": "",
			"SalesDocApprovalStatus": "",
			"OverallSDProcessStatus": "A",
			"TotalCreditCheckStatus": "",
			"OverallSDDocumentRejectionSts": "A",
			"to_Partner": "http://100.21.57.120:8080/sap/opu/odata/sap/API_SALES_QUOTATION_SRV/A_SalesQuotation('20000001')/to_Partner",
			"to_Item": "http://100.21.57.120:8080/sap/opu/odata/sap/API_SALES_QUOTATION_SRV/A_SalesQuotation('20000001')/to_Item"
		}
	],
	"time": "2022-09-20T16:39:12+09:00"
}
```
