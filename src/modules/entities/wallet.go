package entities

import (
	"CrashCourse/GoApp/src/modules/daos"
	r "CrashCourse/GoApp/src/modules/responses"
	"CrashCourse/GoApp/src/modules/vo"
	"fmt"
	"math/rand"
)

type WalletError struct {
	ErrorMsg string
}

func (e WalletError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

type WalletType string

type Money struct {
	Amount   vo.Amount
	Currency vo.Currency
}

const (
	Private  WalletType = "Private"
	Business WalletType = "Business"
)

type Wallet struct {
	Number  string
	Money   Money
	OwnerID vo.Owner
	Type    WalletType
}

func NewMoney(amount vo.Amount, currency vo.Currency) *Money {
	return &Money{
		Amount:   amount,
		Currency: currency,
	}
}

func NewWallet(owner vo.Owner, money Money) *Wallet {
	return &Wallet{
		Number:  generateWalletNumber(),
		OwnerID: owner,
		Type:    Private,
		Money: Money{
			Amount:   money.Amount,
			Currency: money.Currency,
		},
	}
}

func (w *Wallet) Deposit(money Money) {
	w.Money.Amount += money.Amount
}

func (w *Wallet) Withdraw(money Money) {
	w.Money.Amount -= money.Amount
}

func WalletsToDao(wallets []Wallet) *[]daos.WalletDao {
	var walletDao []daos.WalletDao
	for _, w := range wallets {
		walletDao = append(walletDao, daos.WalletDao{
			Number: w.Number,
			Amount: daos.AmountDao{
				Value:    float64(w.Money.Amount),
				Currency: w.Money.Currency.String(),
			},
			Owner: w.OwnerID.String(),
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
				Value:    float64(wallet.Money.Amount),
				Currency: wallet.Money.Currency.String(),
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
			Number:  wallet.Number,
			OwnerID: vo.Owner(wallet.Owner),
			Money: Money{
				Amount:   vo.Amount(wallet.Amount.Value),
				Currency: vo.Currency(wallet.Amount.Currency),
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
			Value:    float64(w.Money.Amount),
			Currency: w.Money.Currency.String(),
		},
		Owner: w.OwnerID.String(),
		Type:  string(w.Type),
	}
}

func generateWalletNumber() string {
	return fmt.Sprintf("NL%02dGOAPP%08d", rand.Int63n(10), rand.Int63n(100000000))
}
