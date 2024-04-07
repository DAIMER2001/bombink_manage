package api

type RetomaRegistroResponseErrorInexperado struct {
	Message string `json:"message" example:"Ocurrio un error inexperado"`
	Data    Data   `json:"data"`
}

type RetomaRegistroResponseUIDVacio struct {
	Message string `json:"message" example:"Datos enviados inválidos"`
	Data    Data   `json:"data"`
}

type Data struct {
	Code string `json:"code" example:"IDMissionNoValid"`
}
