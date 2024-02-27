package main

import (
	"log"

	_ "github.com/sebigdev/GoApp/docs"
	"github.com/sebigdev/GoApp/infrastructures/db"
	"github.com/sebigdev/GoApp/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title GoApp Wallet API
// @version 2.0
// @description This is a GoApp project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /
// @schemes http https
func main() {
	storeResponse := db.MongoInit()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	routes.MapCommon(app, &storeResponse)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))

	err := app.Listen("127.0.0.1:3001")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("An error has occurred while starting the server")
		return
	}
	log.Println("Server running on port", 3000)
}
