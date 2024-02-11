package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/user/dto"
	"CrashCourse/GoApp/src/modules/user/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	CreatePerson(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	GetPersonById(ctx *fiber.Ctx) error
}

type userHandler struct {
	UserService services.IUserService
}

func NewUserHandler(userService services.IUserService) IUserHandler {
	return &userHandler{
		UserService: userService,
	}
}

func (s userHandler) CreatePerson(ctx *fiber.Ctx) error {
	createPersonReqBody := new(dto.CreatePerson)

	if err := ctx.BodyParser(createPersonReqBody); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Error parsing request")
	}

	err := s.UserService.CreateNewPerson(*createPersonReqBody)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Person created successfully"})
}

func (s userHandler) GetAllUsers(ctx *fiber.Ctx) error {

	persons, err := s.UserService.GetAllUsers()

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(persons)
}

func (s userHandler) GetPersonById(ctx *fiber.Ctx) error {

	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Error parsing request")
	}

	person, err := s.UserService.GetPersonById(userId)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(person)
}