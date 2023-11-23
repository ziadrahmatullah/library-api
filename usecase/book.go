package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
)

type BookUsecase interface {
	GetAllBooks() []*entity.Book
	FindBooksByTitle(title string) []*entity.Book
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: repo,
	}
}
func (u *bookUsecase) GetAllBooks() []*entity.Book {
	return u.bookRepo.FindAll()
}

func (u *bookUsecase) FindBooksByTitle(name string) []*entity.Book {
	return u.bookRepo.FindByTitle(name)
}
