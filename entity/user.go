package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Phone     string `gorm:"not null;unique"`
	Books     []Book `gorm:"many2many:borrowing_records"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
