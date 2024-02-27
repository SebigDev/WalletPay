package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/SebigDev/GoApp/src/modules/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPaymentRequestRepository interface {
	MakeRequest(request *entities.PayRequest) (string, error)
	AcknowlegeRequest(requestId string) error
	DeclineRequest(requestId string) error
}

type paymentRequestRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewPaymentRequestRepository(collection *mongo.Collection, context context.Context) IPaymentRequestRepository {
	return &paymentRequestRepository{
		collection: collection,
		context:    context,
	}
}

func (pr *paymentRequestRepository) MakeRequest(request *entities.PayRequest) (string, error) {
	res, err := pr.collection.InsertOne(pr.context, request)
	if err != nil {
		return "", fmt.Errorf("an error occurred submitting request: %w", err)
	}
	log.Printf("Request submitted successfully : %s", res.InsertedID)
	return request.Id, nil
}

func (pr *paymentRequestRepository) AcknowlegeRequest(requestId string) error {

	var request *entities.PayRequest
	filter := bson.M{"id": requestId}

	if err := pr.collection.FindOne(pr.context, filter).Decode(&request); err != nil {
		return fmt.Errorf("an error occurred retrieving request: %w", err)
	}

	request.AcknowledgeRequest()
	_, err := pr.collection.ReplaceOne(pr.context, filter, request)
	if err != nil {
		return fmt.Errorf("an error occurred acknowledging request: %w", err)
	}
	return nil
}

func (pr *paymentRequestRepository) DeclineRequest(requestId string) error {

	var request *entities.PayRequest
	filter := bson.M{"id": requestId}

	if err := pr.collection.FindOne(pr.context, filter).Decode(&request); err != nil {
		return fmt.Errorf("an error occurred retrieving request: %w", err)
	}

	request.DeclineRequest()
	_, err := pr.collection.ReplaceOne(pr.context, filter, request)
	if err != nil {
		return fmt.Errorf("an error occurred declining request: %w", err)
	}
	return nil
}
