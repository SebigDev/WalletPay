package entities

import (
	"CrashCourse/GoApp/src/modules/user/daos"
	"CrashCourse/GoApp/src/modules/user/dto"
	"CrashCourse/GoApp/src/modules/user/vo"
	w "CrashCourse/GoApp/src/modules/wallet/daos"
	"CrashCourse/GoApp/src/modules/wallet/entities"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	user        *User
	houseNumber string
	streetName  string
	postalCode  string
	city        string
	wallets     []entities.Wallet
}

func NewPerson(p dto.CreatePerson) (daos.PersonDao, error) {
	agg := &UserAggregate{
		UserId: uuid.New(),
	}
	emailAddress, err := vo.CreateEmailAddress(p.EmailAddress)
	if err != nil {
		return daos.PersonDao{}, err
	}
	password, err := vo.CreatePassword(p.Password)
	if err != nil {
		return daos.PersonDao{}, err
	}

	user := &User{
		Aggregate:    agg,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		EmailAddress: emailAddress,
		Password:     password,
		CreateAt:     time.Now().UTC(),
		IsActive:     true,
		IsVerified:   true,
	}

	person := Person{
		user:        user,
		houseNumber: p.HouseNumber,
		streetName:  p.StreetName,
		postalCode:  p.PostalCode,
		city:        p.City,
		wallets:     []entities.Wallet{},
	}
	return person.mapToDao(), nil
}

func (p *Person) mapToDao() daos.PersonDao {
	return daos.PersonDao{
		UserId:       p.user.Aggregate.UserId,
		FirstName:    p.user.FirstName,
		LastName:     p.user.LastName,
		EmailAddress: daos.EmailAddressDao(p.user.EmailAddress),
		Password:     daos.PasswordDao(p.user.Password),
		HouseNumber:  p.houseNumber,
		PostalCode:   p.postalCode,
		StreetName:   p.streetName,
		City:         p.city,
		IsActive:     p.user.IsActive,
		IsVerified:   p.user.IsVerified,
		CreatedAt:    p.user.CreateAt,
		Wallets:      *w.MapToDaoFrom(p.wallets),
	}
}
