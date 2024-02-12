package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/wallet/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IWalletHandler interface {
	AddWallet(ctx *fiber.Ctx) error
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = w.WalletService.AddWallet(userId)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Wallet created successfully"})
}
