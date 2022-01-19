package exceptions

import (
	"github.com/go-playground/validator/v10"
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{})  {
	if notFoundError(writer, request, err){
		return
	}
	if validationError(writer, request, err){

		return
	}

	internalServerError(writer, request, err)

}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{})bool{
	exception, ok := err.(validator.ValidationErrors)
	if ok{
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusBadRequest)
		response := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "bad request",
			Data: exception.Error(),
		}
		helper.WriteToResponseBody(writer,response)
		return true
	}else{
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool  {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusNotFound)
		response := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "not found",
			Data: exception.Error,
		}
		helper.WriteToResponseBody(writer,response)
		return true
	}else{
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{})  {
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		response := web.WebResponse{
			Code: http.StatusInternalServerError,
			Status: "internal server error",
			Data: err,
		}
		helper.WriteToResponseBody(writer,response)
}
