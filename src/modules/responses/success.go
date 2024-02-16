package responses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ResponseType interface {
	string | int64 | float32 | float64 | AuthResponse | PersonResponse | WalletResponse | []PersonResponse
}

type ApiSuccessResponse[T ResponseType] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponse[T ResponseType](data T) *ApiSuccessResponse[T] {
	return &ApiSuccessResponse[T]{
		Data:    Transform(data),
		Message: "Success",
	}
}

func Transform[T any](data T) T {
	switch any(data).(type) {
	case string:
		return capitalise[string](any(data).(string)).(T)
	default:
		return data
	}
}

func capitalise[T string](s string) any {
	if strings.HasPrefix(s, "eyJh") { //token string should not be capitalised
		return s
	}
	str := cases.Title(language.English, cases.NoLower).String(s)
	return str
}
