package payments

type PaymentGroupPayU struct {
	Data []PaymentPayU `json:"data"`
}

type PaymentPayU struct {
	CompleteName  string `json:"completeName"`
	Document      string `json:"document"`
	CardNumber    string `json:"cardNumber"`
	ExpMonth      string `json:"expMonth"`
	ExpYear       string `json:"expYear"`
	Type          string `json:"type"`
	Phone         string `json:"phone"`
	Mail          string `json:"mail"`
}

type PaymentClient struct {
	IdClient int64 `json:"idClient"`
}
