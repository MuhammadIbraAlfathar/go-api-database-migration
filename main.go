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

	server := InitializedServer()
	err := server.ListenAndServe()
	panic(err)

}
