package entities

import (
	"CrashCourse/GoApp/src/modules/daos"
	r "CrashCourse/GoApp/src/modules/responses"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type WalletError struct {
	ErrorMsg string
}

func (e WalletError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

type WalletType string

type Amount struct {
	Value    float64
	Currency string
}

const (
	Private  WalletType = "Private"
	Business WalletType = "Business"
)

type Currency string

const (
	EURO Currency = "EUR"
	USD  Currency = "USD"
	YEN  Currency = "YEN"
	NGN  Currency = "NGN"
)

type Wallet struct {
	Number string
	Amount Amount
	Owner  uuid.UUID
	Type   WalletType
}

func NewWallet(userId string) (*Wallet, error) {
	owner, err := uuid.Parse(userId)
	if err != nil {
		return nil, WalletError{ErrorMsg: "Wallet cannot e created"}
	}
	return &Wallet{
		Number: generateWalletNumber(),
		Owner:  owner,
		Type:   Private,
		Amount: Amount{
			Value:    0,
			Currency: string(EURO),
		},
	}, nil
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return WalletError{ErrorMsg: "Amount has to be greater than 0"}
	}
	w.Amount.Value += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.Amount.Value {
		return WalletError{ErrorMsg: "You have insufficient balance"}
	}
	w.Amount.Value -= amount
	return nil
}

func WalletsToDao(wallets []Wallet) *[]daos.WalletDao {
	var walletDao []daos.WalletDao
	for _, w := range wallets {
		walletDao = append(walletDao, daos.WalletDao{
			Number: w.Number,
			Amount: daos.AmountDao{
				Value:    w.Amount.Value,
				Currency: w.Amount.Currency,
			},
			Owner: w.Owner,
			Type:  string(w.Type),
		})
	}
	return &walletDao
}

func PersonToResponse(wallets []Wallet) *[]r.WalletResponse {
	var res []r.WalletResponse
	for _, wallet := range wallets {
		res = append(res, r.WalletResponse{
			Number: wallet.Number,
			Amount: r.Amount{
				Value:    float64(wallet.Amount.Value),
				Currency: wallet.Amount.Currency,
			},
			Type: string(wallet.Type),
		})
	}
	return &res
}

func FromDao(wallets []daos.WalletDao) *[]Wallet {
	var res []Wallet
	for _, wallet := range wallets {
		res = append(res, Wallet{
			Number: wallet.Number,
			Owner:  wallet.Owner,
			Amount: Amount{
				Value:    wallet.Amount.Value,
				Currency: wallet.Amount.Currency,
			},
			Type: WalletType(wallet.Type),
		})
	}
	return &res
}

func WalletToDao(w *Wallet) *daos.WalletDao {
	return &daos.WalletDao{
		Number: w.Number,
		Amount: daos.AmountDao{
			Value:    w.Amount.Value,
			Currency: w.Amount.Currency,
		},
		Owner: w.Owner,
		Type:  string(w.Type),
	}
}

func generateWalletNumber() string {
	return fmt.Sprintf("NL%02dGOAPP%08d", rand.Int63n(10), rand.Int63n(100000000))
}
