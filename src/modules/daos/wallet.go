package daos

type AmountDao struct {
	Value    float64 `bson:"value"`
	Currency string  `bson:"currency"`
}

type WalletDao struct {
	Number string    `bson:"number"`
	Amount AmountDao `bson:"amount"`
	Owner  string    `bson:"owner"`
	Type   string    `bson:"type"`
}
