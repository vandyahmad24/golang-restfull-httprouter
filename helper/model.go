package helper

import (
	"golang-restapi-httprouter/model/domain"
	"golang-restapi-httprouter/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse{
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponseAll(categories []domain.Category) []web.CategoryResponse{

	var categoriesResponse []web.CategoryResponse
	for _, category := range categories{
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}