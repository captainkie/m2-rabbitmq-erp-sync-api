package controller

import (
	"fmt"
	"net/http"

	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

// Login		godoc
// @Summary		Login
// @Description	login with username and password
// @Param    Request body request.LoginRequest{} true "Login"
// @Produce  application/json
// @tags Authentication
// @Success 200 {object} response.LoginResponse{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /authentication/login [post]
func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helpers.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username, password or status not active",
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Register		godoc
// @Summary		Register
// @Description	Register to websync systems
// @Param    Request body request.CreateUsersRequest true "Register"
// @Produce  application/json
// @tags Authentication
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /authentication/register [post]
func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helpers.ErrorPanic(err)

	controller.authenticationService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
