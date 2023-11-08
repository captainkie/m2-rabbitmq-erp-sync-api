package service

import (
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/utils"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements UsersService
func (u *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helpers.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	u.UsersRepository.Create(newUser)
}

// Update implements UsersService
func (u *UsersServiceImpl) Update(userId int, users request.UpdateUsersRequest) {
	userData, err := u.UsersRepository.FindById(userId)
	helpers.ErrorPanic(err)

	if users.Password != "" {
		hashedPassword, err := utils.HashPassword(users.Password)
		helpers.ErrorPanic(err)
		userData.Password = hashedPassword
	}

	userData.Role = users.Role
	userData.Status = users.Status
	u.UsersRepository.Update(userId, userData)
}

// Delete implements UsersService
func (u *UsersServiceImpl) Delete(userId int) {
	u.UsersRepository.Delete(userId)
}

// FindAll implements UsersService
func (u *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := u.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			ID:       value.ID,
			Username: value.Username,
			Email:    value.Email,
			Role:     value.Role,
			Status:   value.Status,
		}

		users = append(users, user)
	}

	return users
}

// FindById implements UsersService
func (u *UsersServiceImpl) FindById(userId int) response.UsersResponse {
	userData, err := u.UsersRepository.FindById(userId)
	helpers.ErrorPanic(err)

	userResponse := response.UsersResponse{
		ID:       userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
		Role:     userData.Role,
		Status:   userData.Status,
	}

	return userResponse
}
