package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(ctx context.Context, query *valueobject.Query) ([]*entity.Book, error)
	GetSingleBook(ctx context.Context, query *valueobject.Query) (*entity.Book, error)
	AddBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
}

type bookUsecase struct {
	bookRepo   repository.BookRepository
	authorRepo repository.AuthorRepository
}

func NewBookUsecase(bookRepo repository.BookRepository, authorRepo repository.AuthorRepository) BookUsecase {
	return &bookUsecase{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}
func (u *bookUsecase) GetAllBooks(ctx context.Context, query *valueobject.Query) ([]*entity.Book, error) {
	return u.bookRepo.Find(ctx, query)
}

func (u *bookUsecase) GetSingleBook(ctx context.Context, query *valueobject.Query) (*entity.Book, error) {
	return u.bookRepo.First(ctx, query)
}

func (u *bookUsecase) AddBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	bookQuery := valueobject.NewQuery().Condition("title", valueobject.Equal, book.Title)
	b, err := u.GetSingleBook(ctx, bookQuery)
	if err != nil {
		return nil, err
	}
	if b != nil {
		return nil, apperror.NewResourceAlreadyExist("book", "title", book.Title)
	}
	authorQuery := valueobject.NewQuery().Condition("id", valueobject.Equal, book.AuthorId)
	author, err := u.authorRepo.First(ctx, authorQuery)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, apperror.NewResourceNotFound("author", "id", book.AuthorId)
	}
	createdBook, err := u.bookRepo.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}
