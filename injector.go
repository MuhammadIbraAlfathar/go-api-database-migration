//go:build wireinject
// +build wireinject

package main

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/app"
	"github.com/MuhammadIbraAlfathar/go-restful-api/controller"
	"github.com/MuhammadIbraAlfathar/go-restful-api/middleware"
	"github.com/MuhammadIbraAlfathar/go-restful-api/repository"
	"github.com/MuhammadIbraAlfathar/go-restful-api/service"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	controller.NewCategoryController,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer)

	return nil
}
