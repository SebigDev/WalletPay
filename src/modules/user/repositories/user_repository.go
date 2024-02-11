package repositories

import (
	"CrashCourse/GoApp/src/modules/user/daos"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoNoDocumentError string = "mongo: no documents in result"

type IUserRepository interface {
	AddUser(person daos.PersonDao) error
	GetUserByEmailAddress(emailAddress string) (*daos.PersonDao, error)
	GetUserById(id string) (*daos.PersonDao, error)
	GetUsers() (*[]daos.PersonDao, error)
}

type userRepository struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewUserRepository(collection *mongo.Collection, ctx context.Context) IUserRepository {
	return &userRepository{
		Collection: collection,
		Ctx:        ctx,
	}
}

func (u *userRepository) AddUser(person daos.PersonDao) error {
	p, err := u.GetUserByEmailAddress(person.EmailAddress.Value)
	if err != nil {
		return err
	}

	if p.IsVerified {
		return errors.New("user with email address already exists")
	}

	_, err = u.Collection.InsertOne(u.Ctx, person)
	if err != nil {
		return err
	}

	log.Println("User added successfully")
	return nil
}

func (u *userRepository) GetUserByEmailAddress(emailAddress string) (*daos.PersonDao, error) {
	var dao daos.PersonDao

	filter := bson.M{
		"emailAddress": bson.M{
			"value": emailAddress,
		}}

	err := u.Collection.FindOne(u.Ctx, filter).Decode(&dao)

	if err != nil {
		if err.Error() == MongoNoDocumentError {
			return &daos.PersonDao{}, nil
		}
		return &daos.PersonDao{}, err
	}

	return &dao, nil
}

func (u *userRepository) GetUserById(id string) (*daos.PersonDao, error) {
	var dao daos.PersonDao

	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"userId": uuid}

	err = u.Collection.FindOne(u.Ctx, filter).Decode(&dao)

	if err != nil {
		if err.Error() == MongoNoDocumentError {
			return &daos.PersonDao{}, nil
		}
		return &daos.PersonDao{}, err
	}

	return &dao, nil
}

func (u *userRepository) GetUsers() (*[]daos.PersonDao, error) {

	cursor, err := u.Collection.Find(u.Ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var daos []daos.PersonDao
	if err = cursor.All(context.TODO(), &daos); err != nil {
		log.Fatal(err)
	}

	return &daos, nil
}
