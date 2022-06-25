package handlers

type JSONBadRequest struct {
	Message string `json:"message" example:"Bad Request"`
}

type JSONInternalServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}

type JSONStatusOK struct {
	Message string `json:"Message" example:"OK"`
}
