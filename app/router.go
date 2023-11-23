package app

import (
	"technopartner/test/controller"
	"technopartner/test/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, userController controller.UserController, transactionController controller.TransactionController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindByID)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.GET("/api/v1/transactions", transactionController.FindAll)
	router.GET("/api/v1/transactions/:transactionId", transactionController.FindByID)
	router.POST("/api/v1/transactions", transactionController.Create)
	router.PUT("/api/v1/transactions/:transactionId", transactionController.Update)
	router.DELETE("/api/v1/transactions/:transactionId", transactionController.Delete)

	router.POST("/api/v1/users/register", userController.Register)
	router.POST("/api/v1/users/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
