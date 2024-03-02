package responses

import "time"

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type PersonResponse struct {
	UserId       string           `json:"userId"`
	FirstName    string           `json:"firstName"`
	LastName     string           `json:"lastName"`
	EmailAddress string           `json:"emailAddress"`
	CreatedAt    time.Time        `json:"createdAt"`
	IsActive     bool             `json:"isActive"`
	IsVerified   bool             `json:"isVerified"`
	HouseNumber  string           `json:"houseNumber"`
	StreetName   string           `json:"streetName"`
	PostalCode   string           `json:"postalCode"`
	City         string           `json:"city"`
	Wallets      []WalletResponse `json:"wallets"`
}

type WalletResponse struct {
	Number string `json:"number"`
	Amount Amount `json:"balance"`
	Type   string `json:"type"`
}

type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type TransactionResponse struct {
	TransactionId        string `json:"transactionID"`
	TransactionReference string `json:"transactionReference"`
}

type TransactionFullResponse struct {
	ID                   string    `json:"id"`
	TransactionReference string    `json:"transactionReference"`
	BeneficiaryAccount   string    `json:"benficiaryAccount"`
	BeneficiaryCurrency  string    `json:"beneficiaryCurrency"`
	BeneficiaryName      string    `json:"beneficiaryName"`
	OriginatorAccount    string    `json:"originatorAccount"`
	OriginatorCurrency   string    `json:"originatorCurrency"`
	DebitorName          string    `json:"debitorName"`
	Amount               float64   `json:"amount"`
	Description          string    `json:"description"`
	CreatedAt            time.Time `json:"createdAt"`
	UserId               string    `json:"userId"`
}
