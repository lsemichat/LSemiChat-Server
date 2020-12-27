package service

import (
	"app/api/application/repository"
	"app/api/domain/entity"

	"github.com/pkg/errors"
)

type CategoryService interface {
	Create(category string) (*entity.Category, error)
	GetAll() ([]*entity.Category, error)
	GetByID(id string) (*entity.Category, error)
	Update(category *entity.Category, newCategory string) (*entity.Category, error)
	Delete(id string) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: cr,
	}
}

func (cs *categoryService) Create(category string) (*entity.Category, error) {
	uuid, err := GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate id")
	}
	data := &entity.Category{
		ID:       uuid,
		Category: category,
	}
	if err = cs.categoryRepository.Create(data); err != nil {
		return nil, errors.Wrap(err, "failed to create category")
	}
	return data, nil
}

func (cs *categoryService) GetAll() ([]*entity.Category, error) {
	categories, err := cs.categoryRepository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all")
	}
	return categories, nil
}

func (cs *categoryService) GetByID(id string) (*entity.Category, error) {
	category, err := cs.categoryRepository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return category, nil
}

func (cs *categoryService) Update(category *entity.Category, newCategory string) (*entity.Category, error) {
	category.Category = newCategory
	if err := cs.categoryRepository.Update(category); err != nil {
		return nil, errors.Wrap(err, "failed to update")
	}
	return category, nil
}

func (cs *categoryService) Delete(id string) error {
	if err := cs.categoryRepository.Delete(id); err != nil {
		return errors.Wrap(err, "failed to delete")
	}
	return nil
}
