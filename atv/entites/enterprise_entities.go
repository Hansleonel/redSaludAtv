package entites

type EnterpriseSunat struct {
	Success bool           `json:"success"`
	Data    DataEnterprise `json:"data"`
}

type DataEnterprise struct {
	Ruc         string   `json:"ruc"`
	RazonSocial string   `json:"nombre_o_razon_social"`
	Condicion   string   `json:"condicion"`
	Estado      string   `json:"estado"`
	Direccion   string   `json:"direccion_completa"`
	Ubigeo      []string `json:"ubigeo"`
}

type DataQueryRuc struct {
	RucQuery string `json:"ruc_query"`
}
