package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BookUsecase interface {
	GetAllBooks() ([]models.Book, error)
	GetBooksByTitle(string) ([]models.Book, error)
	CreateBook(models.Book) (*models.Book, error)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(b repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepository: b,
	}
}

func (b *bookUsecase) GetAllBooks() ([]models.Book, error) {
	return b.bookRepository.FindBooks()
}

func (b *bookUsecase) GetBooksByTitle(title string) ([]models.Book, error) {
	return b.bookRepository.FindBooksByTitle(title)
}

func (b *bookUsecase) CreateBook(book models.Book) (newBook *models.Book, err error) {
	existBook, _ := b.bookRepository.FindBooksByTitle(book.Title)
	if len(existBook) != 0 {
		err = apperror.ErrBookAlreadyExist
		return
	}
	return b.bookRepository.NewBook(book)
}
