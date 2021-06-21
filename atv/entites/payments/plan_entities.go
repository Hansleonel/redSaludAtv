package payments

type PlanPayU struct {
	AccountID            string                 `json:"accountID"`
	PlanCode             string                 `json:"planCode"`
	Description          string                 `json:"description"`
	Interval             string                 `json:"interval"`
	IntervalCount        string                 `json:"intervalCount"`
	MaxPaymentsAllowed   string                 `json:"maxPaymentsAllowed"`
	PaymentAttemptsDelay string                 `json:"paymentAttemptsDelay"`
	MaxPendingPayments   string                 `json:"maxPendingPayments"`
	MaxPaymentAttempts   string                 `json:"maxPaymentAttempts"`
	AdditionalValues     []AdditionalValuesPayU `json:"additionalValues"`
}

type AdditionalValuesPayU struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

