package entity

import (
	"time"

	"gorm.io/gorm"
)

type BorrowingRecords struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint
	User      User
	BookId    uint
	Book      Book
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
