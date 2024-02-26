package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/responses"
	"CrashCourse/GoApp/src/modules/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	CreatePerson(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	GetPersonById(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	ChangePin(ctx *fiber.Ctx) error
}

type userHandler struct {
	UserService services.IUserService
}

func NewUserHandler(userService services.IUserService) IUserHandler {
	return &userHandler{
		UserService: userService,
	}
}

// CreatePerson func for creates a new person.
// @Description Create a new person.
// @Summary create a new person
// @Tags Person
// @Accept json
// @Produce json
// @Param person body dto.CreatePerson true "Create person"
// @Success 200
// @Router /api/v1/users/onboard [post]
func (s userHandler) CreatePerson(ctx *fiber.Ctx) error {
	createPersonReqBody := new(dto.CreatePerson)

	if err := ctx.BodyParser(createPersonReqBody); err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err := s.UserService.CreateNewPerson(*createPersonReqBody)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Person created successfully")
	return ctx.JSON(response)
}

// GetAllUser func gets all users.
// @Description Get all users.
// @Summary get all users
// @Tags Person
// @Accept json
// @Produce json
// @Success 200 {array} responses.PersonResponse
// @Router /api/v1/users [get]
func (s userHandler) GetAllUsers(ctx *fiber.Ctx) error {

	persons, err := s.UserService.GetAllUsers()

	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse(persons)
	return ctx.JSON(response)
}

// GetPerson func gets person by given ID or 404 error.
// @Description Get person by given ID.
// @Tags Person
// @Summary get user by ID
// @Accept json
// @Produce json
// @Success 200 {object} responses.PersonResponse
// @Router /api/v1/user [get]
func (s userHandler) GetPersonById(ctx *fiber.Ctx) error {

	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	person, err := s.UserService.GetPersonById(userId)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse(person)
	return ctx.JSON(response)
}

// Change Password func changes password for a user.
// @Description Change password.
// @Tags Person
// @Summary change user's password
// @Accept json
// @Produce json
// @Param person body dto.CreatePasswordChangeRequest true "Change Password"
// @Success 200
// @Router /api/v1/user/change-password [post]
func (s userHandler) ChangePassword(ctx *fiber.Ctx) error {

	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	passwordReq := new(dto.CreatePasswordChangeRequest)

	if err := ctx.BodyParser(passwordReq); err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err = s.UserService.ChangePassword(userId, *passwordReq)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Password Changed successfully.")
	return ctx.JSON(response)
}

// Change Pin func changes password for a user.
// @Description Change pin.
// @Tags Person
// @Summary change user's pin
// @Accept json
// @Produce json
// @Param person body dto.CreatePinChangeRequest true "Change Pin"
// @Success 200
// @Router /api/v1/user/change-pin [post]
func (s userHandler) ChangePin(ctx *fiber.Ctx) error {

	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	pinRequest := new(dto.CreatePinChangeRequest)

	if err := ctx.BodyParser(pinRequest); err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err = s.UserService.ChangePin(userId, *pinRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Pin Changed successfully.")
	return ctx.JSON(response)
}
