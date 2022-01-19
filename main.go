package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server{
	return &http.Server{
		Handler: authMiddleware,
		Addr: "localhost:3000",
	}
}

func main()  {
	//validate := validator.New()
	//db := app.NewDB()
	//categoryRepository := repository.NewCategoryRepository()
	//NewCategoryService := service.NewCategoryService(categoryRepository, db, validate)
	//NewCategoryController := controller.NewCategoryController(NewCategoryService)
	//router := app.NewRouter(NewCategoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)
	//
	//server := NewServer(authMiddleware)
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicError(err)
}