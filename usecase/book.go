package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(c valueobject.Clause) []*entity.Book
	FindBooksByTitle(title string) []*entity.Book
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
func (u *bookUsecase) GetAllBooks(c valueobject.Clause) []*entity.Book {
	return u.bookRepo.FindAll(c)
}

func (u *bookUsecase) FindBooksByTitle(name string) []*entity.Book {
	return u.bookRepo.FindByTitle(name)
}

func (u *bookUsecase) AddBook(book *entity.Book) (*entity.Book, error) {
	return u.bookRepo.CreateBook(book)
}
