package response

import "app/api/domain/entity"

type CategoryResponse struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type CategoriesResponse struct {
	Categories []*CategoryResponse `json:"categories"`
}

func ConvertToCategoryResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:       category.ID,
		Category: category.Category,
	}
}

func ConvertToCategoriesResponse(categories []*entity.Category) *CategoriesResponse {
	res := make([]*CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, ConvertToCategoryResponse(category))
	}
	return &CategoriesResponse{
		Categories: res,
	}
}
