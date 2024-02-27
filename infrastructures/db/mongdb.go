package db

import (
	"context"
	"fmt"
	"log"

	"github.com/SebigDev/GoApp/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx    context.Context
	err    error
	client *mongo.Client
)

type MongoResponse struct {
	CtxR    context.Context
	ClientR *mongo.Client
}

func MongoInit() MongoResponse {
	mongoHost := config.GoEnv("MONGO_HOST")
	mongoPort := config.GoEnv("MONGO_PORT")
	mongoUri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return MongoResponse{
		CtxR:    ctx,
		ClientR: client,
	}
}
