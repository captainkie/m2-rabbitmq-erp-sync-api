package service

import (
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(userId int, users request.UpdateUsersRequest)
	Delete(userId int)
	FindById(userId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
