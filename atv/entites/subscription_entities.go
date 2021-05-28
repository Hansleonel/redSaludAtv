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
	Estado            int8   `json:"subscription_estado"`           // 1ro
	TypeProcesamiento int8   `json:"subscription_typeProcesamiento"`
	FechaRegistro     string `json:"subscription_fechaRegistro"`
	IdPlan            int64  `json:"subscription_idPlan"`
	FrecuenciaPago    int8   `json:"subscription_frecuenciaPago"`
	TipoAfiliacion    int8   `json:"subscription_tipoAfiliacion"`
}

type SubscriptionStepOne struct {
	Id              int64  `json:"subscription_id"`
	FechaNacimiento string `json:"subscription_fechaNacimiento"`
	CelNumber       string `json:"subscription_celNumber"`
	Type            int    `json:"subscription_type"`
	TcDatos         string `json:"subscription_tcDatos"`
	IdPlan          int64  `json:"subscription_idPlan"`
	FrecuenciaPago  int8   `json:"subscription_frecuenciaPago"`
	TipoAfiliacion  int8   `json:"subscription_tipoAfiliacion"`
}

type SubscriptionStepTwo struct {
	Id     int64  `json:"subscription_id"`
	NumDoc string `json:"subscription_numDoc"`
	Ape1   string `json:"subscription_ape1"`
	Ape2   string `json:"subscription_ape2"`
	Nom1   string `json:"subscription_nom1"`
	Nom2   string `json:"subscription_nom2"`
	Mail   string `json:"subscription_mail"`
}

type SubscriptionStepThree struct {
	DecJur         string `json:"declaracion_jurada"`
	QuestionFirst  string `json:"question_01"`
	QuestionSecond string `json:"question_02"`
	QuestionThird  string `json:"question_03"`
}

type SubsError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
