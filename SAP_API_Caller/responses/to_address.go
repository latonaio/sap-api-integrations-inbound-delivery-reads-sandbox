package responses

type ToAddress struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		AddressID                  string `json:"AddressID"`
		Building                   string `json:"Building"`
		BusinessPartnerName1       string `json:"BusinessPartnerName1"`
		CityName                   string `json:"CityName"`
		CorrespondenceLanguage     string `json:"CorrespondenceLanguage"`
		Country                    string `json:"Country"`
		FaxNumber                  string `json:"FaxNumber"`
		Nation                     string `json:"Nation"`
		PhoneNumber                string `json:"PhoneNumber"`
		PostalCode                 string `json:"PostalCode"`
		Region                     string `json:"Region"`
		StreetName                 string `json:"StreetName"`
	} `json:"d"`
}
