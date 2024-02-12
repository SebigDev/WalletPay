package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IWalletHandler interface {
	AddWallet(ctx *fiber.Ctx) error
	Deposit(ctx *fiber.Ctx) error
	Withdraw(ctx *fiber.Ctx) error
}

type walletHandler struct {
	WalletService services.IWalletService
}

func NewWalletHandler(walletService services.IWalletService) IWalletHandler {
	return &walletHandler{
		WalletService: walletService,
	}
}

func (w walletHandler) AddWallet(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	err = w.WalletService.AddWallet(userId)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Wallet created successfully"})
}

func (w walletHandler) Deposit(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	depositRequest := new(dto.DepositRequest)
	if err := ctx.BodyParser(depositRequest); err != nil {
		return ctx.Status(400).JSON("Error parsing request")
	}

	err = w.WalletService.Deposit(userId, *depositRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Wallet credited successfully"})
}

func (w walletHandler) Withdraw(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	withdrawRequest := new(dto.WithdrawRequest)
	if err := ctx.BodyParser(withdrawRequest); err != nil {
		return ctx.Status(400).JSON("Error parsing request")
	}

	err = w.WalletService.Withdraw(userId, *withdrawRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Wallet debited successfully"})
}
