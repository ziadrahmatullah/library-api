package usecase

import (
	"context"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(ctx context.Context, query valueobject.Query) []*entity.Book
	GetSingleBook(ctx context.Context, query valueobject.Query) *entity.Book
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
func (u *bookUsecase) GetAllBooks(ctx context.Context, query valueobject.Query) []*entity.Book {
	return u.bookRepo.Find(ctx, query)
}

func (u *bookUsecase) GetSingleBook(ctx context.Context, query valueobject.Query) *entity.Book {
	return u.bookRepo.First(ctx, query)
}

func (u *bookUsecase) AddBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	bookCondition := *valueobject.NewCondition("title", valueobject.Equal, book.Title)
	bookQuery := valueobject.Query{
		Conditions: []valueobject.Condition{bookCondition},
	}
	b := u.GetSingleBook(ctx, bookQuery)
	if b != nil {
		return nil, apperror.ErrAlreadyExist{
			Resource: "book",
			Field:    "title",
			Value:    b.Title,
		}
	}
	authorCondition := *valueobject.NewCondition("id", valueobject.Equal, strconv.Itoa(int(book.AuthorId)))
	authorQuery := valueobject.Query{Conditions: []valueobject.Condition{authorCondition}}
	author := u.authorRepo.First(ctx, authorQuery)
	if author == nil {
		return nil, apperror.ErrNotFound{
			Resource: "author",
			Field:    "id",
			Value:    strconv.Itoa(int(book.AuthorId)),
		}
	}
	return u.bookRepo.Create(ctx, book)
}
