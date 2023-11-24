package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book
	AddBook(book *entity.Book) (*entity.Book, error)
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: repo,
	}
}
func (u *bookUsecase) GetAllBooks(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book {
	return u.bookRepo.FindAll(clause, conditions)
}

func (u *bookUsecase) AddBook(book *entity.Book) (*entity.Book, error) {
	return u.bookRepo.CreateBook(book)
}
