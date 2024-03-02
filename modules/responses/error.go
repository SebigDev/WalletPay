package responses

type ApiErrorResponse struct {
	Error  string `json:"message"`
	Status string `json:"status"`
}

func CreateErrorResponse(message string) *ApiErrorResponse {
	return &ApiErrorResponse{
		Error:  message,
		Status: "Error",
	}
}
