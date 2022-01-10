# sap-api-integrations-inbound-delivery-reads 
sap-api-integrations-inbound-delivery-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で入荷データ を取得するマイクロサービスです。    
sap-api-integrations-inbound-delivery-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-inbound-delivery-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_INBOUND_DELIVERY_SRV_0002/overview  

## 動作環境  
sap-api-integrations-inbound-delivery-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-inbound-delivery-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## 本レポジトリ が 対応する API サービス
sap-api-integrations-inbound-delivery-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_INBOUND_DELIVERY_SRV_0002/overview  
* APIサービス名(=baseURL): API_INBOUND_DELIVERY_SRV;v=0002

## 本レポジトリ に 含まれる API名
sap-api-integrations-inbound-delivery-reads には、次の API をコールするためのリソースが含まれています。  

* A_InbDeliveryHeader（入荷伝票 - ヘッダ）※入荷伝票の詳細データを取得するために、ToPartner、ToAddress、ToItemと合わせて利用されます。
* A_InbDeliveryItem（入荷伝票 - 明細）
* ToPartner（入荷伝票 - 取引先）
* ToAddress（入荷伝票 - アドレス）
* ToItem（入荷伝票 - 明細）

## API への 値入力条件 の 初期値
sap-api-integrations-inbound-delivery-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.InboundDelivery.DeliveryDocument（入荷伝票）
* inoutSDC.InboundDelivery.DeliveryDocumentItem.DeliveryDocumentItem（入荷伝票明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.inbounddelivery.v1.InboundDelivery.Created.v1",
	"accepter": ["Header"],
	"delivery_document": "180000000",
	"deleted": ""
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.inbounddelivery.v1.InboundDelivery.Created.v1",
	"accepter": ["Item"],
	"delivery_document": "180000000",
	"deleted": ""
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetInboundDelivery(deliveryDocument, deliveryDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(deliveryDocument)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(deliveryDocument, deliveryDocumentItem)
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
以下の sample.json の例は、SAP 入荷伝票　の　ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ReceivingLocationTimeZone" ～ "ToPartner" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-inbound-delivery-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-inbound-delivery-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"ReceivingLocationTimeZone": "UTC",
			"ActualDeliveryRoute": "",
			"ActualGoodsMovementDate": "/Date(1484092800000)/",
			"ActualGoodsMovementTime": "PT00H00M00S",
			"BillingDocumentDate": "",
			"CompleteDeliveryIsDefined": false,
			"ConfirmationTime": "PT00H00M00S",
			"CreationDate": "/Date(1484092800000)/",
			"CreationTime": "PT11H32M52S",
			"CustomerGroup": "",
			"DeliveryBlockReason": "",
			"DeliveryDate": "/Date(1485734400000)/",
			"DeliveryDocument": "180000000",
			"DeliveryDocumentBySupplier": "ASN#451435",
			"DeliveryDocumentType": "EL",
			"DeliveryIsInPlant": false,
			"DeliveryPriority": "00",
			"DeliveryTime": "PT22H30M00S",
			"DocumentDate": "/Date(1484092800000)/",
			"GoodsIssueOrReceiptSlipNumber": "",
			"GoodsIssueTime": "PT00H00M00S",
			"HeaderBillgIncompletionStatus": "C",
			"HeaderBillingBlockReason": "",
			"HeaderDelivIncompletionStatus": "C",
			"HeaderGrossWeight": "10.000",
			"HeaderNetWeight": "9.000",
			"HeaderPackingIncompletionSts": "C",
			"HeaderPickgIncompletionStatus": "C",
			"HeaderVolume": "0.000",
			"HeaderVolumeUnit": "",
			"HeaderWeightUnit": "KG",
			"IncotermsClassification": "",
			"IsExportDelivery": "",
			"LastChangeDate": "/Date(1484092800000)/",
			"LoadingDate": "",
			"LoadingPoint": "",
			"LoadingTime": "PT00H00M00S",
			"MeansOfTransport": "",
			"OrderCombinationIsAllowed": true,
			"OrderID": "",
			"PickedItemsLocation": "",
			"PickingDate": "",
			"PickingTime": "PT00H00M00S",
			"PlannedGoodsIssueDate": "",
			"ProposedDeliveryRoute": "",
			"ReceivingPlant": "",
			"RouteSchedule": "",
			"SalesDistrict": "",
			"SalesOffice": "",
			"SalesOrganization": "",
			"SDDocumentCategory": "7",
			"ShipmentBlockReason": "",
			"ShippingCondition": "01",
			"ShippingPoint": "",
			"ShippingType": "",
			"ShipToParty": "",
			"SoldToParty": "",
			"Supplier": "17300080",
			"TotalBlockStatus": "",
			"TotalCreditCheckStatus": "",
			"TotalNumberOfPackage": "00000",
			"TransactionCurrency": "",
			"TransportationGroup": "0001",
			"TransportationPlanningDate": "",
			"TransportationPlanningStatus": "",
			"TransportationPlanningTime": "PT00H00M00S",
			"UnloadingPointName": "",
			"to_Partner": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_INBOUND_DELIVERY_SRV;v=0002/A_InbDeliveryHeader('180000000')/to_DeliveryDocumentPartner",
			"to_Item": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_INBOUND_DELIVERY_SRV;v=0002/A_InbDeliveryHeader('180000000')/to_DeliveryDocumentItem"
		}
	],
	"time": "2022-01-10T10:43:59.499612+09:00"
}
```
