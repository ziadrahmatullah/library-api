package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
)

type LoginReq struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type RegisterReq struct{
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Phone    string `binding:"required" json:"phone"`
	Password string `binding:"required" json:"password"`
}

type RegisterRes struct{
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (r *RegisterReq) ToUserModelFromRegisterDTO(password string) models.User {
	return models.User{
		Name:     r.Name,
		Email:    r.Email,
		Phone:    r.Phone,
		Password: password,
	}
}

func ToUserResponsDTOFromModel(user *models.User) *RegisterRes{
	return &RegisterRes{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}