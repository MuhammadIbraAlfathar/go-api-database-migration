package main

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/app"
	"github.com/MuhammadIbraAlfathar/go-restful-api/controller"
	"github.com/MuhammadIbraAlfathar/go-restful-api/middleware"
	"github.com/MuhammadIbraAlfathar/go-restful-api/repository"
	"github.com/MuhammadIbraAlfathar/go-restful-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	validate := validator.New()

	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	panic(err)

}
