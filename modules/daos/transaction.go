package daos

import "time"

type TransactionDao struct {
	ID                   string    `bson:"id"`
	ReceiverAccount      ToDao     `bson:"receiverAccount"`
	SenderAccount        FromDao   `bson:"senderAccount"`
	Amount               float64   `bson:"amount"`
	Currency             string    `bson:"currency"`
	Description          string    `bson:"description"`
	CreatedAt            time.Time `bson:"createdAt"`
	Sender               string    `bson:"sender"`
	Receiver             string    `bson:"receiver"`
	TransactionReference string    `bson:"transactionReference"`
}

type ToDao struct {
	ToName         string `bson:"toName"`
	ToWalletNumber string `bson:"toWalletNumber"`
	ToCurrency     string `bson:"toCurrency"`
}

type FromDao struct {
	FromName         string `bson:"fromName"`
	FromWalletNumber string `bson:"fromWalletNumber"`
	FromCurrency     string `bson:"fromCurrency"`
}
