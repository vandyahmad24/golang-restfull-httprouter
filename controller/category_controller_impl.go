package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/model/web"
	"golang-restapi-httprouter/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl{
		return &CategoryControllerImpl{
			CategoryService: categoryService,
		}
}


func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(),categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data:categoryResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)


}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryId :=params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)
	categoryCreateRequest.Id=id

	categoryResponse := controller.CategoryService.Update(request.Context(),categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data:categoryResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId :=params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)


	controller.CategoryService.Delete(request.Context(),id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId :=params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)


	result := controller.CategoryService.FindById(request.Context(),id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: result,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	result := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: result,
	}
	helper.WriteToResponseBody(writer,webResponse)
}
