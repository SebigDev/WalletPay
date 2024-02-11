package services

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/user/dto"
	"CrashCourse/GoApp/src/modules/user/entities"
	"CrashCourse/GoApp/src/modules/user/repositories"
	"CrashCourse/GoApp/src/modules/user/responses"
	"CrashCourse/GoApp/src/modules/user/vo"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IUserService interface {
	CreateNewPerson(data dto.CreatePerson) error
	LoginPerson(req dto.LoginRequest) (responses.AuthResponse, error)
	GetPersonById(id string) (responses.PersonResponse, error)
	GetAllUsers() ([]responses.PersonResponse, error)
}

type userService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (u *userService) CreateNewPerson(data dto.CreatePerson) error {
	person, err := entities.NewPerson(data)
	if err != nil {
		return err
	}
	err = u.UserRepository.AddUser(person)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetPersonById(id string) (responses.PersonResponse, error) {

	person, err := u.UserRepository.GetUserById(id)

	if err != nil {
		return responses.PersonResponse{}, err
	}

	return person.MapToResponse(), nil
}

func (u *userService) GetAllUsers() ([]responses.PersonResponse, error) {

	daos, err := u.UserRepository.GetUsers()

	if err != nil {
		return []responses.PersonResponse{}, err
	}

	var persons []responses.PersonResponse

	for _, p := range *daos {
		persons = append(persons, p.MapToResponse())
	}
	return persons, nil
}

func (u *userService) LoginPerson(req dto.LoginRequest) (responses.AuthResponse, error) {
	person, err := u.UserRepository.GetUserByEmailAddress(req.EmailAddress)

	if err != nil {
		return responses.AuthResponse{}, errors.New("hi there! You have provided an invalid credentials")
	}

	if err := vo.VerifyPassword(req.Password, person.Password.Value); err != nil {
		return responses.AuthResponse{}, err
	}

	hasher := sha256.New()
	_, err = hasher.Write([]byte(fmt.Sprintf("%+v", req)))

	if err != nil {
		log.Println("hash err:", err)
		return responses.AuthResponse{}, err
	}

	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["hash"] = hashString
	claims["sub"] = person.UserId.String()
	claims["exp"] = time.Now().UTC().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte("SecretKey"))

	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.AuthResponse{
		Token:        tokenString,
		RefreshToken: utils.GenerateSecureToken(),
	}, nil
}
