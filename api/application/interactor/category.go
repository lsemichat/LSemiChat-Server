package interactor

import (
	"app/api/domain/entity"
	"app/api/domain/service"

	"github.com/pkg/errors"
)

type CategoryInteractor interface {
	Create(category string) (*entity.Category, error)
	GetAll() ([]*entity.Category, error)
	GetByID(id string) (*entity.Category, error)
	Update(id, category string) (*entity.Category, error)
	Delete(id string) error
}

type categoryInteractor struct {
	categoryService service.CategoryService
}

func NewCategoryInteractor(cs service.CategoryService) CategoryInteractor {
	return &categoryInteractor{
		categoryService: cs,
	}
}

func (ci *categoryInteractor) Create(category string) (*entity.Category, error) {
	data, err := ci.categoryService.Create(category)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create")
	}
	return data, nil
}

func (ci *categoryInteractor) GetAll() ([]*entity.Category, error) {
	categories, err := ci.categoryService.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return categories, nil
}

func (ci *categoryInteractor) GetByID(id string) (*entity.Category, error) {
	category, err := ci.categoryService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return category, nil
}

func (ci *categoryInteractor) Update(id, category string) (*entity.Category, error) {
	data, err := ci.categoryService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	newCategory, err := ci.categoryService.Update(data, category)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update")
	}
	return newCategory, nil
}

func (ci *categoryInteractor) Delete(id string) error {
	if err := ci.categoryService.Delete(id); err != nil {
		return errors.Wrap(err, "failed to delete")
	}
	return nil
}
