package luggage

type MinimalQuoteRequest struct {
	From struct {
		TownCity    string `json:"TownCity"`
		PostZipCode string `json:"PostZipCode"`
		CountryCode string `json:"CountryCode"`
	} `json:"From"`
	To struct {
		TownCity    string `json:"TownCity"`
		PostZipCode string `json:"PostZipCode"`
		CountryCode string `json:"CountryCode"`
	} `json:"To"`
	Return    bool   `json:"Return"`
	Reference string `json:"Reference"`
}
