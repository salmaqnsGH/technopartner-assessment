package controller

import (
	"net/http"
	"technopartner/test/helper"
	"technopartner/test/model/web"
	"technopartner/test/service"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)

	token, err := helper.GenerateToken(userResponse.ID)
	helper.PanicIfError(err)

	newUser := helper.FormatUser(userResponse, token)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   newUser,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse := controller.UserService.Login(request.Context(), userLoginRequest)

	token, err := helper.GenerateToken(userResponse.ID)
	helper.PanicIfError(err)

	newUser := helper.FormatUser(userResponse, token)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   newUser,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
