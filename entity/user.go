package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Books     []Book `gorm:"many2many:borrowing_records"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
