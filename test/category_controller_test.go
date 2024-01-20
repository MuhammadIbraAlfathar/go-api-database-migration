package test

import (
	"database/sql"
	"encoding/json"
	"github.com/MuhammadIbraAlfathar/go-restful-api/app"
	"github.com/MuhammadIbraAlfathar/go-restful-api/controller"
	"github.com/MuhammadIbraAlfathar/go-restful-api/helper"
	"github.com/MuhammadIbraAlfathar/go-restful-api/middleware"
	"github.com/MuhammadIbraAlfathar/go-restful-api/repository"
	"github.com/MuhammadIbraAlfathar/go-restful-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {

	db := setupTestDB()

	truncateCategory(db)

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "ibrasss"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateCategoryFail(t *testing.T) {

}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFail(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFail(t *testing.T) {

}

func TestFindByIdCategorySuccess(t *testing.T) {

}

func TestFindByIdCategoryFail(t *testing.T) {

}

func TestListCategorySuccess(t *testing.T) {

}

func TestListCategoryFail(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
