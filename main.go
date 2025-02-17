package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mithushikaS/Block/routes"
	"github.com/mithushikaS/Block/database"
)

func main() {
	database.Connect()
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port:=os.Getenv("PORT")
	app := fiber.New(fiber.Config{})
	routes.Setup(app)
	app.Listen(":"+port)

}