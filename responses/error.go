package responses

type ErrorResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}
