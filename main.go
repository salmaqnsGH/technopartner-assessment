package main

import (
	"net/http"
	"technopartner/test/app"
	"technopartner/test/controller"
	"technopartner/test/db"
	"technopartner/test/helper"
	"technopartner/test/middleware"
	"technopartner/test/repository"
	"technopartner/test/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := db.NewDB()

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)

	router := app.NewRouter(categoryController, userController, transactionController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
