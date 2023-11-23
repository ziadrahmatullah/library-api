package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    uint   `gorm:"not null"`
	Cover       string `gorm:"not null"`
}
