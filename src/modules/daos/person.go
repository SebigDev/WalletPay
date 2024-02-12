package daos

import (
	"time"

	"github.com/google/uuid"
)

type PersonDao struct {
	UserId       uuid.UUID       `bson:"userId"`
	FirstName    string          `bson:"firstName"`
	LastName     string          `bson:"lastName"`
	EmailAddress EmailAddressDao `bson:"emailAddress"`
	Password     PasswordDao     `bson:"password"`
	CreatedAt    time.Time       `bson:"createdAt"`
	IsActive     bool            `bson:"isActive"`
	IsVerified   bool            `bson:"isVerified"`
	HouseNumber  string          `bson:"houseNumber"`
	StreetName   string          `bson:"streetName"`
	PostalCode   string          `bson:"postalCode"`
	City         string          `bson:"city"`
	Wallets      []WalletDao     `bson:"wallets"`
}

type EmailAddressDao struct {
	Value string `bson:"value"`
}

type PasswordDao struct {
	Value []byte `bson:"value"`
}
