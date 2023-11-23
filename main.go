package main

import (
	"net/http"
	"technopartner/test/app"
	"technopartner/test/controller"
	"technopartner/test/db"
	"technopartner/test/helper"
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

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
