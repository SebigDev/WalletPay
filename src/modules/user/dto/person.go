package dto

type CreatePerson struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	HouseNumber  string `json:"houseNumber"`
	StreetName   string `json:"streetName"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
}

type LoginRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
