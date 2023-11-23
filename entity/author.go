package entity

type Author struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"string" gorm:"not null"`
}
