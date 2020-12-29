package interactor

import (
	"app/api/domain/entity"
	"app/api/domain/service"
)

func AddCategoryToTag(tags []*entity.Tag, categoryService service.CategoryService) []*entity.Tag {
	result := make([]*entity.Tag, 0, len(tags))
	for _, tag := range tags {
		category, _ := categoryService.GetByID(tag.Category.ID)
		tag.Category = category
		result = append(result, tag)
	}
	return result
}
