package response

import (
	"app/api/domain/entity"
	"time"
)

type UserResponse struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	Name      string         `json:"name"`
	Mail      string         `json:"mail"`
	Image     string         `json:"image"`
	Profile   string         `json:"profile"`
	IsAdmin   int            `json:"is_admin"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	LoginAt   *time.Time     `json:"login_at"`
	Tags      []*TagResponse `json:"tags"`
}

type UsersResponse struct {
	Users []*UserResponse `json:"users"`
}

func ConvertToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		UserID:    user.UserID,
		Name:      user.Name,
		Mail:      user.Mail,
		Image:     user.Image,
		Profile:   user.Profile,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LoginAt:   user.LoginAt,
		Tags:      ConvertToTagsResponse(user.Tags).Tags,
	}
}

func ConvertToUsersResponse(users []*entity.User) *UsersResponse {
	res := make([]*UserResponse, 0, len(users))
	for _, user := range users {
		res = append(res, ConvertToUserResponse(user))
	}
	return &UsersResponse{
		Users: res,
	}
}
