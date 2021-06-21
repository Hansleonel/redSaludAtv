package payments

type RecurrencePayU struct {
	Quantity     string   `json:"quantity"`
	Installments string   `json:"installments"`
	TrialDays    string   `json:"trialDays"`
	Customer     Customer `json:"customer"`
	Plan         Plan     `json:"plan"`
}

type Customer struct {
	Id          string        `json:"id"`
	CreditCards []CreditCards `json:"creditCards"`
}

type CreditCards struct {
	Token string `json:"token"`
}

type Plan struct {
	PlanCode string `json:"planCode"`
}
