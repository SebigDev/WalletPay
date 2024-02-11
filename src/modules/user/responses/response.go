package responses

import "time"

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type PersonResponse struct {
	UserId       string    `json:"userId"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	EmailAddress string    `json:"emailAddress"`
	CreatedAt    time.Time `json:"createdAt"`
	IsActive     bool      `json:"isActive"`
	IsVerified   bool      `json:"isVerified"`
	HouseNumber  string    `json:"houseNumber"`
	StreetName   string    `json:"streetName"`
	PostalCode   string    `json:"postalCode"`
	City         string    `json:"city"`
}
