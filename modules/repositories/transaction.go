package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/sebigdev/walletpay/modules/daos"
	"github.com/sebigdev/walletpay/modules/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransactionRepository interface {
	Submit(trx *entities.Transaction) error
	GetTransaction(userId string) (*[]daos.TransactionDao, error)
}

type transactionRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewTransactionRepository(collection *mongo.Collection, ctx context.Context) ITransactionRepository {
	return &transactionRepository{
		collection: collection,
		context:    ctx,
	}
}

func (tx *transactionRepository) Submit(trx *entities.Transaction) error {
	res, err := tx.collection.InsertOne(tx.context, trx.Create())
	if err != nil {
		return fmt.Errorf("an error occurred submitting transaction: %w", err)
	}
	log.Printf("Payment submitted successfully : %s", res.InsertedID)
	return nil
}

func (tx *transactionRepository) GetTransaction(userId string) (*[]daos.TransactionDao, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"receiver": userId},
			{"sender": userId},
		},
	}
	cursor, err := tx.collection.Find(tx.context, filter)
	if err != nil {
		log.Fatal(err)
	}

	var transactions []daos.TransactionDao

	if err = cursor.All(context.TODO(), &transactions); err != nil {
		log.Fatal(err)
	}

	return &transactions, nil
}
