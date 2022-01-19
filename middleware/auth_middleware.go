package middleware

import (
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY"){
	//	oke
		middleware.Handler.ServeHTTP(writer, request)
	}else{
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		response := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "Key not found",
		}
		helper.WriteToResponseBody(writer,response)
	}
}
