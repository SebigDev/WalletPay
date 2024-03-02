package handlers

import (
	"log"

	"github.com/sebigdev/walletpay/internal/utils"
	"github.com/sebigdev/walletpay/modules/dto"
	"github.com/sebigdev/walletpay/modules/responses"
	"github.com/sebigdev/walletpay/modules/services"

	"github.com/gofiber/fiber/v2"
)

type IWalletHandler interface {
	AddWallet(ctx *fiber.Ctx) error
	Deposit(ctx *fiber.Ctx) error
	Withdraw(ctx *fiber.Ctx) error
	CreateTransaction(ctx *fiber.Ctx) error
	GetTransactions(ctx *fiber.Ctx) error
}

type walletHandler struct {
	WalletService      services.IWalletService
	TransactionService services.ITransactionService
}

func NewWalletHandler(walletService services.IWalletService, txService services.ITransactionService) IWalletHandler {
	return &walletHandler{
		WalletService:      walletService,
		TransactionService: txService,
	}
}

// CreateWallet func for creates a new wallet.
// @Description Create a new wallet.
// @Summary create a new wallet
// @Tags Transaction
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
// @Tags Transaction
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
// @Tags Transaction
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

// Createtransaction func for creates a new transaction.
// @Description Create a new transaction.
// @Summary create a new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body dto.CreateTransaction true "Create transaction"
// @Success 200
// @Router /api/v1/wallet/transaction [post]
func (w walletHandler) CreateTransaction(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}

	createTransactionRequest := new(dto.CreateTransaction)
	if err := ctx.BodyParser(createTransactionRequest); err != nil {
		return ctx.Status(500).JSON(responses.CreateErrorResponse("Error parsing request"))
	}

	res, err := w.TransactionService.Submit(userId, *createTransactionRequest)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse(res)
	return ctx.JSON(response)
}

// GetTransactions func gets all transactions.
// @Description Get all transactions.
// @Summary get all transactions
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {array} responses.TransactionFullResponse
// @Router /api/v1/wallet/transactions [get]
func (w walletHandler) GetTransactions(ctx *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromToken(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(err.Error())
	}
	persons, err := w.TransactionService.GetTransactions(userId)

	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(responses.CreateErrorResponse(err.Error()))
	}

	response := responses.CreateResponse(persons)
	return ctx.JSON(response)
}
