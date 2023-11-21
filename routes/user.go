package routes

import (
	"cc-auth-service/handler"
	"cc-auth-service/repo"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, repoUser repo.UserRepo) {

	userHandler := handler.NewUserHandler(repoUser)

	userRoutes := app.Group("/api/v1")

	userRoutes.Post("/register", userHandler.Register)
	userRoutes.Post("/login", userHandler.Login)

	userRoutes.Get("/users", userHandler.GetAllUser)
}
