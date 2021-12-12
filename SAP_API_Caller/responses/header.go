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
			ReceivingLocationTimeZone      string      `json:"ReceivingLocationTimeZone"`
			ActualDeliveryRoute            string      `json:"ActualDeliveryRoute"`
			ActualGoodsMovementDate        string      `json:"ActualGoodsMovementDate"`
			ActualGoodsMovementTime        string      `json:"ActualGoodsMovementTime"`
			BillingDocumentDate            string      `json:"BillingDocumentDate"`
			CompleteDeliveryIsDefined      bool        `json:"CompleteDeliveryIsDefined"`
			ConfirmationTime               string      `json:"ConfirmationTime"`
			CreationDate                   string      `json:"CreationDate"`
			CreationTime                   string      `json:"CreationTime"`
			CustomerGroup                  string      `json:"CustomerGroup"`
			DeliveryBlockReason            string      `json:"DeliveryBlockReason"`
			DeliveryDate                   string      `json:"DeliveryDate"`
			DeliveryDocument               string      `json:"DeliveryDocument"`
			DeliveryDocumentBySupplier     string      `json:"DeliveryDocumentBySupplier"`
			DeliveryDocumentType           string      `json:"DeliveryDocumentType"`
			DeliveryIsInPlant              bool        `json:"DeliveryIsInPlant"`
			DeliveryPriority               string      `json:"DeliveryPriority"`
			DeliveryTime                   string      `json:"DeliveryTime"`
			DocumentDate                   string      `json:"DocumentDate"`
			GoodsIssueOrReceiptSlipNumber  string      `json:"GoodsIssueOrReceiptSlipNumber"`
			GoodsIssueTime                 string      `json:"GoodsIssueTime"`
			HeaderBillgIncompletionStatus  string      `json:"HeaderBillgIncompletionStatus"`
			HeaderBillingBlockReason       string      `json:"HeaderBillingBlockReason"`
			HeaderDelivIncompletionStatus  string      `json:"HeaderDelivIncompletionStatus"`
			HeaderGrossWeight              string      `json:"HeaderGrossWeight"`
			HeaderNetWeight                string      `json:"HeaderNetWeight"`
			HeaderPackingIncompletionSts   string      `json:"HeaderPackingIncompletionSts"`
			HeaderPickgIncompletionStatus  string      `json:"HeaderPickgIncompletionStatus"`
			HeaderVolume                   string      `json:"HeaderVolume"`
			HeaderVolumeUnit               string      `json:"HeaderVolumeUnit"`
			HeaderWeightUnit               string      `json:"HeaderWeightUnit"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			IsExportDelivery               string      `json:"IsExportDelivery"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			LoadingDate                    string      `json:"LoadingDate"`
			LoadingPoint                   string      `json:"LoadingPoint"`
			LoadingTime                    string      `json:"LoadingTime"`
			MeansOfTransport               string      `json:"MeansOfTransport"`
			OrderCombinationIsAllowed      bool        `json:"OrderCombinationIsAllowed"`
			OrderID                        string      `json:"OrderID"`
			PickedItemsLocation            string      `json:"PickedItemsLocation"`
			PickingDate                    string      `json:"PickingDate"`
			PickingTime                    string      `json:"PickingTime"`
			PlannedGoodsIssueDate          string      `json:"PlannedGoodsIssueDate"`
			ProposedDeliveryRoute          string      `json:"ProposedDeliveryRoute"`
			ReceivingPlant                 string      `json:"ReceivingPlant"`
			RouteSchedule                  string      `json:"RouteSchedule"`
			SalesDistrict                  string      `json:"SalesDistrict"`
			SalesOffice                    string      `json:"SalesOffice"`
			SalesOrganization              string      `json:"SalesOrganization"`
			SDDocumentCategory             string      `json:"SDDocumentCategory"`
			ShipmentBlockReason            string      `json:"ShipmentBlockReason"`
			ShippingCondition              string      `json:"ShippingCondition"`
			ShippingPoint                  string      `json:"ShippingPoint"`
			ShippingType                   string      `json:"ShippingType"`
			ShipToParty                    string      `json:"ShipToParty"`
			SoldToParty                    string      `json:"SoldToParty"`
			Supplier                       string      `json:"Supplier"`
			TotalBlockStatus               string      `json:"TotalBlockStatus"`
			TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
			TotalNumberOfPackage           string      `json:"TotalNumberOfPackage"`
			TransactionCurrency            string      `json:"TransactionCurrency"`
			TransportationGroup            string      `json:"TransportationGroup"`
			TransportationPlanningDate     string      `json:"TransportationPlanningDate"`
			TransportationPlanningStatus   string      `json:"TransportationPlanningStatus"`
			TransportationPlanningTime     string      `json:"TransportationPlanningTime"`
			UnloadingPointName             string      `json:"UnloadingPointName"`
			ToPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_DeliveryDocumentPartner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_DeliveryDocumentItem"`
		} `json:"results"`
	} `json:"d"`
}
