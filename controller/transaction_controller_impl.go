package controller

import (
	"net/http"
	"strconv"
	"technopartner/test/helper"
	"technopartner/test/model/web"
	"technopartner/test/service"

	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionCreateRequest)

	transactionResponse := controller.TransactionService.Create(request.Context(), transactionCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionUpdateRequest := web.TransactionUpdateRequest{}
	helper.ReadFromRequestBody(request, &transactionUpdateRequest)

	transactionID := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionID)
	helper.PanicIfError(err)

	transactionUpdateRequest.ID = id

	transactionResponse := controller.TransactionService.Update(request.Context(), transactionUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionID := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionID)
	helper.PanicIfError(err)

	controller.TransactionService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionID := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionID)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.FindByID(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionResponses := controller.TransactionService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transactionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
