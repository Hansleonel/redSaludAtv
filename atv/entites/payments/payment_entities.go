package payments

type PaymentPayU struct {
	CompleteName string `json:"completeName"`
	Document     string `json:"document"`
	CardNumber   string `json:"cardNumber"`
	ExpMonth     string `json:"expMonth"`
	ExpYear      string `json:"expYear"`
	Type         string `json:"type"`
	Phone        string `json:"phone"`
	Mail         string `json:"mail"`
}
