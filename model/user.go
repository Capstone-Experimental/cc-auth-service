package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primary_key;" json:"id"`
	Name     string `gorm:"size:255;" json:"name"`
	Email    string `gorm:"unique_index" json:"email"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
