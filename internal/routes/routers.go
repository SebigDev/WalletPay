package routes

import (
	"CrashCourse/GoApp/infrastructures/db"
	"CrashCourse/GoApp/internal/middlewares"
	"CrashCourse/GoApp/src/handlers"
	"CrashCourse/GoApp/src/modules/user/repositories"
	"CrashCourse/GoApp/src/modules/user/services"

	"github.com/gofiber/fiber/v2"
)

var authMiddleware *middlewares.AuthMiddleware

func MapRoute(app *fiber.App, store *db.MongoResponse) {

	//USERS
	userCollection := store.ClientR.Database("goapp").Collection("users")
	userRepository := repositories.NewUserRepository(userCollection, store.CtxR)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	authHandler := handlers.NewAuthHandler(userService)

	//person Routes

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Post("/users/onboard", userHandler.CreatePerson)
	v1.Post("/auth/login", authHandler.Authenticate)

	//Authenticated routes
	v1.Get("/user", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetPersonById)
	v1.Get("/users", authMiddleware.UserAuthMiddlewareHandler, userHandler.GetAllUsers)
}
