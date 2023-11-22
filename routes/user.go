package routes

import (
	"cc-auth-service/handler"
	"cc-auth-service/repo"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, repoUser repo.UserRepo) {

	userHandler := handler.NewUserHandler(repoUser)

	userRoutes := app.Group("/api/v1")

	userRoutes.Post("/register", userHandler.Register)   // Register
	userRoutes.Post("/login", userHandler.Login)         // Login
	userRoutes.Get("/user/:id", userHandler.GetUserByID) // Get User
	userRoutes.Put("/user/:id", userHandler.UpdateUser)  // Update User

	userRoutes.Post("/forgot-password", userHandler.ForgotPassword) // Forgot Password, Check Email, Send OTP
	userRoutes.Post("/verify-otp", userHandler.VerifyOTP)           // Verify OTP
	userRoutes.Put("/reset-password", userHandler.ResetPassword)    // Reset Password

	userRoutes.Get("/users", userHandler.GetAllUser) // Get All User
}
