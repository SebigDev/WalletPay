package dto

type DepositRequest struct {
	Amount       float64 `json:"amount"`
	WalletNumber string  `json:"walletNo"`
	Currency     string  `json:"currency"`
}

type WithdrawRequest struct {
	Amount       float64 `json:"amount"`
	WalletNumber string  `json:"walletNo"`
	Currency     string  `json:"currency"`
}

type CreateWalletRequest struct {
	Currency string `json:"currency"`
}
