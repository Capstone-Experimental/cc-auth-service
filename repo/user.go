package repo

import (
	"cc-auth-service/helper"
	"cc-auth-service/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

// GetAll returns all users
func (ur *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	result := ur.Db.Find(&users)

	// hide password
	for i := range users {
		users[i].Password = ""
		users[i].OTP = ""
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetByID returns a user by id
func (ur *UserRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	result := ur.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	user.OTP = ""
	user.Password = ""

	return &user, nil
}

// GetByEmail returns a user by email
func (ur *UserRepo) GetByEmail(email string) (model.User, error) {
	var user model.User
	result := ur.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

// Create creates a new user
func (ur *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := ur.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates a user
func (ur *UserRepo) Update(userID string, updateRequest *model.UpdateUserRequest) (*model.User, error) {
	var user model.User
	result := ur.Db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return &model.User{}, result.Error
	}

	if updateRequest.Name != nil {
		user.Name = *updateRequest.Name
	}
	if updateRequest.Email != nil {
		user.Email = *updateRequest.Email
	}
	if updateRequest.Password != nil {
		hash, err := helper.GeneratePassword(*updateRequest.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hash
	}

	if err := ur.Db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete deletes a user
func (ur *UserRepo) Delete(user *model.User) (*model.User, error) {
	if err := ur.Db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ForgotPassword in case user forgot password
func (ur *UserRepo) ForgotPassword(email string) error {

	otp, err := helper.SendOTP(email)
	if err != nil {
		return err
	}

	user, err := ur.GetByEmail(email)
	if err != nil {
		return err
	}
	// save otp to db
	if err := ur.Db.Model(&user).Update("otp", otp).Error; err != nil {
		return err
	}

	return nil
}

// CheckOTP checks if otp is valid
func (ur *UserRepo) CheckOTP(otpInput string) (bool, error) {
	var user model.User
	result := ur.Db.Where("otp = ?", otpInput).First(&user)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

// ResetPasswordAndOTP resets password and otp
func (ur *UserRepo) ResetPasswordAndOTP(email, password string) error {
	user, err := ur.GetByEmail(email)
	if err != nil {
		return err
	}

	if user.OTP == "" {
		return errors.New("OTP tidak ada")
	}

	user.OTP = ""

	result := ur.Db.Model(&user).Updates(map[string]interface{}{"password": password, "otp": ""})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
