package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
}
