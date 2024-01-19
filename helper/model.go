package helper

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/model/domain"
	"github.com/MuhammadIbraAlfathar/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}

	return categoriesResponse
}
