package entities

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type WalletType string

const (
	Private  WalletType = "Private"
	Business WalletType = "Business"
)

type Wallet struct {
	Number  string
	Balance float64
	Owner   uuid.UUID
	Type    WalletType
}

func NewWallet(owner uuid.UUID) *Wallet {
	return &Wallet{
		Number: generateWalletNumber(),
		Owner:  owner,
		Type:   Private,
	}
}

func generateWalletNumber() string {
	return fmt.Sprintf("NL%02dGOAPP%08d", rand.Int63n(10), rand.Int63n(100000000))
}
