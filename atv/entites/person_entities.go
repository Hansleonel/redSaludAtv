package entites

type PersonReniec struct {
	Success bool       `json:"success"`
	Message int8       `json:"origen"`
	Data    DataPerson `json:"data"`
}

type DataPerson struct {
	Dni                string `json:"numero"`
	Nombres            string `json:"nombres"`
	ApellidoPaterno    string `json:"apellido_paterno"`
	ApellidoMaterno    string `json:"apellido_materno"`
	Sexo               string `json:"sexo"`
	CodigoVerificacion string `json:"codigo_verificacion"`
	FechaNacimiento    string `json:"fecha_nacimiento"`
	NombreCompleto     string `json:"nombre_completo"`
}

type DataQuery struct {
	DniQuery string `json:"dni_query"`
}

type ValidationPerson struct {
	DniValidation string `json:"dni_validation"`
	Validate      bool   `json:"validate"`
}
