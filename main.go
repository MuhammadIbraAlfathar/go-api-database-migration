package main

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/app"
	"github.com/MuhammadIbraAlfathar/go-restful-api/controller"
	"github.com/MuhammadIbraAlfathar/go-restful-api/repository"
	"github.com/MuhammadIbraAlfathar/go-restful-api/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()

	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
