package entity

type BorrowingRecords struct {
	Id     uint `gorm:"primaryKey;autoIncrement"`
	UserId uint
	User   User
	BookId uint
	Book   Book
	Status int
}
