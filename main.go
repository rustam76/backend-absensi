package main

import (
	"backend-absensi/config"
	"backend-absensi/routes"
	"fmt"
	"log"
	"os"

	_ "backend-absensi/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Absensi API
// @version 1.0
// @description REST API untuk sistem absensi
// @host localhost:3030
// @BasePath /api
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.ConnectDB()

	// Setup Fiber
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://frontend-absensi-omega.vercel.app",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3030"
	}
	log.Printf("🚀 Server running at http://localhost:%s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
