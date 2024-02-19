package vo

import (
	"fmt"
	"strings"
)

type Owner string
type Amount float64
type Currency string
type WalletNumber string
type CreditorName string
type MoneyConverted float64

const (
	EURO Currency = "EUR"
	USD  Currency = "USD"
	GBP  Currency = "GBP"
)

type TypeError struct {
	ErrorMsg string
}

func (e TypeError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

func NewMoneyConverted(amount Amount, to, from Currency) MoneyConverted {
	mult := demoConvert(to, from)
	return MoneyConverted(amount * Amount(mult))
}

func NewOwner(userId string) (Owner, error) {
	if len(userId) == 0 {
		return "", TypeError{ErrorMsg: "Invalid owner"}
	}
	return Owner(userId), nil
}

func NewCurrency(currency string) (Currency, error) {
	if len(currency) == 0 {
		return "", TypeError{ErrorMsg: "Currency type must be provided"}
	}
	if currency == string(EURO) || currency == string(GBP) {
		return Currency(currency), nil
	}
	return "", TypeError{ErrorMsg: "Invalid currency type indicated"}
}

func NewWalletNumber(walletNo string) (WalletNumber, error) {
	walletNumber := strings.TrimSpace(walletNo)
	if len(walletNumber) == 0 {
		return "", TypeError{ErrorMsg: "Wallet number cannot be empty"}
	}
	return WalletNumber(walletNo), nil
}

func NewAmount(amount float64, isNew bool) (Amount, error) {
	if amount <= 0 && !isNew {
		return 0, TypeError{ErrorMsg: "Amount has to a postive number"}
	}
	return Amount(amount), nil
}

func NewCreditor(name string) (CreditorName, error) {
	if len(name) == 0 || name == "" {
		return "", TypeError{ErrorMsg: "Creditor name must be provided"}
	}
	return CreditorName(name), nil
}

func (w *WalletNumber) String() string {
	return string(*w)
}

func (cr *CreditorName) String() string {
	return string(*cr)
}

func (c *Currency) String() string {
	return string(*c)
}

func (o *Owner) String() string {
	return string(*o)
}

func demoConvert(toCurrency, fromCurrency Currency) float64 {
	if (toCurrency == EURO && fromCurrency == EURO) || (toCurrency == GBP && fromCurrency == GBP) {
		return 1
	} else if toCurrency == GBP && fromCurrency == EURO {
		return 0.89
	} else {
		return 1.2
	}
}
