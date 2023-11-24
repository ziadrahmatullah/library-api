package entity

type Author struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}
