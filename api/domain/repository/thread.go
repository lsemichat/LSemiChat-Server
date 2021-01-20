package repository

import "app/api/domain/entity"

type ThreadRepository interface {
	Create(thread *entity.Thread) error
	FindAll() ([]*entity.Thread, error)
	FindByID(id string) (*entity.Thread, error)
	FindOnlyPublic() ([]*entity.Thread, error)
	FindMembersByThreadID(id string) ([]*entity.User, error)
	FindByTags(tagIDs []string) ([]*entity.Thread, error)
	FindByUserUUIDs(userUUIDs []string) ([]*entity.Thread, error)
	FindByNames(names []string) ([]*entity.Thread, error)
	Update(thread *entity.Thread) error
	AddMember(id, threadID, userID string, isAdmin int) error
	RemoveMember(threadID, userID string) error
	Delete(id string) error
}
