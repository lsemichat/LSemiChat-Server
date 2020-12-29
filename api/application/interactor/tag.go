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
	GetByUserUUID(id string) ([]*entity.Tag, error)
	GetByThreadID(id string) ([]*entity.Tag, error)
	AddToUser(tagValue, categoryID, userID string) error
	RemoveFromUser(tagID, userID string) error
	AddToThread(tagValue, categoryID, threadID string) error
	RemoveFromThread(tagID, threadID string) error
}

type tagInteractor struct {
	tagService      service.TagService
	categoryService service.CategoryService
	userService     service.UserService
	threadService   service.ThreadService
}

func NewTagInteractor(ts service.TagService, cs service.CategoryService, us service.UserService, ths service.ThreadService) TagInteractor {
	return &tagInteractor{
		tagService:      ts,
		categoryService: cs,
		userService:     us,
		threadService:   ths,
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

func (ti *tagInteractor) GetByUserUUID(id string) ([]*entity.Tag, error) {
	tags, err := ti.tagService.GetByUserUUID(id)
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

func (ti *tagInteractor) GetByThreadID(id string) ([]*entity.Tag, error) {
	tags, err := ti.tagService.GetByThreadID(id)
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

func (ti *tagInteractor) AddToUser(tagValue, categoryID, userID string) error {
	user, err := ti.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	tag, err := ti.tagService.GetByTagAndCategoryID(tagValue, categoryID)
	if err != nil {
		tag, err = ti.tagService.New(tagValue, categoryID)
		if err != nil {
			return errors.Wrap(err, "failed to create tag")
		}
	}

	if err = ti.tagService.AddToUser(tag.ID, user.ID); err != nil {
		return errors.Wrap(err, "failed to add tag to user")
	}
	return nil
}

func (ti *tagInteractor) RemoveFromUser(tagID, userID string) error {
	user, err := ti.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}
	if err = ti.tagService.RemoveFromUser(tagID, user.ID); err != nil {
		return errors.Wrap(err, "failed to remove tag from user")
	}
	return nil
}

func (ti *tagInteractor) AddToThread(tagValue, categoryID, threadID string) error {
	_, err := ti.threadService.GetByID(threadID)
	if err != nil {
		return errors.Wrap(err, "failed to find thread")
	}

	tag, err := ti.tagService.GetByTagAndCategoryID(tagValue, categoryID)
	if err != nil {
		tag, err = ti.tagService.New(tagValue, categoryID)
		if err != nil {
			return errors.Wrap(err, "failed to create tag")
		}
	}

	if err = ti.tagService.AddToThread(tag.ID, threadID); err != nil {
		return errors.Wrap(err, "failed to add tag to user")
	}
	return nil
}

func (ti *tagInteractor) RemoveFromThread(tagID, threadID string) error {
	if err := ti.tagService.RemoveFromThread(tagID, threadID); err != nil {
		return errors.Wrap(err, "failed to remove tag from thread")
	}
	return nil
}
