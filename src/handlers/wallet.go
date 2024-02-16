package handlers

import (
	"CrashCourse/GoApp/internal/utils"
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/responses"
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

// CreateWallet func for creates a new wallet.
// @Description Create a new wallet.
// @Summary create a new wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body dto.CreateWalletRequest true "Create wallet"
// @Success 200
// @Router /api/v1/wallet/add [post]
func (w walletHandler) AddWallet(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	createWalletRequest := new(dto.CreateWalletRequest)
	if err := ctx.BodyParser(createWalletRequest); err != nil {
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err = w.WalletService.AddWallet(userId, createWalletRequest.Currency)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Wallet created successfully")
	return ctx.JSON(response)
}

// Deposit func for deposits money in a wallet.
// @Description Deposit in a wallet.
// @Summary deposit in a wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body dto.DepositRequest true "Deposit into a wallet"
// @Success 200
// @Router /api/v1/wallet/deposit [post]
func (w walletHandler) Deposit(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	depositRequest := new(dto.DepositRequest)
	if err := ctx.BodyParser(depositRequest); err != nil {
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err = w.WalletService.Deposit(userId, *depositRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Wallet credited successfully")
	return ctx.JSON(response)
}

// Withdraw func for withdraws money from a wallet.
// @Description Withdraw from a wallet.
// @Summary Withdraw from a wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body dto.WithdrawRequest true "Withdraws from a wallet"
// @Success 200
// @Router /api/v1/wallet/withdraw [post]
func (w walletHandler) Withdraw(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	withdrawRequest := new(dto.WithdrawRequest)
	if err := ctx.BodyParser(withdrawRequest); err != nil {
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	err = w.WalletService.Withdraw(userId, *withdrawRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse("Wallet debited successfully")
	return ctx.JSON(response)
}
