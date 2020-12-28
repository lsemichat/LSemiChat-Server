package service

import (
	"app/api/domain/entity"
	"app/api/domain/repository"

	"github.com/pkg/errors"
)

type TagService interface {
	New(tag, categoryID string) (*entity.Tag, error)
	GetAll() ([]*entity.Tag, error)
	GetByID(id string) (*entity.Tag, error)
	GetByCategoryID(id string) ([]*entity.Tag, error)
}

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagService(tr repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tr,
	}
}

func (ts *tagService) New(tag, categoryID string) (*entity.Tag, error) {
	id, err := GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate uuid")
	}
	newTag := &entity.Tag{
		ID:  id,
		Tag: tag,
		Category: &entity.Category{
			ID: categoryID,
		},
	}
	err = ts.tagRepository.Create(newTag)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create")
	}
	return newTag, nil

}

func (ts *tagService) GetAll() ([]*entity.Tag, error) {
	tags, err := ts.tagRepository.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return tags, nil
}

func (ts *tagService) GetByID(id string) (*entity.Tag, error) {
	tag, err := ts.tagRepository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return tag, nil
}

func (ts *tagService) GetByCategoryID(id string) ([]*entity.Tag, error) {
	tags, err := ts.tagRepository.FindByCategoryID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return tags, nil
}
