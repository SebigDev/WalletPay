package responses

type ResponseType interface {
	string | int64 | float32 | float64 | AuthResponse | PersonResponse | WalletResponse | TransactionResponse | []PersonResponse | []TransactionFullResponse
}

type ApiSuccessResponse[T ResponseType] struct {
	Status string `json:"message"`
	Data   T      `json:"data"`
}

func CreateResponse[T ResponseType](data T) *ApiSuccessResponse[T] {
	return &ApiSuccessResponse[T]{
		Data:   data,
		Status: "Success",
	}
}
