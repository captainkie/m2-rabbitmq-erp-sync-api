package repository

import (
	"errors"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Create implements UsersRepository
func (u *UsersRepositoryImpl) Create(users model.Users) {
	result := u.Db.Create(&users)
	helpers.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(userId int, users model.Users) {
	var updateUsers = request.UpdateUsersRequest{}

	if users.Password != "" {
		updateUsers.Password = users.Password
	}

	updateUsers.Role = users.Role
	updateUsers.Status = users.Status

	result := u.Db.Model(&users).Where("id = ?", userId).Updates(updateUsers)
	helpers.ErrorPanic(result.Error)
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(userId int) {
	var users model.Users
	result := u.Db.Where("id = ?", userId).Delete(&users)
	helpers.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := u.Db.Find(&users)
	helpers.ErrorPanic(results.Error)
	return users
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(userId int) (model.Users, error) {
	var users model.Users
	result := u.Db.Find(&users, userId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("users is not found")
	}
}

// FindByUsername implements UsersRepository
func (u *UsersRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var users model.Users
	result := u.Db.Where("username = ? AND status = ?", username, 1).First(&users)

	if result.Error != nil {
		return users, errors.New("invalid username, password or status not active")
	}
	return users, nil
}
