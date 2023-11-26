package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    uint   `gorm:"not null"`
	Cover       string
	AuthorId    uint   `gorm:"not null"`
	Author      Author `gorm:"foreignKey:author_id;references:id"`
}
