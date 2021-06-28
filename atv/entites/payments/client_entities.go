package payments

type ClientPayU struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type ClientPayUResponse struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Id       string `json:"id"`
}
