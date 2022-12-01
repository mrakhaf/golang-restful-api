package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrakhaf/golang-restful-api/config"
	"github.com/mrakhaf/golang-restful-api/controller"
	"github.com/mrakhaf/golang-restful-api/helper"
	"github.com/mrakhaf/golang-restful-api/middleware"
	"github.com/mrakhaf/golang-restful-api/repository"
	"github.com/mrakhaf/golang-restful-api/service"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	serviceCategory := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(serviceCategory)

	router := config.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
