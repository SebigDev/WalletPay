package main

import (
	"CrashCourse/GoApp/infrastructures/db"
	"CrashCourse/GoApp/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	storeResponse := db.MongoInit()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	routes.MapRoute(app, &storeResponse)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("An error has occurred while starting the server")
		return
	}
	log.Println("Server running on port", 3000)
}
