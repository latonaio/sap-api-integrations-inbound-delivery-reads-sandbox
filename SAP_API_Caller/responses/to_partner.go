package responses

type ToPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			AddressID       string `json:"AddressID"`
			ContactPerson   string `json:"ContactPerson"`
			Customer        string `json:"Customer"`
			PartnerFunction string `json:"PartnerFunction"`
			Personnel       string `json:"Personnel"`
			SDDocument      string `json:"SDDocument"`
			SDDocumentItem  string `json:"SDDocumentItem"`
			Supplier        string `json:"Supplier"`
			ToAddress       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Address"`
		} `json:"results"`
	} `json:"d"`
}
