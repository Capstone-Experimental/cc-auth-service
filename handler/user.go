package handler

import (
	"cc-auth-service/helper"
	"cc-auth-service/model"
	"cc-auth-service/repo"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	RepoUser  repo.UserRepo
	Validator *validator.Validate
}

func NewUserHandler(repoUser repo.UserRepo) *UserHandler {
	return &UserHandler{
		RepoUser:  repoUser,
		Validator: validator.New(),
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

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uh.RepoUser.GetByID(id)
	if err != nil {
		return helper.Response(c, 400, "Failed to get user", nil)
	}

	return helper.Response(c, 200, "Success to get user", user)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var update model.UpdateUserRequest

	if err := c.BodyParser(&update); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}

	user, err := uh.RepoUser.Update(id, &update)
	if err != nil {
		return helper.Response(c, 400, "Failed to update user", nil)
	}

	return helper.Response(c, 200, "Success to update user", user)
}

func (uh *UserHandler) ForgotPassword(c *fiber.Ctx) error {
	var forgot model.ForgotPasswordRequest

	if err := c.BodyParser(&forgot); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}

	user, err := uh.RepoUser.GetByEmail(forgot.Email)
	if err != nil {
		return helper.Response(c, 400, "Failed to get user", nil)
	}

	err = uh.RepoUser.ForgotPassword(user.Email)
	if err != nil {
		return helper.Response(c, 400, "Failed to send OTP", nil)
	}

	return helper.Response(c, 200, "Success to send OTP, Please check your email", nil)

}

func (uh *UserHandler) VerifyOTP(c *fiber.Ctx) error {
	var verify model.VerifyOTPRequest

	if err := c.BodyParser(&verify); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}

	access, err := uh.RepoUser.CheckOTP(verify.OTP)
	if err != nil {
		return helper.Response(c, 400, "Failed to verify OTP", nil)
	}

	if !access {
		return helper.Response(c, 400, "OTP is wrong", nil)
	}

	return helper.Response(c, 200, "Success to verify OTP", nil)
}

func (uh *UserHandler) ResetPassword(c *fiber.Ctx) error {
	var reset model.ResetPasswordRequest

	if err := c.BodyParser(&reset); err != nil {
		return helper.Response(c, 400, "Error Parsing Body", nil)
	}
	if err := uh.Validator.Struct(reset); err != nil {
		return helper.Response(c, 400, "Error valiidate", nil)
	}
	if reset.Password1 != reset.Password2 {
		return helper.Response(c, 400, "Password not match", nil)
	}

	password, err := helper.GeneratePassword(reset.Password1)
	if err != nil {
		return helper.Response(c, 400, "Failed to hash password", nil)
	}

	err = uh.RepoUser.ResetPasswordAndOTP(reset.Email, password)
	if err != nil {
		return helper.Response(c, 400, "Failed to reset password", nil)
	}

	return helper.Response(c, 200, "Success to reset password", nil)
}
