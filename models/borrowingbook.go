package models

import "gorm.io/gorm"

type BorrowingBooks struct {
	gorm.Model
	BookId uint   `gorm:"column:book_id;not null" json:"book_id"`
	UserId uint   `gorm:"column:user_id;not null" json:"user_id"`
	Status string `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
	Book   Book   `gorm:"foreignKey:BookID;references:ID"`
}
