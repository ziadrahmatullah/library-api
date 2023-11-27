package entity

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

type BorrowingRecords struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	UserId       uint `gorm:"not null"`
	User         User
	BookId       uint `gorm:"not null"`
	Book         Book
	Status       int `gorm:"not null"`
	BorrowedDate time.Time
	ReturnedDate valueobject.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
