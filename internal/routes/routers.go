package routes

import (
	"CrashCourse/GoApp/infrastructures/db"
	"CrashCourse/GoApp/internal/middlewares"
	"CrashCourse/GoApp/src/handlers"
	"CrashCourse/GoApp/src/modules/repositories"
	"CrashCourse/GoApp/src/modules/services"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

var authMiddleware *middlewares.AuthMiddleware

func MapCommon(app *fiber.App, store *db.MongoResponse) {

	//INITIATE EVENT BUS
	eventBus := services.NewEventBus()

	//COLLECTION
	userCollection := store.ClientR.Database("goapp").Collection("users")
	trxCollection := store.ClientR.Database("goapp").Collection("transactions")
	payReqCollection := store.ClientR.Database("goapp").Collection("payments")

	//REPOSITORIES
	userRepository := repositories.NewUserRepository(userCollection, store.CtxR)
	trxRepository := repositories.NewTransactionRepository(trxCollection, store.CtxR)
	payReqRepository := repositories.NewPaymentRequestRepository(payReqCollection, store.CtxR)

	//SERVICES
	userService := services.NewUserService(userRepository, eventBus)
	walletService := services.NewWalletService(userRepository)
	trxService := services.NewTransactionService(trxRepository, userRepository)
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
	//Authenticated routes
	v1.Get("/user", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetPersonById)
	v1.Get("/users", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetAllUsers)
	v1.Post("/wallet/add", authMiddleware.UserAuthMiddlewareHandler, walletHander.AddWallet)
	v1.Post("/wallet/deposit", authMiddleware.UserAuthMiddlewareHandler, walletHander.Deposit)
	v1.Post("/wallet/withdraw", authMiddleware.UserAuthMiddlewareHandler, walletHander.Withdraw)
	v1.Post("/wallet/transaction", authMiddleware.UserAuthMiddlewareHandler, walletHander.CreateTransaction)
	v1.Get("/wallet/transactions", authMiddleware.UserAuthMiddlewareHandler, walletHander.GetTransactions)
	v1.Post("/request", authMiddleware.UserAuthMiddlewareHandler, payReqHandler.SendRequest)
	v1.Post("/request/acknowledge", authMiddleware.UserAuthMiddlewareHandler, payReqHandler.AcknowldgeRequest)

	//SRTVICES FOREVENT BUS
	services.New(eventBus,
		services.WithWalletService(walletService),
		services.WithUserService(userService),
	)
}
