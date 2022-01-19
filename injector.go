//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang-restapi-httprouter/app"
	"golang-restapi-httprouter/controller"
	"golang-restapi-httprouter/middleware"
	"golang-restapi-httprouter/repository"
	"golang-restapi-httprouter/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImplementation)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
	)

func InitializeServer() *http.Server{
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler),new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
		)
	return nil
}
