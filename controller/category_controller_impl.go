package controller

import (
	"net/http"
	"strconv"
	"technopartner/test/helper"
	"technopartner/test/model/web"
	"technopartner/test/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryUpdateRequest.ID = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
