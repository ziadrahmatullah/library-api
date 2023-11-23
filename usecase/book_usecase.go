package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BookUsecase interface{
	GetAllBooks() ([]models.Book, error)
}

type bookUsecase struct{
	bookRepository repository.BookRepository
}

func NewBookUsecase(b repository.BookRepository) BookUsecase{
	return &bookUsecase{
		bookRepository: b,
	}
}

func (b *bookUsecase) GetAllBooks() (books []models.Book, err error){
	return b.bookRepository.FindAllBooks()
}
