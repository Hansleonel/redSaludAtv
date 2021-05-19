package entites

type Subscription struct {
	Id                int64  `json:"subscription_id"`
	TypeDocument      int16  `json:"subscription_typeDocument"`
	NumDoc            string `json:"subscription_numDoc"`
	Ape1              string `json:"subscription_ape1"`
	Ape2              string `json:"subscription_ape2"`
	Nom1              string `json:"subscription_nom1"`
	Nom2              string `json:"subscription_nom2"`
	FechaNacimiento   string `json:"subscription_fechaNacimiento"` // 1ro
	Mail              string `json:"subscription_mail"`
	CelNumber         string `json:"subscription_celNumber"` // 1ro
	Type              int    `json:"subscription_type"`      // 1ro
	IdContratante     string `json:"subscription_idContratante"`
	TcDatos           string `json:"subscription_tcDatos"`          // 1ro
	TcComunicaciones  string `json:"subscription_tcComunicaciones"` // 1ro
	TcPagos           string `json:"subscription_tcPagos"`          // 1ro
	Estado            string `json:"subscription_estado"`           // 1ro
	TypeProcesamiento int8   `json:"subscription_typeProcesamiento"`
	FechaRegistro     string `json:"subscription_fechaRegistro"`
	IdPlan            int64  `json:"subscription_idPlan"`
	FrecuenciaPago    int8   `json:"subscription_frecuenciaPago"`
	TipoAfiliacion    int8   `json:"subscription_tipoAfiliacion"`
}
