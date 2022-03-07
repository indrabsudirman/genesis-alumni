package main

import (
	"genesis-alumni/database"
	"genesis-alumni/migration"
	"genesis-alumni/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initial Database
	database.DatabaseInit()

	//Initial Migration
	migration.Migration()

	app := fiber.New()

	//Initial Route
	route.RouteInit(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Failed to listen gofiber server:", errListen)
		os.Exit(1)
	}
}
