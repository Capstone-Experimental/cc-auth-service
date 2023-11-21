package repo

import (
	"cc-auth-service/model"

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

func (ur *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	result := ur.Db.Find(&users)

	// hiden password
	for i := range users {
		users[i].Password = ""
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *UserRepo) GetByID(id string) (model.User, error) {
	var user model.User
	result := ur.Db.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (ur *UserRepo) GetByEmail(email string) (model.User, error) {
	var user model.User
	result := ur.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (ur *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := ur.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) Update(user *model.User) (*model.User, error) {
	if err := ur.Db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) Delete(user *model.User) (*model.User, error) {
	if err := ur.Db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
