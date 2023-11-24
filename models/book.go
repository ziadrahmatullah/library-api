package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null" binding:"required,max=35"`
	Description string `gorm:"not null" binding:"required"`
	Quantity    uint   `gorm:"not null" binding:"required,min=0"`
	Cover       string
	AuthorId    uint   `gorm:"column:author_id;not null" json:"author_id"`
	Author      Author `gorm:"foreignKey:author_id;references:id"`
}

//dto
