package main

import (
	sap_api_caller "sap-api-integrations-inbound-delivery-reads/SAP_API_Caller"
	"sap-api-integrations-inbound-delivery-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Inbound_Delivery_Item_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}

	caller.AsyncGetInboundDelivery(
		inoutSDC.InboundDelivery.DeliveryDocument,
		inoutSDC.InboundDelivery.DeliveryDocumentItem.DeliveryDocumentItem,
		accepter,
	)
}
