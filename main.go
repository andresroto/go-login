package main

import (
	"log"

	"github.com/andresroto/go-login/database"
	"github.com/andresroto/go-login/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	database.Connect()

	app := fiber.New()

	// Setup CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Import routes
	routes.Setup(app)

	app.Listen(":3000")
}
