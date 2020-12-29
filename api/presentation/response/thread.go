package response

import (
	"app/api/domain/entity"
	"time"
)

type ThreadResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Descirption string         `json:"description"`
	LimitUsers  int            `json:"limit_users"`
	IsPublic    int            `json:"is_public"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	Author      *UserResponse  `json:"author"`
	Tags        []*TagResponse `json:"tags"`
}

type ThreadsResponse struct {
	Threads []*ThreadResponse `json:"threads"`
}

func ConvertToThreadResponse(thread *entity.Thread) *ThreadResponse {
	return &ThreadResponse{
		ID:          thread.ID,
		Name:        thread.Name,
		Descirption: thread.Description,
		LimitUsers:  thread.LimitUsers,
		IsPublic:    thread.IsPublic,
		CreatedAt:   thread.CreatedAt,
		UpdatedAt:   thread.UpdatedAt,
		Author:      ConvertToUserResponse(thread.Author),
		Tags:        ConvertToTagsResponse(thread.Tags).Tags,
	}
}

func ConvertToThreadsResponse(threads []*entity.Thread) *ThreadsResponse {
	res := make([]*ThreadResponse, 0, len(threads))
	for _, thread := range threads {
		res = append(res, ConvertToThreadResponse(thread))
	}
	return &ThreadsResponse{
		Threads: res,
	}
}
