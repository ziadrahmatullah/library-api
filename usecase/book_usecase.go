package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BookUsecase interface {
	GetAllBooks(context.Context) ([]models.Book, error)
	GetBooksByTitle(context.Context, string) ([]models.Book, error)
	CreateBook(context.Context, models.Book) (*models.Book, error)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(b repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepository: b,
	}
}

func (b *bookUsecase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return b.bookRepository.FindBooks(ctx)
}

func (b *bookUsecase) GetBooksByTitle(ctx context.Context, title string) ([]models.Book, error) {
	return b.bookRepository.FindBooksByTitle(ctx, title)
}

func (b *bookUsecase) CreateBook(ctx context.Context, book models.Book) (newBook *models.Book, err error) {
	existBook, _ := b.bookRepository.FindBooksByTitle(ctx, book.Title)
	if len(existBook) != 0 {
		err = apperror.ErrBookAlreadyExist
		return
	}
	return b.bookRepository.NewBook(ctx,book)
}
