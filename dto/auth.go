package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func (r *RegisterRequest) ToUser() *entity.User {
	return &entity.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
	}
}
