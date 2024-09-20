package db

import (
	"context"
	"fmt"
	"log"

	"github.com/sebigdev/walletpay/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx    context.Context
	err    error
	client *mongo.Client
)

type DbProps struct {
	Context               context.Context
	UserCollection        *mongo.Collection
	PaymentCollection     *mongo.Collection
	TransactionCollection *mongo.Collection
}

func Init() *DbProps {
	mongoHost := config.GoEnv("MONGO_HOST")
	mongoPort := config.GoEnv("MONGO_PORT")
	mongoUri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	database := client.Database("walletpay")
	userCollection := database.Collection("users")
	trxCollection := database.Collection("transactions")
	payReqCollection := database.Collection("payments")

	log.Println("Connected to MongoDB")

	opts := &DbProps{
		Context:               ctx,
		UserCollection:        userCollection,
		TransactionCollection: trxCollection,
		PaymentCollection:     payReqCollection,
	}
	return opts
}
