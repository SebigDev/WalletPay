package dto

type DepositRequest struct {
	Amount       float64 `json:"amount"`
	WalletNumber string  `json:"walletNo"`
}

type WithdrawRequest struct {
	Amount       float64 `json:"amount"`
	WalletNumber string  `json:"walletNo"`
}
