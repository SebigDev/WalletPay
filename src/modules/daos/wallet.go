package daos

import (
	"github.com/google/uuid"
)

type AmountDao struct {
	Value    float64 `bson:"value"`
	Currency string  `bson:"currency"`
}

type WalletDao struct {
	Number string    `bson:"number"`
	Amount AmountDao `bson:"amount"`
	Owner  uuid.UUID `bson:"owner"`
	Type   string    `bson:"type"`
}
