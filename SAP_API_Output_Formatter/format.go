package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-inbound-delivery-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			ReceivingLocationTimeZone:     data.ReceivingLocationTimeZone,
			ActualDeliveryRoute:           data.ActualDeliveryRoute,
			ActualGoodsMovementDate:       data.ActualGoodsMovementDate,
			ActualGoodsMovementTime:       data.ActualGoodsMovementTime,
			BillingDocumentDate:           data.BillingDocumentDate,
			CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
			ConfirmationTime:              data.ConfirmationTime,
			CreationDate:                  data.CreationDate,
			CreationTime:                  data.CreationTime,
			CustomerGroup:                 data.CustomerGroup,
			DeliveryBlockReason:           data.DeliveryBlockReason,
			DeliveryDate:                  data.DeliveryDate,
			DeliveryDocument:              data.DeliveryDocument,
			DeliveryDocumentBySupplier:    data.DeliveryDocumentBySupplier,
			DeliveryDocumentType:          data.DeliveryDocumentType,
			DeliveryIsInPlant:             data.DeliveryIsInPlant,
			DeliveryPriority:              data.DeliveryPriority,
			DeliveryTime:                  data.DeliveryTime,
			DocumentDate:                  data.DocumentDate,
			GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
			GoodsIssueTime:                data.GoodsIssueTime,
			HeaderBillgIncompletionStatus: data.HeaderBillgIncompletionStatus,
			HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
			HeaderDelivIncompletionStatus: data.HeaderDelivIncompletionStatus,
			HeaderGrossWeight:             data.HeaderGrossWeight,
			HeaderNetWeight:               data.HeaderNetWeight,
			HeaderPackingIncompletionSts:  data.HeaderPackingIncompletionSts,
			HeaderPickgIncompletionStatus: data.HeaderPickgIncompletionStatus,
			HeaderVolume:                  data.HeaderVolume,
			HeaderVolumeUnit:              data.HeaderVolumeUnit,
			HeaderWeightUnit:              data.HeaderWeightUnit,
			IncotermsClassification:       data.IncotermsClassification,
			IsExportDelivery:              data.IsExportDelivery,
			LastChangeDate:                data.LastChangeDate,
			LoadingDate:                   data.LoadingDate,
			LoadingPoint:                  data.LoadingPoint,
			LoadingTime:                   data.LoadingTime,
			MeansOfTransport:              data.MeansOfTransport,
			OrderCombinationIsAllowed:     data.OrderCombinationIsAllowed,
			OrderID:                       data.OrderID,
			PickedItemsLocation:           data.PickedItemsLocation,
			PickingDate:                   data.PickingDate,
			PickingTime:                   data.PickingTime,
			PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
			ProposedDeliveryRoute:         data.ProposedDeliveryRoute,
			ReceivingPlant:                data.ReceivingPlant,
			RouteSchedule:                 data.RouteSchedule,
			SalesDistrict:                 data.SalesDistrict,
			SalesOffice:                   data.SalesOffice,
			SalesOrganization:             data.SalesOrganization,
			SDDocumentCategory:            data.SDDocumentCategory,
			ShipmentBlockReason:           data.ShipmentBlockReason,
			ShippingCondition:             data.ShippingCondition,
			ShippingPoint:                 data.ShippingPoint,
			ShippingType:                  data.ShippingType,
			ShipToParty:                   data.ShipToParty,
			SoldToParty:                   data.SoldToParty,
			Supplier:                      data.Supplier,
			TotalBlockStatus:              data.TotalBlockStatus,
			TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
			TotalNumberOfPackage:          data.TotalNumberOfPackage,
			TransactionCurrency:           data.TransactionCurrency,
			TransportationGroup:           data.TransportationGroup,
			TransportationPlanningDate:    data.TransportationPlanningDate,
			TransportationPlanningStatus:  data.TransportationPlanningStatus,
			TransportationPlanningTime:    data.TransportationPlanningTime,
			UnloadingPointName:            data.UnloadingPointName,
			ToPartner:                     data.ToPartner.Deferred.URI,
			ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			ActualDeliveredQtyInBaseUnit:   data.ActualDeliveredQtyInBaseUnit,
			ActualDeliveryQuantity:         data.ActualDeliveryQuantity,
			AdditionalCustomerGroup1:       data.AdditionalCustomerGroup1,
			AdditionalCustomerGroup2:       data.AdditionalCustomerGroup2,
			AdditionalCustomerGroup3:       data.AdditionalCustomerGroup3,
			AdditionalCustomerGroup4:       data.AdditionalCustomerGroup4,
			AdditionalCustomerGroup5:       data.AdditionalCustomerGroup5,
			BaseUnit:                       data.BaseUnit,
			Batch:                          data.Batch,
			BatchBySupplier:                data.BatchBySupplier,
			BOMExplosion:                   data.BOMExplosion,
			BusinessArea:                   data.BusinessArea,
			ControllingArea:                data.ControllingArea,
			CostCenter:                     data.CostCenter,
			CreationDate:                   data.CreationDate,
			CreationTime:                   data.CreationTime,
			DeliveryDocument:               data.DeliveryDocument,
			DeliveryDocumentItem:           data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:   data.DeliveryDocumentItemCategory,
			DeliveryDocumentItemText:       data.DeliveryDocumentItemText,
			DeliveryGroup:                  data.DeliveryGroup,
			DeliveryQuantityUnit:           data.DeliveryQuantityUnit,
			DeliveryRelatedBillingStatus:   data.DeliveryRelatedBillingStatus,
			DistributionChannel:            data.DistributionChannel,
			Division:                       data.Division,
			GLAccount:                      data.GLAccount,
			GoodsMovementReasonCode:        data.GoodsMovementReasonCode,
			GoodsMovementStatus:            data.GoodsMovementStatus,
			GoodsMovementType:              data.GoodsMovementType,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			IsNotGoodsMovementsRelevant:    data.IsNotGoodsMovementsRelevant,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			ItemBillingIncompletionStatus:  data.ItemBillingIncompletionStatus,
			ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
			ItemGdsMvtIncompletionSts:      data.ItemGdsMvtIncompletionSts,
			ItemGeneralIncompletionStatus:  data.ItemGeneralIncompletionStatus,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemIsBillingRelevant:          data.ItemIsBillingRelevant,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemPackingIncompletionStatus:  data.ItemPackingIncompletionStatus,
			ItemPickingIncompletionStatus:  data.ItemPickingIncompletionStatus,
			ItemVolume:                     data.ItemVolume,
			ItemVolumeUnit:                 data.ItemVolumeUnit,
			ItemWeightUnit:                 data.ItemWeightUnit,
			LastChangeDate:                 data.LastChangeDate,
			LoadingGroup:                   data.LoadingGroup,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialFreightGroup:           data.MaterialFreightGroup,
			MaterialGroup:                  data.MaterialGroup,
			MaterialIsBatchManaged:         data.MaterialIsBatchManaged,
			OrderID:                        data.OrderID,
			OrderItem:                      data.OrderItem,
			OriginalDeliveryQuantity:       data.OriginalDeliveryQuantity,
			PackingStatus:                  data.PackingStatus,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			PickingConfirmationStatus:      data.PickingConfirmationStatus,
			PickingStatus:                  data.PickingStatus,
			Plant:                          data.Plant,
			ProductAvailabilityDate:        data.ProductAvailabilityDate,
			ProductAvailabilityTime:        data.ProductAvailabilityTime,
			ProfitabilitySegment:           data.ProfitabilitySegment,
			ProfitCenter:                   data.ProfitCenter,
			QuantityIsFixed:                data.QuantityIsFixed,
			ReceivingPoint:                 data.ReceivingPoint,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
			SalesDocumentItemType:          data.SalesDocumentItemType,
			SalesGroup:                     data.SalesGroup,
			SalesOffice:                    data.SalesOffice,
			SDDocumentCategory:             data.SDDocumentCategory,
			SDProcessStatus:                data.SDProcessStatus,
			ShelfLifeExpirationDate:        data.ShelfLifeExpirationDate,
			StockType:                      data.StockType,
			StorageLocation:                data.StorageLocation,
			TransportationGroup:            data.TransportationGroup,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		})
	}

	return item, nil
}

func ConvertToToPartner(raw []byte, l *logger.Logger) ([]ToPartner, error) {
	pm := &responses.ToPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toPartner := make([]ToPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPartner = append(toPartner, ToPartner{
			AddressID:       data.AddressID,
			ContactPerson:   data.ContactPerson,
			Customer:        data.Customer,
			PartnerFunction: data.PartnerFunction,
			Personnel:       data.Personnel,
			SDDocument:      data.SDDocument,
			SDDocumentItem:  data.SDDocumentItem,
			Supplier:        data.Supplier,
			ToAddress:       data.ToAddress.Deferred.URI,
		})
	}

	return toPartner, nil
}

func ConvertToToAddress(raw []byte, l *logger.Logger) (*ToAddress, error) {
	pm := &responses.ToAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToAddress. unmarshal error: %w", err)
	}

	return &ToAddress{
		AddressID:              pm.D.AddressID,
		Building:               pm.D.Building,
		BusinessPartnerName1:   pm.D.BusinessPartnerName1,
		CityName:               pm.D.CityName,
		CorrespondenceLanguage: pm.D.CorrespondenceLanguage,
		Country:                pm.D.Country,
		FaxNumber:              pm.D.FaxNumber,
		Nation:                 pm.D.Nation,
		PhoneNumber:            pm.D.PhoneNumber,
		PostalCode:             pm.D.PostalCode,
		Region:                 pm.D.Region,
		StreetName:             pm.D.StreetName,
	}, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			ActualDeliveredQtyInBaseUnit:   data.ActualDeliveredQtyInBaseUnit,
			ActualDeliveryQuantity:         data.ActualDeliveryQuantity,
			AdditionalCustomerGroup1:       data.AdditionalCustomerGroup1,
			AdditionalCustomerGroup2:       data.AdditionalCustomerGroup2,
			AdditionalCustomerGroup3:       data.AdditionalCustomerGroup3,
			AdditionalCustomerGroup4:       data.AdditionalCustomerGroup4,
			AdditionalCustomerGroup5:       data.AdditionalCustomerGroup5,
			BaseUnit:                       data.BaseUnit,
			Batch:                          data.Batch,
			BatchBySupplier:                data.BatchBySupplier,
			BOMExplosion:                   data.BOMExplosion,
			BusinessArea:                   data.BusinessArea,
			ControllingArea:                data.ControllingArea,
			CostCenter:                     data.CostCenter,
			CreationDate:                   data.CreationDate,
			CreationTime:                   data.CreationTime,
			DeliveryDocument:               data.DeliveryDocument,
			DeliveryDocumentItem:           data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:   data.DeliveryDocumentItemCategory,
			DeliveryDocumentItemText:       data.DeliveryDocumentItemText,
			DeliveryGroup:                  data.DeliveryGroup,
			DeliveryQuantityUnit:           data.DeliveryQuantityUnit,
			DeliveryRelatedBillingStatus:   data.DeliveryRelatedBillingStatus,
			DistributionChannel:            data.DistributionChannel,
			Division:                       data.Division,
			GLAccount:                      data.GLAccount,
			GoodsMovementReasonCode:        data.GoodsMovementReasonCode,
			GoodsMovementStatus:            data.GoodsMovementStatus,
			GoodsMovementType:              data.GoodsMovementType,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			IsNotGoodsMovementsRelevant:    data.IsNotGoodsMovementsRelevant,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			ItemBillingIncompletionStatus:  data.ItemBillingIncompletionStatus,
			ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
			ItemGdsMvtIncompletionSts:      data.ItemGdsMvtIncompletionSts,
			ItemGeneralIncompletionStatus:  data.ItemGeneralIncompletionStatus,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemIsBillingRelevant:          data.ItemIsBillingRelevant,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemPackingIncompletionStatus:  data.ItemPackingIncompletionStatus,
			ItemPickingIncompletionStatus:  data.ItemPickingIncompletionStatus,
			ItemVolume:                     data.ItemVolume,
			ItemVolumeUnit:                 data.ItemVolumeUnit,
			ItemWeightUnit:                 data.ItemWeightUnit,
			LastChangeDate:                 data.LastChangeDate,
			LoadingGroup:                   data.LoadingGroup,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialFreightGroup:           data.MaterialFreightGroup,
			MaterialGroup:                  data.MaterialGroup,
			MaterialIsBatchManaged:         data.MaterialIsBatchManaged,
			OrderID:                        data.OrderID,
			OrderItem:                      data.OrderItem,
			OriginalDeliveryQuantity:       data.OriginalDeliveryQuantity,
			PackingStatus:                  data.PackingStatus,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			PickingConfirmationStatus:      data.PickingConfirmationStatus,
			PickingStatus:                  data.PickingStatus,
			Plant:                          data.Plant,
			ProductAvailabilityDate:        data.ProductAvailabilityDate,
			ProductAvailabilityTime:        data.ProductAvailabilityTime,
			ProfitabilitySegment:           data.ProfitabilitySegment,
			ProfitCenter:                   data.ProfitCenter,
			QuantityIsFixed:                data.QuantityIsFixed,
			ReceivingPoint:                 data.ReceivingPoint,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
			SalesDocumentItemType:          data.SalesDocumentItemType,
			SalesGroup:                     data.SalesGroup,
			SalesOffice:                    data.SalesOffice,
			SDDocumentCategory:             data.SDDocumentCategory,
			SDProcessStatus:                data.SDProcessStatus,
			ShelfLifeExpirationDate:        data.ShelfLifeExpirationDate,
			StockType:                      data.StockType,
			StorageLocation:                data.StorageLocation,
			TransportationGroup:            data.TransportationGroup,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		})
	}

	return toItem, nil
}
