package entites

type PersonReniec struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    DataPerson `json:"data"`
}

type DataPerson struct {
	Verificador     int8   `json:"verificador"`
	Dni             string `json:"dni"`
	NombreCompleto  string `json:"nombre_completo"`
	Nombres         string `json:"nombres"`
	ApellidoPaterno string `json:"ap_paterno"`
	ApellidoMaterno string `json:"ap_materno"`
	ApellidoCasada  string `json:"ap_casada"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Ubigeo          string `json:"ubigeo"`
	Direccion       string `json:"direccion"`
}

type DataQuery struct {
	DniQuery string `json:"dni_query"`
}