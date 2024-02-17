package repositories

import (
	"CrashCourse/GoApp/src/modules/daos"
	"CrashCourse/GoApp/src/modules/entities"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransactionRepository interface {
	Submit(trx entities.Transaction) error
	GetTransaction(userId string) (*[]daos.TransactionDao, error)
}

type TransactionRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewTransactionRepository(collection *mongo.Collection, ctx context.Context) ITransactionRepository {
	return &TransactionRepository{
		collection: collection,
		context:    ctx,
	}
}

func (tx *TransactionRepository) Submit(trx entities.Transaction) error {
	res, err := tx.collection.InsertOne(tx.context, trx.Create())
	if err != nil {
		return fmt.Errorf("an error occurred submitting transaction: %w", err)
	}
	log.Printf("Payment submitted successfully : %s", res.InsertedID)
	return nil
}

func (tx *TransactionRepository) GetTransaction(userId string) (*[]daos.TransactionDao, error) {
	cursor, err := tx.collection.Find(tx.context, bson.M{"userId": userId})
	if err != nil {
		log.Fatal(err)
	}

	var transactions []daos.TransactionDao

	if err = cursor.All(context.TODO(), &transactions); err != nil {
		log.Fatal(err)
	}

	return &transactions, nil
}
