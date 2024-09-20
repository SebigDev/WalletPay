package routes

import (
	"time"

	"github.com/sebigdev/walletpay/handlers"
	"github.com/sebigdev/walletpay/infrastructures/db"
	"github.com/sebigdev/walletpay/internal/middlewares"
	"github.com/sebigdev/walletpay/modules/repositories"
	"github.com/sebigdev/walletpay/modules/services"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var authMiddleware *middlewares.AuthMiddleware
var oContext *middlewares.OperationContext

func Init(app *fiber.App, props *db.DbProps) {

	//SET OPERATION CONTEXT
	oContext.SetXRequestIDContext(app)
	oContext.LivenessAndHealthCheck(app)

	//INITIATE EVENT BUS
	eventBus := services.NewEventBus()

	//REPOSITORIES
	userRepository := repositories.NewUserRepository(props.UserCollection, props.Context)
	trxRepository := repositories.NewTransactionRepository(props.TransactionCollection, props.Context)
	payReqRepository := repositories.NewPaymentRequestRepository(props.PaymentCollection, props.Context)

	//SERVICES
	userService := services.NewUserService(userRepository, eventBus)
	walletService := services.NewWalletService(userRepository)
	trxService := services.NewTransactionService(trxRepository, userRepository, eventBus)
	payReqService := services.NewPaymentRequestService(payReqRepository, userRepository)

	//HANDLERS
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService)
	walletHander := handlers.NewWalletHandler(walletService, trxService)
	payReqHandler := handlers.NewPaymentRequestHandler(payReqService)

	//person Routes

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Post("/users/onboard", userHandler.CreatePerson)
	v1.Post("/auth/login", authHandler.Authenticate)

	//swagger routes
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "WalletPay Metrics", Refresh: time.Hour}))

	//Authenticated routes
	v1.Get("/user", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetPersonById)
	v1.Get("/users", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetAllUsers)
	v1.Post("/user/change-password", authMiddleware.UserAuthMiddlewareHandler, userHandler.ChangePassword)
	v1.Post("/user/change-pin", authMiddleware.UserAuthMiddlewareHandler, userHandler.ChangePin)

	//WALLET
	v1.Post("/wallet/add", authMiddleware.UserAuthMiddlewareHandler, walletHander.AddWallet)
	v1.Post("/wallet/deposit", authMiddleware.UserAuthMiddlewareHandler, walletHander.Deposit)
	v1.Post("/wallet/withdraw", authMiddleware.UserAuthMiddlewareHandler, walletHander.Withdraw)
	v1.Post("/wallet/transaction", authMiddleware.UserAuthMiddlewareHandler, walletHander.CreateTransaction)
	v1.Get("/wallet/transactions", authMiddleware.UserAuthMiddlewareHandler, walletHander.GetTransactions)

	//REQUEST
	v1.Post("/request", authMiddleware.UserAuthMiddlewareHandler, payReqHandler.SendRequest)
	v1.Post("/request/acknowledge", authMiddleware.UserAuthMiddlewareHandler, payReqHandler.AcknowldgeRequest)

	//SRTVICES FOR EVENT BUS
	services.New(eventBus,
		services.WithWalletService(walletService),
		services.WithUserService(userService),
		services.WithTransactionService(trxService),
	)
}
