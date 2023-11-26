package entity

import (
	"time"

	"gorm.io/gorm"
)

type BorrowingRecords struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint `gorm:"not null"`
	User      User
	BookId    uint `gorm:"not null"`
	Book      Book
	Status    int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
