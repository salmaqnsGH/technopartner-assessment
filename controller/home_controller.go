package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController interface {
	TotalSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	TotalSpend(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	TotalIncome(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
