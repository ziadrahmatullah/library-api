package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewFromUser(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}

func NewFromUsers(users []*entity.User) []*UserResponse {
	return newResponsesFromEntities(users, NewFromUser)
}
