package entity

type Cover string

type Book struct {
	Id          uint   `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Quantity    int    `json:"quantity" gorm:"quantity"`
	Cover       Cover  `json:"cover" gorm:"cover"`
}
