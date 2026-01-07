package repository

import model "github.com/captainkie/websync-api/internal/app/models"

type UsersRepository interface {
	Create(users model.Users)
	Update(userId int, users model.Users)
	Delete(userId int)
	FindAll() []model.Users
	FindById(userId int) (model.Users, error)
	FindByUsername(username string) (model.Users, error)
}
