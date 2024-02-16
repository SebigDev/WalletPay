package responses

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ApiErrorResponse struct {
	Error  string `json:"message"`
	Status string `json:"status"`
}

func CreateErrorResponse(message string) *ApiErrorResponse {
	return &ApiErrorResponse{
		Error:  cases.Title(language.English, cases.NoLower).String(message),
		Status: "Error",
	}
}
