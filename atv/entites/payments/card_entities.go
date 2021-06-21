package payments

type cardPayU struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Number   string `json:"number"`
	ExpMonth string `json:"expMonth"`
	ExpYear  string `json:"expYear"`
	Type     string `json:"type"`
	Address Address `json:"address"`
}

type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	Line3      string `json:"line3"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
}