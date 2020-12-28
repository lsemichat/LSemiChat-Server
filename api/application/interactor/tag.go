package interactor

import (
	"app/api/domain/entity"
	"app/api/domain/service"

	"github.com/pkg/errors"
)

type TagInteractor interface {
	Create(tag, categoryID string) (*entity.Tag, error)
	GetAll() ([]*entity.Tag, error)
	GetByID(id string) (*entity.Tag, error)
	GetByCategoryID(id string) ([]*entity.Tag, error)
}

type tagInteractor struct {
	tagService      service.TagService
	categoryService service.CategoryService
}

func NewTagInteractor(ts service.TagService, cs service.CategoryService) TagInteractor {
	return &tagInteractor{
		tagService:      ts,
		categoryService: cs,
	}
}

func (ti *tagInteractor) Create(tag, categoryID string) (*entity.Tag, error) {
	category, err := ti.categoryService.GetByID(categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find category")
	}
	newTag, err := ti.tagService.New(tag, categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tag")
	}
	newTag.Category = category
	return newTag, nil
}

func (ti *tagInteractor) GetAll() ([]*entity.Tag, error) {
	tags, err := ti.tagService.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	result := make([]*entity.Tag, 0, len(tags))
	for _, tag := range tags {
		// OPTIMIZE: 確実に遅いw
		category, _ := ti.categoryService.GetByID(tag.Category.ID)
		tag.Category = category
		result = append(result, tag)
	}
	return result, nil

}

func (ti *tagInteractor) GetByID(id string) (*entity.Tag, error) {
	tag, err := ti.tagService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tag")
	}
	category, err := ti.categoryService.GetByID(tag.Category.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find category")
	}
	tag.Category = category
	return tag, nil
}

func (ti *tagInteractor) GetByCategoryID(id string) ([]*entity.Tag, error) {
	tags, err := ti.tagService.GetByCategoryID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	result := make([]*entity.Tag, 0, len(tags))
	for _, tag := range tags {
		// OPTIMIZE: 確実に遅いw
		category, _ := ti.categoryService.GetByID(tag.Category.ID)
		tag.Category = category
		result = append(result, tag)
	}
	return result, nil
}
