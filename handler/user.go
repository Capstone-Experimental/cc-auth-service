package handler

import (
	"cc-auth-service/helper"
	"cc-auth-service/model"
	"cc-auth-service/repo"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	RepoUser repo.UserRepo
}

func NewUserHandler(repoUser repo.UserRepo) *UserHandler {
	return &UserHandler{
		RepoUser: repoUser,
	}
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {

	var register model.RegisterRequest

	if err := c.BodyParser(&register); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}

	existingUser, _ := uh.RepoUser.GetByEmail(register.Email)

	if existingUser.ID != "" {
		return helper.Response(c, 400, "Email already exists", nil)
	}

	password, err := helper.GeneratePassword(register.Password)
	if err != nil {
		return helper.Response(c, 400, "Error Hash Password", nil)
	}

	u := model.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: password,
	}
	_, err = uh.RepoUser.Create(&u)

	if err != nil {
		return helper.Response(c, 400, "Failed to Create Account", nil)
	}

	return helper.Response(c, 200, "Success to register", nil)
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {

	var login model.LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}

	user, err := uh.RepoUser.GetByEmail(login.Email)
	if err != nil {
		return helper.Response(c, 401, "Email or password wrong", nil)
	}

	if !helper.VerifyPassword(login.Password, user.Password) {
		return helper.Response(c, 401, "Email or password wrong", nil)
	}

	token, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return helper.Response(c, 400, "Failed generate token", nil)
	}
	return helper.Response(c, 200, "Success to login", fiber.Map{"token": token})
}

func (uh *UserHandler) GetAllUser(c *fiber.Ctx) error {
	var users []model.User

	users, err := uh.RepoUser.GetAll()
	if err != nil {
		return helper.Response(c, 400, "Failed to get all user", nil)
	}

	return helper.Response(c, 200, "Success to get all user", users)
}
