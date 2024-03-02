package entities

import (
	"fmt"
	"time"

	"github.com/sebigdev/walletpay/modules/daos"
	"github.com/sebigdev/walletpay/modules/dto"
	"github.com/sebigdev/walletpay/modules/responses"
	"github.com/sebigdev/walletpay/modules/vo"

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

func NewPerson(p dto.CreatePerson) (*Person, error) {
	agg := &UserAggregate{
		UserId: uuid.NewString(),
	}
	emailAddress, err := vo.CreateEmailAddress(p.EmailAddress)
	if err != nil {
		return &Person{}, err
	}
	password, err := vo.CreatePassword(p.Password)
	if err != nil {
		return &Person{}, err
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
		Pin:          vo.Pin{},
	}

	person := Person{
		user:        user,
		houseNumber: p.HouseNumber,
		streetName:  p.StreetName,
		postalCode:  p.PostalCode,
		city:        p.City,
		wallets:     make([]Wallet, 0),
	}
	return &person, nil
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
		Pin: daos.PinDao{
			HashValue:    p.user.Pin.HashValue,
			RecoverValue: p.user.Pin.RecoverValue,
		},
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
			Pin:          vo.Pin(p.Pin),
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
		UserId:       d.user.Aggregate.UserId,
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

func (p *Person) Deposit(money vo.Money, walletNumber vo.WalletNumber) error {
	if !p.HasWallets() {
		return fmt.Errorf("oops!!! user has no wallets")
	}

	for i := range p.wallets {
		if p.wallets[i].Number == walletNumber.String() && p.wallets[i].Money.Currency == money.Currency {
			p.wallets[i].Deposit(money)
			return nil
		}
	}
	return fmt.Errorf("you have provided an invalid wallet address: %s", walletNumber)
}

func (p *Person) Withdraw(money vo.Money, walletNumber vo.WalletNumber) error {
	if !p.HasWallets() {
		return fmt.Errorf("oops!!! user has no wallets")
	}

	for i := range p.wallets {
		if p.wallets[i].Number == walletNumber.String() && p.wallets[i].Money.Currency == money.Currency {
			if p.wallets[i].Money.Amount < money.Amount {
				return fmt.Errorf("insufficient fund")
			}
			p.wallets[i].Withdraw(money)
			return nil
		}
	}
	return fmt.Errorf("you have provided an invalid wallet address: %s", walletNumber)
}

func (p *Person) ChangePassword(oldPass, newPass string) error {
	if err := vo.VerifyPassword(oldPass, p.user.Password.Value); err != nil {
		return err
	}
	password, err := vo.CreatePassword(newPass)
	if err != nil {
		return err
	}
	p.user.Password = password
	return nil
}

func (p *Person) ChangePin(oldPin, newPin vo.PinValue) error {
	if err := vo.Verify(oldPin.String(), p.user.Pin); err != nil {
		return err
	}
	nPin, err := vo.NewPinValue(newPin.String())
	if err != nil {
		return err
	}
	p.user.Pin = *vo.NewPin(nPin)
	return nil
}

func (p *Person) NewPin(pin string) {
	p.user.Pin = *vo.NewPin(vo.PinValue(pin))
}

func (p *Person) VerifyPassword(password string) error {
	return vo.VerifyPassword(password, p.user.Password.Value)
}

func (p *Person) VerifyPin(pin string) error {
	return vo.Verify(pin, p.user.Pin)
}

func (p *Person) VerifyPinRecovery(pin string) error {
	return vo.VerifyRecover(pin, p.user.Pin)
}

func (p *Person) GetUserID() string {
	return p.user.Aggregate.UserId
}

func (p *Person) GetRecoverPin() string {
	return p.user.Pin.RecoverValue
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
