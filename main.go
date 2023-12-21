package main

import (
	"cc-auth-service/db"
	"cc-auth-service/handler"
	"cc-auth-service/middleware"
	"cc-auth-service/repo"
	"cc-auth-service/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {

	app := fiber.New()

	// Limiter(?)
	app.Use(limiter.New(
		limiter.Config{
			Max:        1,
			Expiration: 1 * time.Second,
		},
	))

	app.Use(middleware.Logger())
	// app.Use(middleware.ApiKey())

	db.InitDatabase()

	userRepo := repo.NewUserRepo(db.DB)

	// Protected Route with ApiKey
	routes.InitRoutes(app, *userRepo)

	// Test Protected Route with JWT and ApiKey
	app.Use(middleware.JWTProtected())
	app.Get("/api/v1/test-protected", handler.Dashboard)

	app.Listen(":3000")
}
