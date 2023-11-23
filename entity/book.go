package entity

type Cover string

type Book struct {
	Id          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	Cover       Cover  `json:"cover" gorm:"not null"`
}
