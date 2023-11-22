package model

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"email"`
	Password *string `json:"password" validate:"min=6,max=20"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyOTPRequest struct {
	OTP string `json:"otp" validate:"required,min=6,max=6"`
}

type ResetPasswordRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password1 string `json:"password1" validate:"required,min=6,max=20"`
	Password2 string `json:"password2" validate:"required,min=6,max=20"`
}
