package entities

import (
	"CrashCourse/GoApp/src/modules/daos"
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/responses"
	"CrashCourse/GoApp/src/modules/vo"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	user        *User
	houseNumber string
	streetName  string
	postalCode  string
	city        string
	wallets     []Wallet
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
		EmailAddress: *emailAddress,
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
		wallets:     []Wallet{},
	}
	return person.MapToDao(), nil
}

func (p *Person) MapToDao() daos.PersonDao {
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
		Wallets:      *WalletsToDao(p.wallets),
	}
}

func MapFromDao(p *daos.PersonDao) Person {
	return Person{
		user: &User{
			Aggregate: &UserAggregate{
				UserId: p.UserId,
			},
			FirstName:    p.FirstName,
			LastName:     p.LastName,
			EmailAddress: vo.EmailAddress(p.EmailAddress),
			Password:     vo.Password(p.Password),
			IsActive:     p.IsActive,
			IsVerified:   p.IsVerified,
			CreateAt:     p.CreatedAt,
		},
		houseNumber: p.HouseNumber,
		postalCode:  p.PostalCode,
		streetName:  p.StreetName,
		city:        p.City,
		wallets:     *FromDao(p.Wallets),
	}
}

func MapToResponse(d *Person) responses.PersonResponse {
	return responses.PersonResponse{
		UserId:       d.user.Aggregate.UserId.String(),
		FirstName:    d.user.FirstName,
		LastName:     d.user.LastName,
		EmailAddress: d.user.EmailAddress.Value,
		HouseNumber:  d.houseNumber,
		PostalCode:   d.postalCode,
		StreetName:   d.streetName,
		City:         d.city,
		IsActive:     d.user.IsActive,
		IsVerified:   d.user.IsVerified,
		CreatedAt:    d.user.CreateAt,
		Wallets:      *PersonToResponse(d.wallets),
	}
}

func (p *Person) Deposit(amount float64, walletNo string) error {
	if !p.HasWallets() {
		return fmt.Errorf("oops!!! user has no wallets")
	}

	for _, wallet := range p.wallets {
		if wallet.Number == walletNo {
			err := wallet.Deposit(amount)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("you have provided an invalid wallet address: %s", walletNo)

}

func (p *Person) Withdraw(amount float64, walletNo string) error {
	if !p.HasWallets() {
		return fmt.Errorf("oops!!! user has no wallets")
	}

	for _, wallet := range p.wallets {
		if wallet.Number == walletNo {
			err := wallet.Withdraw(amount)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("you have provided an invalid wallet address: %s", walletNo)
}

func (p *Person) VerifyPassword(password string) error {
	return vo.VerifyPassword(password, p.user.Password.Value)
}

func (p *Person) GetUserID() string {
	return p.user.Aggregate.UserId.String()
}

func (p *Person) HasWallets() bool {
	return p.wallets != nil || len(p.wallets) > 0
}

func (p *Person) SetWallet(wallet Wallet) {
	p.wallets = append(p.wallets, wallet)
}

func (p *Person) IsNotNil() bool {
	return p.user != nil
}

func (p *Person) GetWallets() *[]Wallet {
	return &p.wallets
}
