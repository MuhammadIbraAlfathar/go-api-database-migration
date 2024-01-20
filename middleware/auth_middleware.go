package middleware

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/helper"
	"github.com/MuhammadIbraAlfathar/go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "secret" == request.Header.Get("X-API-Key") {
		//oke
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error (unauthorized)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
