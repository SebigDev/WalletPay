package db


import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx      context.Context
	err      error
	client   *mongo.Client
	MongoUri string = "mongodb://localhost:27017/"
)

type MongoResponse struct {
	CtxR    context.Context
	ClientR *mongo.Client
}

func MongoInit() MongoResponse {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(MongoUri))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return MongoResponse{
		CtxR:    ctx,
		ClientR: client,
	}
}
