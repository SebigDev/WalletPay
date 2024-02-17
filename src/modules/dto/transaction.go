package dto

type CreateTransaction struct {
	Amount                float64 `json:"amount"`
	CreditorCurrency      string  `json:"creditorCurrency"`
	DebitorCurrency       string  `json:"debitorCurrency"`
	CreditorWalletAddress string  `json:"creditorWalletAddress"`
	DebitorWalletAddress  string  `json:"debitorWalletAddress"`
	Description           string  `json:"description"`
	CreditorName          string  `json:"creditorName"`
}
