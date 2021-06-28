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

// TODO RESPONSES

type RecurrencePayUResponse struct {
	Id                 string             `json:"id"`
	Plan               PlanRecurrence     `json:"plan"`
	Customer           CustomerRecurrence `json:"customer"`
	Quantity           string             `json:"quantity"`
	Installments       string             `json:"installments"`
	CurrentPeriodStart int64              `json:"currentPeriodStart"`
	CurrentPeriodEnd   int64              `json:"currentPeriodEnd"`
	NotifyUrl          string             `json:"notifyUrl"`
}

type PlanRecurrence struct {
	Id               string                       `json:"id"`
	PlanCode         string                       `json:"planCode"`
	Description      string                       `json:"description"`
	AccountId        string                       `json:"accountId"`
	IntervalCount    int16                        `json:"intervalCount"`
	AdditionalValues []AdditionalValuesRecurrence `json:"additionalValues"`
}

type AdditionalValuesRecurrence struct {
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type CustomerRecurrence struct {
	Id          string                  `json:"id"`
	FullName    string                  `json:"fullName"`
	Email       string                  `json:"email"`
	CreditCards []CreditCardsRecurrence `json:"creditCards"`
}

type CreditCardsRecurrence struct {
	Token string `json:"token"`
}
