package response

import "app/api/domain/entity"

type TagResponse struct {
	ID       string            `json:"id"`
	Tag      string            `json:"tag"`
	Category *CategoryResponse `json:"category"`
}

type TagsResponse struct {
	Tags []*TagResponse `json:"tags"`
}

func ConvertToTagResponse(tag *entity.Tag) *TagResponse {
	return &TagResponse{
		ID:       tag.ID,
		Tag:      tag.Tag,
		Category: ConvertToCategoryResponse(tag.Category),
	}
}

func ConvertToTagsResponse(tags []*entity.Tag) *TagsResponse {
	result := make([]*TagResponse, 0, len(tags))
	for _, tag := range tags {
		result = append(result, ConvertToTagResponse(tag))
	}
	return &TagsResponse{
		Tags: result,
	}
}
