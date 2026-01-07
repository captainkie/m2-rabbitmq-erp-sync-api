package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UsersService
}

func NewUserController(service service.UsersService) *UserController {
	return &UserController{userService: service}
}

// Create		godoc
// @Summary		Create User
// @Description	create new user
// @Param    Request body request.CreateUsersRequest true "CreateUser"
// @Produce  application/json
// @tags Users
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users [post]
// @Security BearerAuth
func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helpers.ErrorPanic(err)

	controller.userService.Create(createUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully create new user!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update		godoc
// @Summary		Update User
// @Description	update user
// @Param    id path int true "id"
// @Param    Request body request.UpdateUsersRequest true "UpdateUser"
// @Produce  application/json
// @tags Users
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/{id} [patch]
// @Security BearerAuth
func (controller *UserController) Update(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, errId := strconv.Atoi(userId)
	if errId != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", errId),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	updateUserRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	controller.userService.Update(id, updateUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully update user!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete		godoc
// @Summary		Delete User
// @Description	delete user
// @Param    id path int true "id"
// @Produce  application/json
// @tags Users
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/{id} [delete]
// @Security BearerAuth
func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	controller.userService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully delete user!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll		godoc
// @Summary		Find All User
// @Description	find all user
// @Produce  application/json
// @tags Users
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users [get]
// @Security BearerAuth
func (controller *UserController) FindAll(ctx *gin.Context) {
	users := controller.userService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById		godoc
// @Summary		Find By Id User
// @Description	find by id user
// @Param    id path int true "id"
// @Produce  application/json
// @tags Users
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/{id} [get]
// @Security BearerAuth
func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	user := controller.userService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch user data!",
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
