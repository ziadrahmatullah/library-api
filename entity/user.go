package entity

type User struct {
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Email string `gorm:"unique"`
	Phone string `gorm:"unique"`
	Books []Book `gorm:"many2many:borrowing_records"`
}
