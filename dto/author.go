package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type AuthorResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewFromAuthor(author *entity.Author) *AuthorResponse {
	return &AuthorResponse{
		Id:   author.Id,
		Name: author.Name,
	}
}
