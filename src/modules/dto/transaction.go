package dto

type CreateTransaction struct {
	Amount                float64 `json:"amount"`
	CreditorCurrency      string  `json:"creditorCurrency"`
	DebitorCurrency       string  `json:"debitorCurrency"`
	CreditorWalletAddress string  `json:"creditorWalletAddress"`
	DebitorWalletAddress  string  `json:"debitorWalletAddress"`
	Description           string  `json:"description"`
	CreditorName          string  `json:"creditorName"`
	Pin                   string  `json:"pin"`
}

type CreatePayRequest struct {
	CreditorWallet string  `json:"creditorWallet"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	RequestPartyId string  `json:"requestPartyId"`
	Pin            string  `json:"pin"`
}
type AckRequest struct {
	RequestId string `json:"requestId"`
}
