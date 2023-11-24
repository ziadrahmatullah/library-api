package entity

type Cover string

type Book struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	Cover       Cover
	AuthorId    uint
	Author      *Author
	User        []User `gorm:"many2many:borrowing_records"`
}
