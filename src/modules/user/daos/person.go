package daos

import (
	"CrashCourse/GoApp/src/modules/user/responses"
	"CrashCourse/GoApp/src/modules/wallet/daos"
	"time"

	"github.com/google/uuid"
)

type PersonDao struct {
	UserId       uuid.UUID        `bson:"userId"`
	FirstName    string           `bson:"firstName"`
	LastName     string           `bson:"lastName"`
	EmailAddress EmailAddressDao  `bson:"emailAddress"`
	Password     PasswordDao      `bson:"password"`
	CreatedAt    time.Time        `bson:"createdAt"`
	IsActive     bool             `bson:"isActive"`
	IsVerified   bool             `bson:"isVerified"`
	HouseNumber  string           `bson:"houseNumber"`
	StreetName   string           `bson:"streetName"`
	PostalCode   string           `bson:"postalCode"`
	City         string           `bson:"city"`
	Wallets      []daos.WalletDao `bson:"wallets"`
}

type EmailAddressDao struct {
	Value string `bson:"value"`
}

type PasswordDao struct {
	Value []byte `bson:"value"`
}

func (d *PersonDao) MapToResponse() responses.PersonResponse {
	return responses.PersonResponse{
		UserId:       d.UserId.String(),
		FirstName:    d.FirstName,
		LastName:     d.LastName,
		EmailAddress: d.EmailAddress.Value,
		HouseNumber:  d.HouseNumber,
		PostalCode:   d.PostalCode,
		StreetName:   d.StreetName,
		City:         d.City,
		IsActive:     d.IsActive,
		IsVerified:   d.IsActive,
		CreatedAt:    d.CreatedAt,
		Wallets:      *daos.MapDaoToResponse(d.Wallets),
	}
}
