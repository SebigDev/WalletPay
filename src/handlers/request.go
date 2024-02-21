package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/responses"
	"CrashCourse/GoApp/src/modules/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IPaymentRequestHandler interface {
	SendRequest(c *fiber.Ctx) error
	AcknowldgeRequest(c *fiber.Ctx) error
}

type paymentRequestHandler struct {
	paymentRequestService services.IPaymentRequestService
}

func NewPaymentRequestHandler(paymentReqService services.IPaymentRequestService) IPaymentRequestHandler {
	return &paymentRequestHandler{
		paymentRequestService: paymentReqService,
	}
}

func (pr *paymentRequestHandler) SendRequest(c *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(err.Error())
	}

	createPayRequest := new(dto.CreatePayRequest)
	if err := c.BodyParser(createPayRequest); err != nil {
		return c.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}
	resp, err := pr.paymentRequestService.Request(userId, *createPayRequest)
	if err != nil {
		return err
	}
	return c.JSON(responses.CreateResponse(resp))
}

func (pr *paymentRequestHandler) AcknowldgeRequest(c *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(err.Error())
	}

	ackRequest := new(dto.AckRequest)
	if err := c.BodyParser(ackRequest); err != nil {
		return c.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}
	err = pr.paymentRequestService.AcknowledgeRequest(userId, ackRequest.RequestId)
	if err != nil {
		return err
	}
	return c.JSON(responses.CreateResponse("Acknowldgement was successfully"))
}
