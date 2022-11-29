package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/mrakhaf/golang-restful-api/config"
	"github.com/mrakhaf/golang-restful-api/controller"
	"github.com/mrakhaf/golang-restful-api/exeption"
	"github.com/mrakhaf/golang-restful-api/helper"
	"github.com/mrakhaf/golang-restful-api/repository"
	"github.com/mrakhaf/golang-restful-api/service"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	serviceCategory := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(serviceCategory)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
