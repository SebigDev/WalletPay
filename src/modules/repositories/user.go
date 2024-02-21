package repositories

import (
	"CrashCourse/GoApp/src/modules/daos"
	"CrashCourse/GoApp/src/modules/entities"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoNoDocumentError string = "mongo: no documents in result"

type IUserRepository interface {
	AddUser(person daos.PersonDao) error
	GetUserByEmailAddress(emailAddress string) (*entities.Person, error)
	GetUserById(id string) (*entities.Person, error)
	GetUsers() (*[]entities.Person, error)
	UpdatePerson(person daos.PersonDao) error
}

type userRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepository(collection *mongo.Collection, ctx context.Context) IUserRepository {
	return &userRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (u *userRepository) AddUser(person daos.PersonDao) error {
	p, err := u.GetUserByEmailAddress(person.EmailAddress.Value)
	if err != nil {
		return err
	}

	if p.IsNotNil() {
		return errors.New("user with email address already exists")
	}

	_, err = u.collection.InsertOne(u.ctx, person)
	if err != nil {
		return err
	}

	log.Println("User added successfully")
	return nil
}

func (u *userRepository) GetUserByEmailAddress(emailAddress string) (*entities.Person, error) {
	var dao daos.PersonDao

	filter := bson.M{
		"emailAddress": bson.M{
			"value": emailAddress,
		}}

	err := u.collection.FindOne(u.ctx, filter).Decode(&dao)

	if err != nil {
		if err.Error() == MongoNoDocumentError {
			return &entities.Person{}, nil
		}
		return &entities.Person{}, err
	}
	person := entities.MapFromDao(&dao)
	return &person, nil
}

func (u *userRepository) GetUserById(id string) (*entities.Person, error) {
	var dao daos.PersonDao

	filter := bson.M{"userId": id}
	err := u.collection.FindOne(u.ctx, filter).Decode(&dao)

	if err != nil {
		if err.Error() == MongoNoDocumentError {
			return &entities.Person{}, nil
		}
		return &entities.Person{}, err
	}
	person := entities.MapFromDao(&dao)
	return &person, nil
}

func (u *userRepository) GetUsers() (*[]entities.Person, error) {

	cursor, err := u.collection.Find(u.ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var daos []daos.PersonDao
	var persons []entities.Person

	if err = cursor.All(context.TODO(), &daos); err != nil {
		log.Fatal(err)
	}

	for _, dao := range daos {
		persons = append(persons, entities.MapFromDao(&dao))
	}
	return &persons, nil
}

func (u *userRepository) UpdatePerson(person daos.PersonDao) error {
	filter := bson.M{"userId": person.UserId}

	_, err := u.collection.ReplaceOne(u.ctx, filter, person)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
