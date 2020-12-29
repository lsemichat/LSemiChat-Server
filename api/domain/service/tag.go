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
	GetByUserUUID(id string) ([]*entity.Tag, error)
	GetByThreadID(id string) ([]*entity.Tag, error)
	GetByTagAndCategoryID(tagValue, categoryID string) (*entity.Tag, error)
	AddToUser(tagID, userUUID string) error
	RemoveFromUser(tagID, userUUID string) error
	AddToThread(tagID, threadID string) error
	RemoveFromThread(tagID, threadID string) error
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

func (ts *tagService) GetByUserUUID(id string) ([]*entity.Tag, error) {
	tags, err := ts.tagRepository.FindByUserUUID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return tags, nil
}

func (ts *tagService) GetByThreadID(id string) ([]*entity.Tag, error) {
	tags, err := ts.tagRepository.FindByThreadID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}
	return tags, nil
}

func (ts *tagService) GetByTagAndCategoryID(tagValue, categoryID string) (*entity.Tag, error) {
	tag, err := ts.tagRepository.FindByTagAndCategoryID(tagValue, categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tag")
	}
	return tag, nil
}

func (ts *tagService) AddToUser(tagID, userUUID string) error {
	id, err := GenerateUUID()
	if err != nil {
		return errors.Wrap(err, "failed to generate uuid")
	}

	if err = ts.tagRepository.AddToUser(id, tagID, userUUID); err != nil {
		return errors.Wrap(err, "failed to add tag to user")
	}
	return nil
}

func (ts *tagService) RemoveFromUser(tagID, userUUID string) error {
	if err := ts.tagRepository.RemoveFromUser(tagID, userUUID); err != nil {
		return errors.Wrap(err, "failed to remove tag from user")
	}
	return nil
}

func (ts *tagService) AddToThread(tagID, threadID string) error {
	id, err := GenerateUUID()
	if err != nil {
		return errors.Wrap(err, "failed to generate uuid")
	}

	if err = ts.tagRepository.AddToThread(id, tagID, threadID); err != nil {
		return errors.Wrap(err, "failed to add tag to thread")
	}
	return nil
}

func (ts *tagService) RemoveFromThread(tagID, threadID string) error {
	if err := ts.tagRepository.RemoveFromThread(tagID, threadID); err != nil {
		return errors.Wrap(err, "failed to remove tag from thread")
	}
	return nil
}
