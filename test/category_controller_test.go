package test

import (
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang-restapi-httprouter/app"
	"golang-restapi-httprouter/controller"
	"golang-restapi-httprouter/middleware"
	"golang-restapi-httprouter/repository"
	"golang-restapi-httprouter/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUpRouter() http.Handler  {
	db := app.SetUpTestDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	NewCategoryService := service.NewCategoryService(categoryRepository, db, validate)
	NewCategoryController := controller.NewCategoryController(NewCategoryService)
	router := app.NewRouter(NewCategoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setUpRouter()
	requestBody := strings.NewReader(`{"name":"ahmad"}`)
	request := httptest.NewRequest("POST","http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("X-API-KEY","RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder,request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestCreateCategoryFailed(t *testing.T) {

}