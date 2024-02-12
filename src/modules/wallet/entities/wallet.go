package entities

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Wallet struct {
	Number  string
	Balance float64
	Owner   uuid.UUID
}

func NewWallet(owner uuid.UUID) *Wallet {
	return &Wallet{
		Number: generateWalletNumber(),
		Owner:  owner,
	}
}

func generateWalletNumber() string {
	return fmt.Sprintf("NL%dGOAPP%d", rand.Int63n(4), rand.Int63n(8))
}
