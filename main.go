package main

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//validate := validator.New()
	//
	//db := app.NewDB()
	//
	//categoryRepository := repository.NewCategoryRepositoryImpl()
	//categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)

	server := InitializedServer()
	err := server.ListenAndServe()
	panic(err)

}
