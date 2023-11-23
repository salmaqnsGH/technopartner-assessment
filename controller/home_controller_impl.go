package controller

import (
	"net/http"
	"technopartner/test/helper"
	"technopartner/test/model/web"
	"technopartner/test/service"

	"github.com/julienschmidt/httprouter"
)

type HomeControllerImpl struct {
	HomeService service.HomeService
}

func NewHomeController(homeService service.HomeService) HomeController {
	return &HomeControllerImpl{
		HomeService: homeService,
	}
}

func (controller *HomeControllerImpl) TotalSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	totalSaldo := controller.HomeService.TotalSaldoCount(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   totalSaldo,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HomeControllerImpl) TotalSpend(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	totalSpend := controller.HomeService.TotalSpendCount(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   totalSpend,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HomeControllerImpl) TotalIncome(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	totalIncome := controller.HomeService.TotalIncomeCount(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   totalIncome,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
